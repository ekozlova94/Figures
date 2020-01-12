package main

import (
	"fmt"
	"math"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	counter()
}

//Треугольники
func counter() {
	var figures = []geometricFigure{
		triangle{1, 1, 1},
		triangle{2, 2, 2},
		triangle{1, 0, 1},
		triangle{3, 4, 5},
		triangle{1, 2, 3},
		quadrangle{1, 1, 1, 1},
		quadrangle{2, 2, 2, 2},
		quadrangle{1, 0, 1, 0},
		quadrangle{3, 4, 4, 3},
		quadrangle{1, 2, 3, 4},
	}
	var m = make(map[string][]geometricFigure)
	for i := 0; i < len(figures); i++ {
		m[figures[i].tip()] = append(m[figures[i].tip()], figures[i])
	}
	//fmt.Println(m)
	perimeter := func(f geometricFigure) float64 { return f.perimeter() }
	area := func(f geometricFigure) float64 { return f.area() }

	for key, value := range m {
		resultMaxPerimeter := max(perimeter, value)
		resultMaxArea := max(area, value)
		resultMinPerimeter := min(perimeter, value)
		resultMinArea := min(area, value)
		fmt.Println(key, "наибольшая площадь:", resultMaxArea, "наибольший периметр:", resultMaxPerimeter)
		fmt.Println(key, "наименьшая площадь:", resultMinArea, "наименьший периметр:", resultMinPerimeter)
	}
}

/*для примера
func area(f geometricFigure) float64 {
	return f.area()
}*/

//Наибольшая площадь и периметр без if для треугольника
/*func max(ahaha []geometricFigure) (float64, float64) {
	var maxArea, maxPerimeter float64 = ahaha[0].area(), ahaha[0].perimeter()
	for i := 0; i < len(ahaha); i++ {
		maxArea = math.Max(ahaha[i].area(), maxArea)
		maxPerimeter = math.Max(ahaha[i].perimeter(), maxPerimeter)
	}
	return maxArea, maxPerimeter
}*/
func max(selector func(geometricFigure) float64, ahaha []geometricFigure) float64 {
	var max = selector(ahaha[0])
	for i := 0; i < len(ahaha); i++ {
		max = math.Max(selector(ahaha[i]), max)
	}
	return max
}

// Наименьшая площадь и периметр с if для треугольника
func min(s selector, ahaha []geometricFigure) float64 {
	var min = s(ahaha[0])
	for i := 0; i < len(ahaha); i++ {
		if currArea := s(ahaha[i]); currArea <= min {
			min = currArea
		}
	}
	return min
}
