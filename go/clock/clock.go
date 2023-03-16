package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hr  int
	min int
}

// h%24               => will give the rounded off hour
// h%24 + m/60        => will give the roll over from minutes if its greater than 60 mins, else 0
// (h%24 + m/60) % 24 => will give roll over for hour if h + min/60 > 24
func New(h, m int) Clock {
	return Clock{
		hr:  hourRollOver(h, m),
		min: minRollOver(m),
	}
}

func hourRollOver(h, m int) (hr int) {
	hr = ((h % 24) + (m / 60)) % 24
	fmt.Println((h % 24), (m / 60), (h%24)+(m/60), ((h%24)+(m/60))%24, hr)

	if hr < 0 {
		hr = (24 - hr + (m / 60)) % 24
		fmt.Println(hr, m)

		if hr < 0 {
			hr *= -1
		}
		fmt.Println(hr, m)
		if m < 0 {
			hr++
		}
	} else if m < 0 {
		hr = hr + (m / 60) - 1
	}
	fmt.Println(hr)
	return
}

func minRollOver(m int) (min int) {
	min = m % 60
	if min < 0 {
		min = 60 + min
	}
	return
}

func (c Clock) Add(m int) Clock {

	// set init hour as current hour
	hr, m := addMins(c, m)

	// add remaining mins to c.min and reset hr with calculated hour
	c.min = c.min % 60
	c.hr = hr
	return c
}

func addHour(c Clock, h int) int {
	return h % 24
}

func addMins(c Clock, m int) (int, int) {

	hr := c.hr
	// until current min + m > 60, add 1 to hour
	// and reduce 60 mins from m
	for c.min+m > 60 {
		m -= 60
		hr++

		// if hr > 23, reset to 0
		if hr > 23 {
			hr = 0
		}
	}
	return hr, m
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hr, c.min)
}
