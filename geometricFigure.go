package main

type selector func(geometricFigure) float64

type geometricFigure interface {
	perimeter() float64
	area() float64
	tip() string
}
