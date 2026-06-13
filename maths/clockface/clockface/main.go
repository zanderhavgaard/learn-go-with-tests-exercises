package main

import (
	"os"
	"time"

	"github.com/zanderhavgaard/learn-go-with-tests-exercises/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
