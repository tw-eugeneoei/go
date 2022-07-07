package structsmethodsinterfaces

import "math"

func Perimeter(shape Rectangle) float64 {
	return (shape.Length + shape.Breadth) * 2
}

// func Area(shape Rectangle) float64 {
// 	return shape.Length * shape.Breadth
// }

// * In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

// * syntax for declaring methods on types
// * func (receiverName ReceiverType) MethodName(args) {...}
// * It is a convention in Go to have the receiver variable be the first letter of the type.
func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
