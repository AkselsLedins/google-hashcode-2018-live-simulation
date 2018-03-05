package ghashcode

import (
	"math"

	config "../config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Vehicle a structure which represents a vehicle in the simulation
// Quite self-explanatory
type Vehicle struct {
	// current position of the vehicle
	CurrentPosition Coordinates

	// Trips an array of trips' indexes
	// (and not their ids)
	Trips []int32
	// index of the current trip in Trips
	CurrentTrip int
	// if he is currently assigned on a trip
	OnRide  bool
	Enabled bool
}

// AddToImd adds to the imd batch the graphic point of the vehicle
// We only call .draw once for all the vehicles
func (v *Vehicle) AddToImd(imd *imdraw.IMDraw) {
	imd.Color = config.Config.UI.VehicleDefaultColor
	imd.EndShape = imdraw.RoundEndShape
	squareSize := config.Config.UI.SquareSize

	offsetX := v.CurrentPosition.X*squareSize + squareSize
	offsetY := v.CurrentPosition.Y*squareSize + squareSize

	drawX := float64(offsetX)
	drawY := float64(offsetY)

	imd.Push(pixel.V(drawX, drawY))
	imd.Push(pixel.V(drawX, drawY))

	imd.Line(config.Config.UI.VehicleSize)
}

func (v *Vehicle) Drive(allTrips []*Trip, step int, bonus int16) (score int) {
	// until the vehicle handles all scheduled rides
	if v.CurrentTrip >= len(v.Trips) {
		return
	}

	tripToGoTo := allTrips[v.Trips[v.CurrentTrip]]
	tripToGoTo.SomeoneIsOnIt()

	// first, the vehicle drives from its current intersection ([0,0] at the beginning of the simulation) to the
	// start intersection of the next ride (unless the vehicle is already in this intersection)
	if !v.OnRide {
		v.DriveTo(tripToGoTo.Start.X, tripToGoTo.Start.Y)
	}

	if v.OnRide {
		tripToGoTo.StartTrip(step, bonus)
		// then, if the current step is earlier than the earliest start of the next ride,
		// the vehicle waits until that step
		if int32(step) < tripToGoTo.EarliestStart {
			tripToGoTo.WarnEarly()
			return
		}

		// then, the vehicle drives to the finish intersection
		v.DriveOnTrip(tripToGoTo.End.X, tripToGoTo.End.Y)
		currentX, currentY := v.GetPosition()

		if currentX == tripToGoTo.End.X && currentY == tripToGoTo.End.Y {
			score += tripToGoTo.Finish(int32(step))
			// then, the process repeats for the next assigned ride,
			v.NextTrip()
			return
		}
		return
	}

	return
}

// NextTrip move forward next trip or disable the vehicle
// if there isnt any trips left
func (v *Vehicle) NextTrip() {
	v.CurrentTrip++
	v.OnRide = false

	// no trip ? hide the car
	if v.CurrentTrip >= len(v.Trips) {
		v.Enabled = false
	}
}

// DriveOnTrip it will make the vehicle
// follow the line drawn on the screen by the trip
func (v *Vehicle) DriveOnTrip(destinationX, destinationY int32) {
	currentX, currentY := v.GetPosition()
	if currentX > destinationX {
		currentX--
	} else if currentX < destinationX {
		currentX++
	} else if currentY > destinationY {
		currentY--
	} else if currentY < destinationY {
		currentY++
	}

	// update vehicle position
	v.SetPosition(currentX, currentY)
}

// DriveTo destinationX, destinationY coordinates
// It will make the vehicle move in a straight line to a point
func (v *Vehicle) DriveTo(destinationX, destinationY int32) {
	// retrieve current position
	currentX, currentY := v.GetPosition()
	if destinationX == currentX && destinationY == currentY {
		v.OnRide = true
		return
	}

	// calculate the absolute difference between the positions
	dx := math.Abs(float64(currentX - destinationX))
	dy := math.Abs(float64(currentY - destinationY))

	// newX and newY
	nx := currentX
	ny := currentY
	if dx < dy {
		if currentY < destinationY {
			ny++
		} else if currentY > destinationY {
			ny--
		}
	} else if dx > dy {
		if currentX < destinationX {
			nx++
		} else if currentX > destinationX {
			nx--
		}
	} else if dx == dy {
		if currentX < destinationX {
			nx++
		} else if currentX > destinationX {
			nx--
		} else if currentY < destinationY {
			ny++
		} else if currentY > destinationX {
			ny--
		}
	}

	// update the vehicle position
	v.SetPosition(nx, ny)

	// it reached his destination so we can consider him ready to start his ride
	if destinationX == nx && destinationY == ny {
		v.OnRide = true
	}
}

// NewVehicle a constructor for Trip
// takes its trips as parameters
func NewVehicle(trips []int32) *Vehicle {
	v := new(Vehicle)

	// default values
	v.SetPosition(0, 0)
	v.Enabled = true
	v.CurrentTrip = 0
	v.Trips = trips

	return v
}

// SetPosition setter for CurrentPosition
func (v *Vehicle) SetPosition(x, y int32) {
	v.CurrentPosition.X = x
	v.CurrentPosition.Y = y
}

// GetPosition getter for CurrentPosition
func (v *Vehicle) GetPosition() (int32, int32) {
	return v.CurrentPosition.X, v.CurrentPosition.Y
}
