package ghashcode

import (
	"math"

	config "../config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Vehicle struct {
	CurrentPosition Coordinates

	CurrentRide int
	OnRide      bool
	Trips       []int32
	Enabled     bool
}

func (v *Vehicle) DrawToWindow(win *pixelgl.Window) {
	imd := imdraw.New(nil)

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

	imd.Draw(win)
}

func (v *Vehicle) SetPosition(x, y int32) {
	v.CurrentPosition.X = x
	v.CurrentPosition.Y = y
}

func (v *Vehicle) Drive(allTrips []*Trip, step int) {
	// it has done every of his trips
	if v.CurrentRide >= len(v.Trips) {
		return
	}

	tripToGoTo := allTrips[v.Trips[v.CurrentRide]]

	if v.OnRide {
		tripToGoTo.StartTrip()
		if int32(step) < tripToGoTo.EarliestStart {
			tripToGoTo.WarnEarly()
			return
		}
		v.DriveOnTrip(tripToGoTo.End.X, tripToGoTo.End.Y)
		cx, cy := v.GetPosition()
		if cx == tripToGoTo.End.X && cy == tripToGoTo.End.Y {
			tripToGoTo.Finish()
			v.NextTrip()
		}
		return
	}

	tripToGoTo.SomeoneIsOnIt()
	v.DriveTo(tripToGoTo.Start.X, tripToGoTo.Start.Y)
}

func (v *Vehicle) GetPosition() (int32, int32) {
	return v.CurrentPosition.X, v.CurrentPosition.Y
}

func (v *Vehicle) NextTrip() {
	// move forward next trip
	v.CurrentRide++
	v.OnRide = false
	if v.CurrentRide >= len(v.Trips) {
		// no trip ? hide the car
		v.Enabled = false
	}
}

func (v *Vehicle) DriveOnTrip(x, y int32) {
	cx, cy := v.GetPosition()
	if cx > x {
		cx--
	} else if cx < x {
		cx++
	} else if cy > y {
		cy--
	} else if cy < y {
		cy++
	}
	v.SetPosition(cx, cy)
}

func (v *Vehicle) DriveTo(x, y int32) {
	cx, cy := v.GetPosition()
	dx := math.Abs(float64(cx - x))
	dy := math.Abs(float64(cy - y))

	// fmt.Printf("goto %d %d\n", x, y)
	nx := cx
	ny := cy
	if dx < dy {
		if cy < y {
			ny++
		} else if cy > y {
			ny--
		}
	} else if dx > dy {
		if cx < x {
			nx++
		} else if cx > x {
			nx--
		}
	} else if dx == dy {
		if cx < x {
			nx++
		} else if cx > x {
			nx--
		} else if cy < y {
			ny++
		} else if cy > y {
			ny--
		}
	}
	// fmt.Printf("driving to %d, %d\n", x, y)
	// fmt.Printf("current pos [%d, %d]", v.CurrentPosition.X, v.CurrentPosition.Y)
	// fmt.Printf("-\n")
	// fmt.Printf("moving from (%d, %d) to (%d, %d)\n", cx, cy, nx, ny)
	v.SetPosition(nx, ny)
	if x == nx && y == ny {
		v.OnRide = true
	}
}
