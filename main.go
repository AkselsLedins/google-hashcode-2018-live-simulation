package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func createGrid(x, y int32) *imdraw.IMDraw {
	imd := imdraw.New(nil)

	imd.Color = colornames.Blueviolet
	imd.EndShape = imdraw.RoundEndShape
	/* create outside rect */
	imd.Push(pixel.V(0+50, 0+50), pixel.V(0+50, 768-50))
	imd.Push(pixel.V(1024-50, 768-50))
	imd.Push(pixel.V(1024-50, 0+50))
	imd.Push(pixel.V(0+50, 0+50))
	imd.Line(1)

	return imd
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Google Hashcode 2018 - Simulator!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		grid := createGrid(100, 100)
		grid.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
