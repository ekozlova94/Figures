package main

import "math"

type quadrangle struct {
	a float64
	b float64
	c float64
	d float64
}

func (q quadrangle) tip() string {
	if q.a == q.b && q.b == q.c && q.c == q.d {
		return "square quadrangle"
	} else if q.a == q.b && q.b != q.c && q.c == q.d || q.a == q.d && q.a != q.b && q.b == q.c || q.a == q.c && q.a != q.b && q.b == q.d {
		return "rectangle quadrangle"
	} else {
		return "random quadrangle"
	}
}

//Периметр четырёхугольника
func (q quadrangle) perimeter() float64 {
	return q.a + q.b + q.c + q.d
}

//Площадь четырёхугольника
func (q quadrangle) area() float64 {
	p := q.perimeter() / 2
	return math.Sqrt(p * (p - q.a) * (p - q.b) * (p - q.c) * (p - q.d))
}
