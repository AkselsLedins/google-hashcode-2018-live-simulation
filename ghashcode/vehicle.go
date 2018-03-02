package ghashcode

import (
	config "../config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Vehicle struct {
	CurrentPosition Coordinates
}

func (v *Vehicle) DrawToWindow(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	imd.Color = config.Config.UI.VehicleDefaultColor
	imd.EndShape = imdraw.RoundEndShape
	squareSize := int32(5)

	offsetX := v.CurrentPosition.X*squareSize + squareSize
	offsetY := v.CurrentPosition.Y*squareSize + squareSize

	drawX := float64(v.CurrentPosition.X + offsetX)
	drawY := float64(v.CurrentPosition.Y + offsetY)
	imd.Push(pixel.V(drawX, drawY))
	imd.Push(pixel.V(drawX, drawY))

	imd.Line(2)

	imd.Draw(win)
}

func (v *Vehicle) SetPosition(x, y int32) {
	v.CurrentPosition.X = x
	v.CurrentPosition.Y = y
}
