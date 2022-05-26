package stopwatch

import (
	"fmt"
	"time"
)

//TimerFunc Function prototype for timers.
type TimerFunc func(Watch)

type Watch interface {
	fmt.Stringer

	// Timer calls a callback with the currently-measured time. This is useful for
	// deferring out of a function.
	Timer(fn TimerFunc)

	//Stop  stops the watch based on the current wall-clock time.
	Stop() Watch

	//Start starts the watch based on the current wall-clock time.
	Start() Watch

	//Duration returns the elapsed duration
	Duration() time.Duration

	//CurrentDuration returns the current duration
	CurrentDuration() time.Duration

	// Milliseconds returns the elapsed duration in milliseconds.
	Milliseconds() int

	// Seconds returns the elapsed duration in seconds.
	Seconds() int

	// Minutes returns the elapsed duration in minutes.
	Minutes() int

	// Hours returns the elapsed duration in hours.
	Hours() int

	// Days returns the elapsed duration in days.
	Days() int

	// CurrentDurationSecond returns the current duration in seconds
	CurrentDurationSecond() int

	// CurrentDurationMillisecond returns the current duration in milliseconds
	CurrentDurationMillisecond() int
}

var Now = func() time.Time {
	return time.Now()
}

type watch struct {
	start, stop time.Time
}

// Timer calls a callback with the currently-measured time. This is useful for
// deferring out of a function.
func (s *watch) Timer(fn TimerFunc) {
	fn(s.Stop())
}

//Stop  stops the watch based on the current wall-clock time.
func (s *watch) Stop() Watch {
	s.stop = Now()
	return s
}

//Start starts the watch based on the current wall-clock time.
func (s *watch) Start() Watch {
	s.start = Now()
	return s
}

// String returns a human-readable representation of the stopwatch's duration.
func (s *watch) String() string {
	// if the watch isn't stopped yet...
	if s.stop.IsZero() {
		return "0m0.00s"
	}

	return s.duration().String()
}

//duration returns the elapsed duration
func (s *watch) duration() time.Duration {
	return s.stop.Sub(s.start)
}

//CurrentDuration returns the current duration
func (s *watch) CurrentDuration() time.Duration {
	timeNow := time.Now()
	duration := timeNow.Sub(s.start)
	return duration
}

// Duration returns the elapsed duration.
func (s *watch) Duration() time.Duration {
	return s.duration()
}

// Milliseconds returns the elapsed duration in milliseconds.
func (s *watch) Milliseconds() int {
	return int(s.duration().Milliseconds())
}

// Seconds returns the elapsed duration in seconds.
func (s *watch) Seconds() int {
	return int(s.duration().Seconds())
}

// Minutes returns the elapsed duration in minutes.
func (s *watch) Minutes() int {
	return int(s.duration().Minutes())
}

// Hours returns the elapsed duration in hours.
func (s *watch) Hours() int {
	return int(s.duration().Hours())
}

// Days returns the elapsed duration in days.
func (s *watch) Days() int {
	return int(s.duration().Hours()) / 24
}

// CurrentDurationSecond returns the current duration in seconds
func (s *watch) CurrentDurationSecond() int {
	return int(s.CurrentDuration() / time.Second)
}

// CurrentDurationMillisecond returns the current duration in milliseconds
func (s *watch) CurrentDurationMillisecond() int {
	return int(s.CurrentDuration() / time.Millisecond)
}
