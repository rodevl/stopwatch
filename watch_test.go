package stopwatch_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/rodevl/stopwatch"
)

func withNow(fn func() time.Time, callback func()) {
	oldNow := stopwatch.Now
	defer func() {
		stopwatch.Now = oldNow
	}()

	stopwatch.Now = fn
	callback()
}

func withNowOffset(t time.Duration, callback func()) {
	fn := func() time.Time {
		return time.Now().Add(t)
	}

	withNow(fn, callback)
}

func TestStopWatchString(t *testing.T) {
	exp := `^30\.(\d+)ms$`
	rexp := regexp.MustCompile(exp)

	var watch stopwatch.Watch

	withNowOffset(-30*time.Millisecond, func() {
		watch = stopwatch.Start()
	})

	watch.Stop()

	// We're not millisecond accurate above, so...
	if !rexp.MatchString(watch.String()) {
		t.Fatalf("expected `%s` to match `%s`", watch, exp)
	}
}

func TestDeferring(t *testing.T) {
	exp := `^30m0\.\d+s$`
	rexp := regexp.MustCompile(exp)

	var called bool

	defer func() {
		if !called {
			t.Fatalf("failed to call defered function")
		}
	}()

	var watch stopwatch.Watch

	// Rewind the clock by 30 minutes so we have a realistic value to check this
	// against.
	withNowOffset(-30*time.Minute, func() {
		watch = stopwatch.Start()
	})

	defer watch.Timer(func(w stopwatch.Watch) {
		called = true

		if !rexp.MatchString(w.String()) {
			t.Fatalf("expected `%s` to match `%s`", watch, exp)
		}
	})
}

func TestWatch_CurrentDurationMillisecond(t *testing.T) {

	var watch stopwatch.Watch

	withNowOffset(-30*time.Millisecond, func() {
		watch = stopwatch.Start()
	})

	defer watch.Stop()
	currentDuration := watch.CurrentDurationMillisecond()
	// We're not millisecond accurate above, so...
	if currentDuration != 30 {
		t.Fatalf("expected `%s` to match `%s`", "30", "30")
	}
}

func TestWatch_CurrentDurationSecond(t *testing.T) {
	var watch stopwatch.Watch

	withNowOffset(-35*time.Second, func() {
		watch = stopwatch.Start()
	})

	defer watch.Stop()
	currentDuration := watch.CurrentDurationSecond()
	if currentDuration != 35 {
		t.Fatalf("expected 35 to match `%d`", currentDuration)
	}
}

func ExampleWatch_Timer() {
	defer stopwatch.StartAt(time.Now().Add(-30 * time.Minute)).Timer(func(w stopwatch.Watch) {
		fmt.Printf("elapsed time: %d minutes", w.Minutes())
	})

	// Output:
	// elapsed time: 30 minutes
}
