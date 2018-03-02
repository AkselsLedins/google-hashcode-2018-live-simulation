package main

import (
	"flag"
	"fmt"
	"time"

	ghashcode "./ghashcode"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	size       *int
	windowSize *float64
	frameRate  *time.Duration

	frames = 0
	second = time.Tick(time.Second)
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

func init() {
	frameRate = flag.Duration("frameRate", 33*time.Millisecond, "The framerate in milliseconds")
	flag.Parse()
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
	vehicle := new(ghashcode.Vehicle)
	vehicle.SetPosition(1, 3)

	trip := new(ghashcode.Trip)
	trip.Start = ghashcode.Coordinates{50, 50}
	trip.End = ghashcode.Coordinates{20, 80}

	tick := time.Tick(*frameRate)
	for !win.Closed() {
		// logic loop
		frames++
		select {
		case <-tick:
			grid.Draw(win)
			vehicle.DrawToWindow(win)
			trip.DrawToWindow(win)
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
