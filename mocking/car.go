package mocking

//go:generate mockgen -destination=./mocks/mock_car.go -package=mocks golab.com/m/v2/mocking Car

type Car interface {
	Vehicle
}
