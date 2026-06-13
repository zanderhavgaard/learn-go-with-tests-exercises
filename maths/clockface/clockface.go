// Package clockface
package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	secondHandPoint := Point{x, y}
	return secondHandPoint
}

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func SecondHand(t time.Time) Point {
	point := secondHandPoint(t)
	point = Point{point.X * secondHandLength, point.Y * secondHandLength} // scale
	point = Point{point.X, -point.Y}                                      // flip
	point = Point{point.X + clockCentreX, point.Y + clockCentreY}         // translate
	return point
}
