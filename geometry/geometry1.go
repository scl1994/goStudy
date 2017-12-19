package geometry

import (
	"math"
	"fmt"
)

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (Point) ShowPoint(q Point) {
	//接收者可以没有名字
	fmt.Printf("{%f, %f}\n", q.X, q.Y)
}

func (p *Point) ShowPoint2() {
	//接收者可以是一个指向命名类型的指针
	//不能使*a.X，他的意思是*(a.X)
	fmt.Printf("{%f, %f}\n", (*p).X, p.Y)
	p.X++
	p.Y++
}