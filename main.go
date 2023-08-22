package main

import (
	"fmt"
)

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Van
)

type ParkingLot struct {
	totalSpaces     int
	availableSpaces map[VehicleType]int
	vansOccupied    int
}

func NewParkingLot(motorcycleSpaces, carSpaces, vanSpaces int) *ParkingLot {
	return &ParkingLot{
		totalSpaces: motorcycleSpaces + carSpaces + vanSpaces,
		availableSpaces: map[VehicleType]int{
			Motorcycle: motorcycleSpaces,
			Car:        carSpaces,
			Van:        vanSpaces,
		},
		vansOccupied: 0,
	}
}

func (p *ParkingLot) SpacesAvailable() int {
	return p.availableSpaces[Motorcycle] + p.availableSpaces[Car] + p.availableSpaces[Van]
}

func (p *ParkingLot) TotalSpaces() int {
	return p.totalSpaces
}

func (p *ParkingLot) IsFull() bool {
	for _, spaces := range p.availableSpaces {
		if spaces > 0 {
			return false
		}
	}
	return true
}

func (p *ParkingLot) IsEmpty() bool {
	for _, spaces := range p.availableSpaces {
		if spaces <= 0 {
			return false
		}
	}
	return true
}

func (p *ParkingLot) IsFullFor(vehicleType VehicleType) bool {
	return p.availableSpaces[vehicleType] == 0
}

func (p *ParkingLot) Park(vehicleType VehicleType) bool {
	switch vehicleType {
	case Motorcycle:
		p.availableSpaces[Motorcycle]--
	case Car:
		if p.availableSpaces[Car] > 0 {
			p.availableSpaces[Car]--
		} else if p.availableSpaces[Van] > 0 {
			p.availableSpaces[Van]--
		} else {
			return false
		}
	case Van:
		if p.availableSpaces[Van] > 0 {
			p.availableSpaces[Van]--
		} else if p.availableSpaces[Car] >= 3 {
			p.availableSpaces[Car] -= 3
		} else {
			return false
		}
		p.vansOccupied++
	default:
		return false
	}
	return true
}

func (p *ParkingLot) SpacesAvailableFor(vehicleType VehicleType) int {
	return p.availableSpaces[vehicleType]
}

func (p *ParkingLot) SpacesOccupiedByVans() int {
	return p.vansOccupied
}

func main() {
	parkinkLot := NewParkingLot(1, 1, 0)

	motorcycleParked := parkinkLot.Park(Motorcycle)
	carParked := parkinkLot.Park(Car)
	vanParked := parkinkLot.Park(Van)

	fmt.Printf("Motorcycle parked: %t\n", motorcycleParked)
	fmt.Printf("Car parked: %t\n", carParked)
	fmt.Printf("Van parked: %t\n", vanParked)

	fmt.Printf("Park is full: %t\n", parkinkLot.IsFull())
	fmt.Printf("Park is empty: %t\n", parkinkLot.IsEmpty())
	fmt.Printf("Spaces available: %d\n", parkinkLot.SpacesAvailable())

	fmt.Printf("Total spaces: %d\n", parkinkLot.TotalSpaces())
	fmt.Printf("Available spaces: %d\n", parkinkLot.SpacesAvailable())
	fmt.Printf("Vans occupied: %d\n", parkinkLot.SpacesOccupiedByVans())
	fmt.Printf("All spaces of car: %d\n", parkinkLot.SpacesAvailableFor(Car))
	fmt.Printf("All spaces of van: %d\n", parkinkLot.SpacesAvailableFor(Van))
}
