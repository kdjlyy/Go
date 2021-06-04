package main

import (
	"fmt"
	"math"
)

// 定义一个接口
type Shape interface {
	Area() float32
	Circumference() float32
}

type Rect struct {
	Width  float32
	Height float32
}
type Circle struct {
	Radius float32
}

func (r Rect) Area() float32 {
	return r.Width * r.Height
}
func (r Rect) Circumference() float32 {
	return 2 * (r.Width + r.Height)
}

func (r Circle) Area() float32 {
	return math.Pi * r.Radius * r.Radius
}
func (r Circle) Circumference() float32 {
	return 2 * math.Pi * r.Radius
}

func showInfo(shape Shape) {
	fmt.Printf("Area:%f Circumference:%f\n", shape.Area(), shape.Circumference())
}

// 在上面的 showInfo 中我们传入了接口类型的对象，如果将实现了接口的类型传递进去，
// 那么会将实际类型的其他特性掩盖住，因此通常我们会想要获取其真正的类型, 可以使用下面的方法:
func showInfo2(shape Shape) {
	switch shape.(type) {
	case Rect:
		fmt.Println("This is Rect")
	case Circle:
		fmt.Println("This is Circle")
	}
	fmt.Printf("Area:%f Circumference:%f\n", shape.Area(), shape.Circumference())
}

func main() {
	r := Rect{5, 10}
	fmt.Printf("Rect:\n\tWidth:%f \tHeight:%f\n\tArea:%f \tCircumference:%f\n",
		r.Width, r.Height, r.Area(), r.Circumference())

	c := Circle{10}
	fmt.Printf("Circle:\n\tRadius:%f\n\tArea:%f \tCircumference:%f\n",
		c.Radius, c.Area(), c.Circumference())

	showInfo2(r)
	showInfo2(c)
}
