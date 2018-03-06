package simulator

import (
	ui "github.com/AkselsLedins/google-hashcode-2018-live-simulation/ui"
	"github.com/faiface/pixel/pixelgl"
)

func (s *Simulation) HandleEvents(win *pixelgl.Window, dt float64) {
	if win.JustPressed(pixelgl.KeySpace) {
		s.Toggle()
	}

	if win.Pressed(pixelgl.KeyLeft) {
		ui.Cam().CamPos.X -= ui.Cam().CamSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		ui.Cam().CamPos.X += ui.Cam().CamSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		ui.Cam().CamPos.Y -= ui.Cam().CamSpeed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		ui.Cam().CamPos.Y += ui.Cam().CamSpeed * dt
	}
}
