package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var drivenDistance int = car.distance
	var remainingBattery int = car.battery

	if car.battery >= car.batteryDrain {
		drivenDistance = car.distance + car.speed
		remainingBattery = car.battery - car.batteryDrain
	}

	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		distance:     drivenDistance,
		battery:      remainingBattery,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	unitDistance := track.distance / car.speed
	return car.battery-(car.batteryDrain*unitDistance) >= 0
}
