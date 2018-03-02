package main

import (
	ghashcode "./ghashcode"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func createGrid(sizeX, sizeY int32) *imdraw.IMDraw {
	imd := imdraw.New(nil)

	imd.Color = colornames.Blueviolet
	imd.EndShape = imdraw.RoundEndShape
	squareSize := int32(5)
	for x := int32(0); x < sizeX; x++ {
		for y := int32(0); y < sizeY; y++ {
			offsetX := x*squareSize + squareSize
			offsetY := y*squareSize + squareSize
			imd.Push(pixel.V(float64(x+offsetX), float64(y+offsetY)))
			imd.Push(pixel.V(float64(x+squareSize+offsetX), float64(y+squareSize+offsetY)))
			imd.Rectangle(2)
		}
	}

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

	grid := createGrid(100, 100)
	driver := new(ghashcode.Vehicle)

	for !win.Closed() {
		grid.Draw(win)
		driver.DrawToWindow(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
