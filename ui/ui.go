package ui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

func DrawScore(win *pixelgl.Window, score int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(900, 685), basicAtlas)

	fmt.Fprintf(txt, "Score : %09d", score)

	txt.Draw(win, pixel.IM)
}

func DrawStepNumber(win *pixelgl.Window, step int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(900, 700), basicAtlas)

	fmt.Fprintf(txt, "Step  :    %06d", step)

	txt.Draw(win, pixel.IM)
}

func DrawNumberOfVehicles(win *pixelgl.Window, numberOfVehicles int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(830, 670), basicAtlas)

	fmt.Fprintf(txt, "Number of cars  :      %04d", numberOfVehicles)

	txt.Draw(win, pixel.IM)
}

func DrawNumberOfTrips(win *pixelgl.Window, numberOfTrips int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(830, 655), basicAtlas)

	fmt.Fprintf(txt, "Number of trips :      %04d", numberOfTrips)

	txt.Draw(win, pixel.IM)
}

func DrawStartHint(win *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(340, 30), basicAtlas)

	fmt.Fprintf(txt, "  Press [T] to display all the trips (at your own risk)\n")
	fmt.Fprintf(txt, " [Arrow Keys] to move around    [Scroll] to Zoom in/out \n")
	fmt.Fprintf(txt, "        Press [SPACE] to start/pause simulation")

	txt.Draw(win, pixel.IM)
}
