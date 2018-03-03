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

func (v *Vehicle) Drive(allTrips []*Trip) {
	// it has done every of his trips
	if v.CurrentRide > len(v.Trips) {
		return
	}

	tripToGoTo := allTrips[v.Trips[v.CurrentRide]]
	tripToGoTo.SomeoneIsOnIt()
	v.DriveTo(tripToGoTo.Start.X, tripToGoTo.Start.Y)
}

func (v *Vehicle) GetPosition() (int32, int32) {
	return v.CurrentPosition.X, v.CurrentPosition.Y
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
}
