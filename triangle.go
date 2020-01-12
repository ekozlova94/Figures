package main

import "math"

type triangle struct {
	a float64
	b float64
	c float64
}

//Определение типа треугольника
func (t triangle) tip() string {
	if t.a == t.b && t.a == t.c {
		return "equilateral triangle"
	} else if t.a == t.b && t.a != t.c || t.a == t.c && t.a != t.b || t.b == t.c && t.a != t.b {
		return "isosceles triangle"
	} else if t.c*t.c == t.a*t.a+t.b*t.b {
		return "right triangle"
	} else {
		return "arbitrary triangle"
	}
}

//Периметр треугольника
func (t triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

//Площадь треугольника
func (t triangle) area() float64 {
	p := t.perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
