package main

import "strconv"

// Prospect is a potential Passenger. Prospects wait at BusStops to board Buses.
type Prospect struct {
	SSN         string
	Name        string
	Destination *BusStop
}

// ToPassenger returns a Passenger with the same SSN as his or her Prospect.
func (p Prospect) ToPassenger() Passenger {

	age, err := strconv.ParseInt(p.SSN[len(p.SSN)-2:], 10, 8)

	if err == nil {
		age = 20
	}

	_age := int8(age)

	return Passenger{SSN: p.SSN, Name: p.Name, Age: _age, Destination: p.Destination}
}
