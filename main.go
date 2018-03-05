package main

import (
	"flag"
	"fmt"
	"math"
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

	frames = 0
	second = time.Tick(time.Second)

	outputFile *string
	inputFile  *string

	camPos       = pixel.ZV
	camSpeed     = 500.0
	camZoom      = 1.0
	camZoomSpeed = 1.2
)

func init() {
	outputFile = flag.String("o", "", "Path to your result")
	inputFile = flag.String("i", "", "Path to the exercice input")
	flag.Parse()
}

func run() {
	vehicles := simulator.ParseOutputFile(*outputFile)
	trips := simulator.ParseInputFile(*inputFile)

	cfg := pixelgl.WindowConfig{
		Title:  config.Config.UI.WindowTitle,
		Bounds: pixel.R(0, 0, 1024, 720),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tick := time.Tick(6 * time.Millisecond)

	score := 0
	step := 0
	lastStep := 0
	imd := imdraw.New(nil)
	last := time.Now()
	for !win.Closed() {
		imd.Clear()
		frames++
		dt := time.Since(last).Seconds()
		last = time.Now()

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		step++
		select {
		case <-tick:
			win.Clear(colornames.Black)
			for _, trip := range trips {
				trip.AddToImd(imd)
			}
			for _, vehicle := range vehicles {
				if !vehicle.Enabled {
					continue
				}
				if lastStep != step {
					score += vehicle.Drive(trips, step)
				}
				vehicle.AddToImd(imd)
			}
			imd.Draw(win)

			if lastStep != step {
				lastStep = step
			}
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		}

		win.SetMatrix(pixel.IM)
		ui.DrawStepNumber(win, step)
		ui.DrawScore(win, score)
		ui.DrawNumberOfVehicles(win, len(vehicles))
		ui.DrawNumberOfTrips(win, len(trips))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
