package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout time.Duration = 10 * time.Second

func ping(url string) chan struct{} {
	// we struct{} as the return type b/c it does not allocate anything in memory
	// and we don't care about the return type here
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string, error error) {
	return ConfigureableRacer(a, b, tenSecondTimeout)
}

func ConfigureableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// when using select with channels here, the first channel to return something gets selected
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for urls %q, and %q", a, b)
	}
}
