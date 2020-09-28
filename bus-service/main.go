package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\n ")
	fmt.Println("*********************************************************************")
	fmt.Println("\nStarting simulation")
	expressLine := NewBus("Express Line")

	s1 := BusStop{Name: "Downtown"}
	s2 := BusStop{Name: "The University"}
	s3 := BusStop{Name: "The Village"}

	expressLine.AddStop(&s1)
	expressLine.AddStop(&s2)
	expressLine.AddStop(&s3)

	john := Prospect{
		SSN:         "12345612-22",
		Name:        "John Smith",
		Destination: &s2,
	}

	betty := Prospect{
		SSN:         "11223322-67",
		Name:        "Betty Miller",
		Destination: &s3,
	}

	s1.NotifyProspectArrival(john)
	s1.NotifyProspectArrival(betty)

	for expressLine.Go() {
		expressLine.VisitPassengers(func(p Passenger) {
			fmt.Printf("Passenger %q is heading to %q\n", p.Name, p.Destination.Name)
		})
	}
	fmt.Println("Simulation done")
	fmt.Println("*********************************************************************")
	fmt.Println("\n\n ")
}
