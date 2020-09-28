package main

// BusStop represents a place where a Bus can stop and signal to prospects (future passengers)
// that they may board.
type BusStop struct {
	Name      string
	prospects []Prospect
	busses    []Bus
}

// Equals returns true if the given BusStop is the same as the receiver.
func (b *BusStop) Equals(busStop *BusStop) bool {
	return b.Name == busStop.Name
}

// NotifyBusArrival is called by Bus upon arrival.
func (b *BusStop) NotifyBusArrival(bus *Bus) {

	bus.VisitPassengers(func(p Passenger) {
		if bus.CurrentStop().Equals(p.Destination) {
			bus.Remove(p)
		}
	})

	for _, p := range b.prospects {
		if bus.StopsAt(p.Destination) {
			pas := p.ToPassenger()
			bus.Board(&pas, bus.Company.GetPricing())
		}
	}
}

// NotifyProspectArrival is called whenever a prospect arrives at Busstop.
func (bs *BusStop) NotifyProspectArrival(p Prospect) {
	bs.prospects = append(bs.prospects, p)

	// Find all Busses on this route.
	for _, bus := range bs.busses {
		if bus.StopsAt(p.Destination) {
			bus.NotifyBoardingIntent(bs)
		}
	}
}
