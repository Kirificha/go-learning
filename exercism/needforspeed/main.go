package speed

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	newcar := Car{battery: 100, batteryDrain: 2, speed: 5, distance: 0}
	return newcar
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

func main() {
	distance := 800
	track := NewTrack(distance)
	car := Car{}
	newcar := NewCar(car.speed, car.batteryDrain)
	car = Drive(car)

	fmt.Println(distance, track, newcar, car)
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	newd := Track{distance: distance}
	return newd
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.

func Drive(car Car) Car {
	battery := car.battery - car.batteryDrain
	if battery < 0 {
		distancen := 0 + car.distance
		return Car{battery: battery, batteryDrain: car.batteryDrain, speed: car.speed, distance: distancen}
	} else {
		distancen := car.distance + car.speed
		return Car{battery: battery, batteryDrain: car.batteryDrain, speed: car.speed, distance: distancen}
	}
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	car_dist := float64(car.battery) * 2.5
	if car_dist < float64(track.distance) {
		return false
	}
	return true
	panic("Please implement the CanFinish function")
}
