package main

import "time"

// BusCompany represents the bus company responsible for the Bus service. BusCompany determines price policies.
type BusCompany string

// WorkdayPricing charges EUR 6 for regular Passengers and EUR 4.5 for seniors during workdays.
func WorkdayPricing(p Passenger) float64 {
	if p.IsSenior() {
		return 4.5
	}
	return 6.0
}

// WeekendPricing charges EUR 5 for regular Passengers and EUR 3.5 for seniors during weekends.
func WeekendPricing(p Passenger) float64 {
	if p.IsSenior() {
		return 3.5
	}
	return 5.0
}

// PriceCalculator is the type used by BusCompany to determine the ticket price for a Passenger.
// PriceCalculator returns the ticket price in the local currency.
type PriceCalculator func(p Passenger) float64

// GetPricing returns a price calculator based on the pricing policy of the day.
func (b BusCompany) GetPricing() PriceCalculator {

	wd := time.Now().Weekday()
	if wd == time.Saturday || wd == time.Sunday {
		return WeekendPricing
	}
	return WorkdayPricing
}
