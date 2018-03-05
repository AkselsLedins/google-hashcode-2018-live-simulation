package ghashcode

import (
	"image/color"
	"math"

	"golang.org/x/image/colornames"

	config "../config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Trip struct {
	ID int

	Start Coordinates
	End   Coordinates

	Distance int
	Bonus    int

	EarliestStart int32
	LatestFinish  int32

	Color color.RGBA

	Taken      bool
	InProgress bool
	Failed     bool

	// precomputed values
	GraphicLine *imdraw.IMDraw
}

func (t *Trip) DrawToWindow(win *pixelgl.Window) {
	if !t.Taken {
		return
	}

	t.GraphicLine.Color = t.Color
	t.GraphicLine.Draw(win)
}

func (t *Trip) AddToImd(imd *imdraw.IMDraw) {
	if !t.Taken {
		return
	}

	imd.Color = t.Color

	/* start point */
	startX := t.Start.X*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	startY := t.Start.Y*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(startX), float64(startY)))
	/* second point */
	intermediateX := (t.End.X)*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	intermediateY := (t.Start.Y)*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(intermediateX), float64(intermediateY)))
	/* final point */
	endX := t.End.X*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	endY := t.End.Y*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(endX), float64(endY)))

	imd.Line(2)

	imd.Color = t.Color
	imd.EndShape = imdraw.RoundEndShape

	imd.Push(pixel.V(float64(startX), float64(startY)))
	imd.Push(pixel.V(float64(startX), float64(startY)))

	imd.Line(10)
}

func (t *Trip) SetStart(x, y int32) {
	t.Start.X = x
	t.Start.Y = y
}

func (t *Trip) SetEnd(x, y int32) {
	t.End.X = x
	t.End.Y = y
}

func (t *Trip) SomeoneIsOnIt() {
	t.Color = colornames.Beige
	t.Taken = true
}

func (t *Trip) StartTrip(step int) {
	if step == int(t.EarliestStart) {
		t.Bonus += 2
	}
	t.Color = colornames.Cyan
}

func (t *Trip) Finish(step int32) int {
	failed := false
	if step > t.LatestFinish {
		failed = true
	}
	t.Failed = failed

	if !failed {
		t.Color = colornames.Green
		return t.Distance + t.Bonus
	} else {
		t.Color = colornames.Red
		return 0
	}
}

func (t *Trip) WarnEarly() {
	t.Color = colornames.Yellow
}

func NewTrip(id int, a, b, x, y, s, f int32) *Trip {
	trip := new(Trip)

	trip.ID = id

	trip.SetStart(a, b)
	trip.SetEnd(x, y)
	trip.EarliestStart = s
	trip.LatestFinish = f

	// default values
	trip.InProgress = false
	trip.Taken = false
	trip.Failed = false
	trip.Color = config.Config.UI.TripDefaultColor

	// precomputed values
	trip.Distance = int(math.Abs(float64(a-x)) + math.Abs(float64(b-y)))
	trip.Bonus = 0
	imd := imdraw.New(nil)

	imd.Color = trip.Color
	imd.EndShape = imdraw.RoundEndShape

	trip.GraphicLine = imd

	return trip
}
