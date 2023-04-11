package main

import (
	"fmt"
	"math"
)

type SubVertex struct {
	Z float64
}

func (s SubVertex) Scale(f float64) {
	s.Z = s.Z * f
}

func (s *SubVertex) Scale2(f float64) {
	s.Z = s.Z * f
}

type Vertex struct {
	X, Y float64
	Z    *SubVertex
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale readonly
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale2 read write
func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) SubScale(f float64) {
	v.Z.Scale(f)
	v.Z = nil
}

func (v Vertex) SubScale2(f float64) {
	v.Z.Scale2(f)
	v.Z = nil
}

func (v Vertex) GetSub() float64 {
	return v.Z.Z
}

func main() {
	v1 := Vertex{3, 4, nil}
	v1.Scale(10)
	fmt.Println(v1.Abs())

	v2 := Vertex{3, 4, nil}
	v2.Scale2(10)
	fmt.Println(v2.Abs())

	v3 := Vertex{0, 0, &SubVertex{3}}
	v3.SubScale(10)
	fmt.Println(v3.GetSub())

	v4 := Vertex{0, 0, &SubVertex{3}}
	v4.SubScale2(10)
	fmt.Println(v4.GetSub())
}
