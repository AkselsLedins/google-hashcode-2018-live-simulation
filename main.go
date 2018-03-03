package main

import (
	"flag"
	"fmt"
	"time"

	config "./config"
	simulator "./simulator"
	ui "./ui"

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

	outputFile *string
	inputFile  *string
)

func createGrid(sizeX, sizeY int32) *imdraw.IMDraw {
	imd := imdraw.New(nil)

	imd.Color = config.Config.UI.GridColor
	imd.EndShape = imdraw.RoundEndShape
	squareSize := config.Config.UI.SquareSize
	for x := int32(0); x < sizeX; x++ {
		for y := int32(0); y < sizeY; y++ {
			offsetX := x*squareSize + squareSize
			offsetY := y*squareSize + squareSize
			imd.Push(pixel.V(float64(x+offsetX), float64(y+offsetY)))
			imd.Push(pixel.V(float64(x+squareSize+offsetX), float64(y+squareSize+offsetY)))
			imd.Rectangle(1)
		}
	}

	return imd
}

func init() {
	frameRate = flag.Duration("frameRate", 1*time.Millisecond, "The framerate in milliseconds")
	outputFile = flag.String("o", "", "Path to your result")
	inputFile = flag.String("i", "", "Path to the exercice input")
	flag.Parse()
}

func run() {
	vehicles := simulator.ParseOutputFile(*outputFile)
	trips := simulator.ParseInputFile(*inputFile)
	fmt.Printf("%v", trips[0])

	cfg := pixelgl.WindowConfig{
		Title:  config.Config.UI.WindowTitle,
		Bounds: pixel.R(0, 0, 1024, 720),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// grid := createGrid(100, 100)

	tick := time.Tick(*frameRate)

	step := 0
	lastStep := 0
	for !win.Closed() {
		// logic loop
		frames++
		select {
		case <-tick:
			win.Clear(colornames.Black)
			// grid.Draw(win)
			// fmt.Printf("STEP (%d)\n", step)
			if win.JustPressed(pixelgl.KeyRight) {
			}
			step++
			for _, trip := range trips {
				trip.DrawToWindow(win)
			}
			for _, vehicle := range vehicles {
				if !vehicle.Enabled {
					continue
				}
				if lastStep != step {
					vehicle.Drive(trips)
				}
				vehicle.DrawToWindow(win)
			}
			if lastStep != step {
				lastStep = step
			}
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		}
		ui.DrawStepNumber(win, step)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
