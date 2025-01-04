package main

import (
	"fmt"
)

// bike interface
type Bike interface {
	RevEngine() string
	GetSpecs() string
	TopSpeed() int
}

// motorcycle implements bike
type Motorcycle struct {
	Brand    string
	Model    string
	CC       int
	Category string
	MaxSpeed int
	HasABS   bool
}

func (m Motorcycle) RevEngine() string {
	if m.Model == "MT-15" {
		return "Vroom Vrroom! * MT-15 VVA engine roars *"
	}
	return "Generic motorcycle sounds"
}

func (m Motorcycle) GetSpecs() string {
	return fmt.Sprintf("Bike: %s %s\nCategory: %s\nEngine: %dcc\nABS: %v",
		m.Brand, m.Model, m.Category, m.CC, m.HasABS)
}

func (m Motorcycle) TopSpeed() int {
	return m.MaxSpeed
}

func main() {
	bike := Motorcycle{
		Brand:"Yamaha",
		Model:"MT-15",
		CC:155,
		Category:"Naked",
		MaxSpeed:140,
		HasABS:true,
	}

	fmt.Println(bike.RevEngine())
	fmt.Println(bike.GetSpecs())
	fmt.Printf("Top Speed: %d km/h\n", bike.TopSpeed())
}
