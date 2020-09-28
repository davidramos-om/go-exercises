package main

import "fmt"

// Bus carries Passengers from A to B if they have a valid bus ticket.
type Bus struct {
	Company     BusCompany
	name        string
	passengers  Passengers
	stops       []*BusStop
	currentStop int16
}

// NewBus returns a new Bus with an empty passenger set.
func NewBus(name string) Bus {
	b := Bus{}
	b.name = name
	b.currentStop = -1
	b.passengers = NewPassengerSet()
	return b
}

// AddStop adds the given BusStop to the list of stops that the Bus will stop at. Each stop is visited in order.
func (b *Bus) AddStop(busStop *BusStop) {
	b.stops = append(b.stops, busStop)
}

// add adds a single Passenger to the Bus. For brevity, we don't care too much about accidentally adding the same Passenger more than once.
func (b *Bus) add(p Passenger) {
	if b.passengers == nil {
		b.passengers = make(map[string]Passenger)
	}
	b.passengers[p.SSN] = p

	fmt.Printf("%s: boarded passenger %q\n", b.name, p.Name)
	fmt.Println("")
}

// Remove removes a single Passenger from the Bus.
func (b *Bus) Remove(p Passenger) {
	delete(b.passengers, p.SSN)

	fmt.Printf("\n%s: unboarded passenger %q\n\n", b.name, p.Name)

}

// Board adds the given Passenger to the Bus and charges them a ticket price calculated by chargeFn if they don't already have a paid ticket.
// Board returns false if the Passenger was not allowed to board the Bus.
func (b *Bus) Board(p *Passenger, chargeFn PriceCalculator) bool {

	var allowed bool // Default value is false

	if p.HasValidTicket {
		allowed = true
	} else {
		amount := chargeFn(*p)
		p2 := p.Charge(amount)
		p = &p2
		allowed = true
	}

	if allowed {
		b.add(*p)
	}
	return allowed
}

// Go takes the Bus to the next BusStop. Go returns true if there are still more stops to visit.
func (b *Bus) Go() bool {

	b.currentStop++
	lastIndex := int16(len(b.stops) - 1)

	if b.currentStop == lastIndex {
		fmt.Printf("%s: reached the end of the line, everybody out\n", b.name)
		b.VisitPassengers(func(p Passenger) {
			b.Remove(p)
		})
		return false
	}

	if b.currentStop == 0 {
		fmt.Printf("%s: starting\n", b.name)
	} else {
		fmt.Println("")
		fmt.Printf("%s: carrying %d passengers: heading for next stop\n", b.name, len(b.passengers))
	}

	curr := b.stops[b.currentStop]
	fmt.Println("")
	fmt.Printf("%s: arriving at %q\n", b.name, curr.Name)
	fmt.Println("")
	curr.NotifyBusArrival(b)
	return b.currentStop < lastIndex
}

// Manifest asks Passengers for a SSN manifest and returns it.
func (b Bus) Manifest() []string {
	return b.passengers.Manifest()
}

// VisitPassengers calls function visitor for each Passenger on the bus.
func (b *Bus) VisitPassengers(visitor func(Passenger)) {
	b.passengers.Visit(visitor)
}

// FindPassenger returns the Passenger that matches the given SSN, if found. Otherwise, an empty Passenger is returned.
func (b *Bus) FindPassenger(ssn string) Passenger {
	if p, ok := b.passengers[ssn]; ok {
		return p
	}
	return Passenger{} // A nobody.
}

// UpdatePassengers calls function visitor for each Passenger on the bus. Passengers are passed by reference and may be modified.
func (b *Bus) UpdatePassengers(visitor func(*Passenger)) {
	ps := make(map[string]Passenger, len(b.passengers))
	for ssn, p := range b.passengers {
		visitor(&p)
		ps[ssn] = p
	}
	b.passengers = ps
}

// NotifyBoardingIntent is called by BusStop every time a Prospect arrives and instructs the Bus to signal its arrival when at that BusStop.
func (b *Bus) NotifyBoardingIntent(busStop *BusStop) {
	if b.StopsAt(busStop) {
		return // We already intend to stop here.
	}
	b.AddStop(busStop)
}

// NotifyArrival notifies the current BusStop that the Bus has arrived.
func (b *Bus) NotifyArrival() {
	curr := b.stops[b.currentStop]
	curr.NotifyBusArrival(b)
}

// StopsAt checks if Bus stops at the given BusStop, and returns true if it does, and false otherwise.
func (b Bus) StopsAt(busStop *BusStop) bool {
	for _, stop := range b.stops {
		if stop.Equals(busStop) {
			return true
		}
	}
	return false
}

// CurrentStop returns the BusStop that the Bus is currently stopped at.
func (b Bus) CurrentStop() *BusStop {
	return b.stops[b.currentStop]
}
