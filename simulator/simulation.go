package simulator

import (
	ghashcode "github.com/AkselsLedins/google-hashcode-2018-live-simulation/ghashcode"

	"github.com/faiface/pixel/imdraw"
)

// Simulation is a struct handling information about the simulation
type Simulation struct {
	Stopped bool
	Ended   bool

	LastStep int
	Step     int
	Score    int

	Trips    []*ghashcode.Trip
	Vehicles []*ghashcode.Vehicle

	// configuration
	gridRows                     int16
	gridColumns                  int16
	numberOfVehicles             int16
	numberOfRides                int16
	perRideBonus                 int16
	numberOfStepsInTheSimulation int
}

func (s *Simulation) Run(imd *imdraw.IMDraw) {
	if imd != nil {
		for _, trip := range s.Trips {
			trip.AddToImd(imd)
		}
	}

	remainingVehicles := 0
	for _, vehicle := range s.Vehicles {
		if !vehicle.Enabled {
			continue
		}
		remainingVehicles++
		if imd != nil {
			vehicle.AddToImd(imd)
		}
		if s.LastStep != s.Step && !s.Stopped {
			s.Score += vehicle.Drive(s.Trips, s.Step, s.perRideBonus)
		}
	}

	if !s.Stopped {
		s.Step++
	}

	// or the simulation reaches its final step T
	if (s.Step >= int(s.numberOfStepsInTheSimulation) && s.numberOfStepsInTheSimulation != 0) || remainingVehicles == 0 {
		s.Ended = true
		s.Stop()
	}
}

func (s *Simulation) Toggle() {
	s.Stopped = !s.Stopped
}

func (s *Simulation) Stop() {
	s.Stopped = true
}

func (s *Simulation) Start() {
	s.Stopped = false
}

// NewSimulation a constructor for the simulation
// which set the default values
func NewSimulation() *Simulation {
	s := new(Simulation)

	s.LastStep = -1
	s.Step = 0
	s.Stopped = true
	s.Ended = false

	return s
}
