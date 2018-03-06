package main

import (
	"flag"
	"fmt"
	"time"

	config "github.com/AkselsLedins/google-hashcode-2018-live-simulation/config"
	simulator "github.com/AkselsLedins/google-hashcode-2018-live-simulation/simulator"
	ui "github.com/AkselsLedins/google-hashcode-2018-live-simulation/ui"

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

	noGuiFlag *bool

	outputFile *string
	inputFile  *string

	simulation *simulator.Simulation
)

func init() {
	simulation = simulator.NewSimulation()

	outputFile = flag.String("o", "", "Path to your result")
	inputFile = flag.String("i", "", "Path to the exercice input")
	noGuiFlag = flag.Bool("noGui", false, "Run the simulation without a GUI")

	flag.Parse()

	simulation.ParseOutputFile(*outputFile)
	simulation.ParseInputFile(*inputFile)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  config.Config.UI.WindowTitle,
		Bounds: pixel.R(0, 0, 1024, 720),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tick := time.Tick(6 * time.Millisecond)

	imd := imdraw.New(nil)
	last := time.Now()

	for !win.Closed() && !simulation.Ended {
		imd.Clear()
		frames++
		dt := time.Since(last).Seconds()
		last = time.Now()

		// set the camera for the scene
		win.SetMatrix(ui.Cam().GetMatrix(win))

		// handle user inputs
		simulation.HandleEvents(win, dt)

		select {
		case <-tick:
			win.Clear(colornames.Black)

			simulation.Run(imd)

			imd.Draw(win)

		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		}

		// reset the matrix
		win.SetMatrix(pixel.IM)

		// drall the UI
		ui.DrawStepNumber(win, simulation.Step)
		ui.DrawScore(win, simulation.Score)
		ui.DrawNumberOfVehicles(win, len(simulation.Vehicles))
		ui.DrawNumberOfTrips(win, len(simulation.Trips))
		ui.DrawStartHint(win)

		win.Update()
	}

	fmt.Printf("Score: %d\n", simulation.Score)
}

func noGui() {
	simulation.Start()
	for !simulation.Ended {
		simulation.Run(nil)
	}
	fmt.Printf("%d\n", simulation.Score)
}

func main() {
	if *noGuiFlag == true {
		noGui()
		return
	}
	pixelgl.Run(run)
}
