package ui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

func DrawStepNumber(win *pixelgl.Window, step int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(920, 700), basicAtlas)

	fmt.Fprintf(basicTxt, "Step : %06d", step)

	basicTxt.Draw(win, pixel.IM)
}
