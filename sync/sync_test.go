package main

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	gotValue := got.Value()

	if gotValue != want {
		t.Errorf("got %d, want %d", gotValue, want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times should leave it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func TestCounterConcurrently(t *testing.T) {
	t.Run("counter runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}
