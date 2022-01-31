package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type square struct {
	base   float64
	altura float64
}

func (s square) area() float64 {
	return s.base * s.altura
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println(s.(circle).area())
	case square:
		fmt.Println(s.(square).area())
	}
}

func main() {
	c := circle{
		radius: 2.0,
	}

	s := square{
		base:   2.0,
		altura: 7.0,
	}

	info(c)
	info(s)

}
