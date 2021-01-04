package scheduler

import "time"

// // Interval defines the
// type Interval uint32

// // Interval type
// const (
// 	Minute Interval = iota
// 	Hour
// 	Day
// 	Week
// 	Month
// 	Year
// )

// ClockOption define the behaviors of clock.
type ClockOption func(c *Clock)

// Clock is a type of job that work at some time.
type Clock struct {
	job
	state   State
	repeat  bool
	daily   []daily
	weekly  []weekly
	monthly []monthly
	yearly  []yearly
}

type daily struct {
	hour   uint8
	minute uint8
	second uint8
}

type weekly struct {
	day time.Weekday
}

type monthly struct {
	day     uint8
	reverse bool
}

type yearly struct {
	month time.Month
}

// Next .
func (c *Clock) Next() {
	// TODO
}

// State .
func (c *Clock) State() State {
	// TODO
	return c.state
}

func (c *Clock) check() error {
	// TODO
	return nil
}

// NewClock create a new clock job.
func NewClock(key string, f interface{}, options ...ClockOption) Job {
	clock := &Clock{
		job: job{
			key:      key,
			caller: f,
		},
	}

	for _, opt := range options {
		opt(clock)
	}

	return clock
}

// Daily .
func Daily(hour uint8, minute uint8, second uint8) ClockOption {
	return func(c *Clock) {
		c.daily = append(c.daily, daily{
			hour:   hour,
			minute: minute,
			second: second,
		})
	}
}

// Weekly .
func Weekly(day time.Weekday) ClockOption {
	return func(c *Clock) {
		c.weekly = append(c.weekly, weekly{
			day: day,
		})
	}
}

// Monthly .
func Monthly(day uint8, reserve ...bool) ClockOption {
	r := false
	if len(reserve) != 0 {
		r = reserve[0]
	}
	return func(c *Clock) {
		c.monthly = append(c.monthly, monthly{
			day:     day,
			reverse: r,
		})
	}
}

// Yearly .
func Yearly(month time.Month) ClockOption {
	return func(c *Clock) {
		c.yearly = append(c.yearly, yearly{
			month: month,
		})
	}
}

// Repeat default true
func Repeat(repeat ...bool) ClockOption {
	r := true
	if len(repeat) != 0 {
		r = repeat[0]
	}
	return func(c *Clock) {
		c.repeat = r
	}
}
