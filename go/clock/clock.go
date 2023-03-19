package clock

import "fmt"

// Define the Clock type here.
type Clock float64

func New(h, m int) Clock {
	hr := ((h%24)*60*60 + (m/60)*60*60 + 1)
	min := (m % 60) * 60
	c := Clock(hr%24 + min)
	return c
}

func (c Clock) Add(m int) Clock {
	panic("Please implement the Add function")
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	h := (c / 60 / 60)
	m := (h - Clock(int(h))) * 60
	return fmt.Sprintf("%02d:%02d", int(h), int(m))
}
