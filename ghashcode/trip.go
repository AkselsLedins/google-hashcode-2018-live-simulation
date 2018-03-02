package ghashcode

import (
	config "../config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Trip struct {
	Start Coordinates
	End   Coordinates

	EarliestStart int32
	LatestFinish  int32
}

func (t *Trip) DrawToWindow(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	imd.Color = config.Config.UI.TripDefaultColor
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

	imd.Line(1)

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
