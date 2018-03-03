package ghashcode

import (
	"image/color"

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

	EarliestStart int32
	LatestFinish  int32

	Color color.RGBA

	Taken      bool
	InProgress bool
}

func (t *Trip) DrawToWindow(win *pixelgl.Window) {
	if !t.Taken {
		return
	}
	imd := imdraw.New(nil)

	imd.Color = t.Color
	imd.EndShape = imdraw.RoundEndShape

	/* start point */
	startX := t.Start.X*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	startY := t.Start.Y*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(startX), float64(startY)))
	/* second point */
	x := (t.End.X)*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	y := (t.Start.Y)*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(x), float64(y)))
	/* final point */
	endX := t.End.X*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	endY := t.End.Y*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(endX), float64(endY)))

	imd.Line(2)

	imd.Draw(win)

	imd.Push(pixel.V(float64(startX), float64(startY)))
	imd.Push(pixel.V(float64(startX), float64(startY)))

	imd.Line(10)

	imd.Draw(win)
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

func (t *Trip) StartTrip() {
	t.Color = colornames.Cyan
}

func (t *Trip) Finish() {
	t.Color = colornames.Green
}

func (t *Trip) WarnEarly() {
	t.Color = colornames.Yellow
}

func NewTrip(id int, a, b, x, y, s, f int32) (t *Trip) {
	trip := new(Trip)

	trip.ID = id

	trip.SetStart(a, b)
	trip.SetEnd(x, y)
	trip.EarliestStart = s
	trip.LatestFinish = f

	// default values
	trip.InProgress = false
	trip.Taken = false
	trip.Color = config.Config.UI.TripDefaultColor

	return trip
}
