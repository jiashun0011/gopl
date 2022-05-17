package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i - 1].Distance(path[i])
		}
	}
	return sum
}

// func Distance(p, q Point) float64 {
// 	return math.Hypot(q.X-p.X, q.Y-p.Y)
// }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

/* func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
} */

/* func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p)
} */

/* func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
} */

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	
	path := Path{p, q}
	fmt.Println(path)

	path.TranslateBy(Point{1, 1}, true)
	fmt.Println(path)

	path.TranslateBy(Point{1, 1}, false)
	fmt.Println(path)
}
