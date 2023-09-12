package mocking

//go:generate mockgen -destination=./mocks/mock_truck.go -package=mocks golab.com/m/v2/mocking Truck

// Truck is a vehicle with a bed
type Truck interface {
	Vehicle
}
