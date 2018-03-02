package ghashcode

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Coordinates struct {
	X, Y int32
}

type Trip struct {
	Start Coordinates
	End   Coordinates
}

func (t *Trip) DrawToWindow(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	imd.Color = colornames.Orange
	imd.EndShape = imdraw.RoundEndShape

	/* start point */
	startX := t.Start.X*5 + 5
	startY := t.Start.Y*5 + 5
	imd.Push(pixel.V(float64(startX), float64(startY)))
	/* second point */
	x := (t.End.X)*5 + 5
	y := (t.Start.Y)*5 + 5
	imd.Push(pixel.V(float64(x), float64(y)))
	/* final point */
	endX := t.End.X*5 + 5
	endY := t.End.Y*5 + 5
	imd.Push(pixel.V(float64(endX), float64(endY)))

	// fmt.Printf("0) %v %v\n", startX, startY)
	// fmt.Printf("1) %v %v\n", x, y)
	// fmt.Printf("2) %v %v\n", endX, endY)
	// fmt.Printf("\n")

	imd.Line(2)
	/*

	   imd.Push(pixel.V(float64(v.X+squareSize+offsetX), float64(v.Y+squareSize+offsetY)))
	   imd.Rectangle(2)
	*/

	imd.Draw(win)
}
