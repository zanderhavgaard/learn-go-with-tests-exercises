package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      string = "Go!"
	countdownStart int    = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeperDuration := 1 * time.Second
	sleeperFunction := time.Sleep
	sleeper := &ConfigurableSleeper{sleeperDuration, sleeperFunction}
	Countdown(os.Stdout, sleeper)
}
