package clock

import "fmt"

const dayInMinutes = 24 * 60

// Clock is a clock with only hours and minutes
type Clock struct {
	hour   int
	minute int
}

// New create a new Clock
func New(h, m int) Clock {
	clock := &Clock{hour: h, minute: m}
	clock.normalize()
	return *clock
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds the given amount in minutes to the Clock and
// returns a new Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract reduces the given amount in minutes from the Clock and
// returns a new Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}

func (c *Clock) normalize() {
	c.minute = c.hour*60 + c.minute
	negative := false
	if c.minute < 0 {
		negative = true
		c.minute = -c.minute
	}

	if c.minute >= dayInMinutes {
		c.minute -= dayInMinutes * (c.minute / dayInMinutes)
	}

	if negative {
		c.minute = dayInMinutes - c.minute
	}

	c.hour = c.minute / 60
	c.minute = c.minute % 60

	return
}
