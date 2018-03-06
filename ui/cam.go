package ui

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type cam struct {
	CamPos       pixel.Vec
	CamSpeed     float64
	CamZoom      float64
	CamZoomSpeed float64
}

var (
	camera *cam
)

func init() {
	camera = new(cam)

	camera.CamPos = pixel.ZV
	camera.CamSpeed = 500.0
	camera.CamZoom = 1.0
	camera.CamZoomSpeed = 1.2
}

func Cam() *cam {
	return camera
}

func (c *cam) GetMatrix(win *pixelgl.Window) pixel.Matrix {
	c.CamZoom *= math.Pow(c.CamZoomSpeed, win.MouseScroll().Y)
	return pixel.IM.Scaled(c.CamPos, c.CamZoom).Moved(win.Bounds().Center().Sub(c.CamPos))
}
