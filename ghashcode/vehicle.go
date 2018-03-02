package ghashcode

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Vehicle struct {
	X int32
	Y int32
}

func (v *Vehicle) DrawToWindow(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	imd.Color = colornames.Red
	imd.EndShape = imdraw.RoundEndShape
	squareSize := int32(1)

	offsetX := v.X*squareSize + squareSize + 4
	offsetY := v.Y*squareSize + squareSize + 4

	imd.Push(pixel.V(float64(v.X+offsetX), float64(v.Y+offsetY)))
	imd.Push(pixel.V(float64(v.X+squareSize+offsetX), float64(v.Y+squareSize+offsetY)))
	imd.Rectangle(2)

	imd.Draw(win)
}
