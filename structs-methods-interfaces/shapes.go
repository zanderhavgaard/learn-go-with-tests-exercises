package main

import "math"

// "A struct is just a named collection of fileds where you can store data"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	// method
	area := math.Pi * c.Radius * c.Radius
	return area
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	area := (t.Base * t.Height) * 0.5
	return area
}

type Shape interface {
	// interface
	Area() float64 // type must have a method with this signature
}

func Perimeter(rectangle Rectangle) float64 {
	// function
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
