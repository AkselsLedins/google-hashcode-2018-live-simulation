package ghashcode

import (
	"image/color"
	"math"

	"golang.org/x/image/colornames"

	config "../config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Trip a structure which represents a trip in the simulation
// Quite self-explanatory
type Trip struct {
	// unique identifier for each trip, used for debugging
	ID int

	// start & end coordinates
	Start Coordinates
	End   Coordinates

	// computed distance from Start to End
	// helps for score computation
	Distance int
	// we set there the bonus defined in Config if the
	// driver starts its trip on time
	Bonus int

	// A driver cannot start its trip before EarlierStart step
	// And if he finish his drive after LatestFinish the trip will
	// be marked as failed
	EarliestStart int32
	LatestFinish  int32

	// current color of the trip
	Color color.RGBA

	// when a driver is going towards the trip
	Taken bool
	// the vehicle is actually doing it
	InProgress bool
	// the vehicle end the trip too late
	Failed bool
}

// AddToImd adds to the imd batch the graphic line of the trip
// We only call .draw once for all the trips
func (t *Trip) AddToImd(imd *imdraw.IMDraw) {
	// we only show taken trips for performance
	// TODO a way to desactivate this
	if !t.Taken {
		return
	}

	// depending of the status of the trip we assign it a color
	imd.Color = t.Color

	// start point
	startX := t.Start.X*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	startY := t.Start.Y*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(startX), float64(startY)))
	// second point
	intermediateX := (t.End.X)*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	intermediateY := (t.Start.Y)*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(intermediateX), float64(intermediateY)))
	// third final point
	endX := t.End.X*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	endY := t.End.Y*config.Config.UI.SquareSize + config.Config.UI.SquareSize
	imd.Push(pixel.V(float64(endX), float64(endY)))

	// registering the line
	imd.Line(2)

	// we create the tip of the trip to see where its start is
	// looks like this:       O---------
	imd.Color = t.Color
	imd.EndShape = imdraw.RoundEndShape

	imd.Push(pixel.V(float64(startX), float64(startY)))
	imd.Push(pixel.V(float64(startX), float64(startY)))

	imd.Line(10)
}

// SetStart setter for StartCoordinates
func (t *Trip) SetStart(x, y int32) {
	t.Start.X = x
	t.Start.Y = y
}

// SetEnd setter for StartCoordinates
func (t *Trip) SetEnd(x, y int32) {
	t.End.X = x
	t.End.Y = y
}

// SomeoneIsOnIt we could store the vehicle id for debugging purposes there
// but atm we only change the trip color and its status
func (t *Trip) SomeoneIsOnIt() {
	t.Color = colornames.Beige
	t.Taken = true
}

// StartTrip takes the current step as parameter
// if we start it on time we earn the bonus points
// a different color is assigned if we start it on time or not
func (t *Trip) StartTrip(step int, bonus int16) {
	if step == int(t.EarliestStart) {
		t.Bonus += int(bonus)
		t.Color = colornames.Gold
		return
	}
	t.Color = colornames.Cyan
}

// Finish the vehicle finished the trip at step.
// we determine if he failed or not to arrive on time
// a different is assigned if we end it on time or not
func (t *Trip) Finish(step int32) int {
	failed := false
	// the vehicle does this even if the arrival step is later than the latest finish
	if step > t.LatestFinish {
		failed = true
	}
	t.Failed = failed

	if !failed {
		t.Color = colornames.Green
		return t.Distance + t.Bonus
	}
	// but no points are earned by such a ride
	t.Color = colornames.Red
	return 0
}

// WarnEarly it changes only the trip color to warn graphically
// that a vehicle is too early on a trip
// We can see that way that some drivers are waiting way too long
func (t *Trip) WarnEarly() {
	t.Color = colornames.Yellow
}

// NewTrip a constructor for Trip
func NewTrip(id int, a, b, x, y, s, f int32) *Trip {
	trip := new(Trip)

	trip.ID = id

	trip.SetStart(a, b)
	trip.SetEnd(x, y)
	trip.EarliestStart = s
	trip.LatestFinish = f

	// default values
	trip.InProgress = false
	trip.Taken = false
	trip.Failed = false
	trip.Color = config.Config.UI.TripDefaultColor

	// precomputed values
	trip.Distance = int(math.Abs(float64(a-x)) + math.Abs(float64(b-y)))
	trip.Bonus = 0
	imd := imdraw.New(nil)

	imd.Color = trip.Color
	imd.EndShape = imdraw.RoundEndShape

	return trip
}
