package ui

type options struct {
	ShowAllTrips bool
}

var (
	opt *options
)

func init() {
	opt = new(options)
	opt.ShowAllTrips = false
}

func Options() *options {
	return opt
}
