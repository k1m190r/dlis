package dlis

// Val is "universal value"
type Val struct {
	// payload with a value
	s *string
	i *int
	f *float64
	v *Val

	c int // count
	e error
}
