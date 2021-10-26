package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printCircleArea(22)
	// fmt.Println("Plosha kryga s radiusom 5cm = ", calcCircleArea(5))
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Radius kryga:", radius)
	fmt.Printf("Formula A=pi*r*r\n")
	fmt.Printf("Площа круга: %f cm, kv\n", circleArea)
}
func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Radius cannot be negative")
	}
	return float32(radius) * float32(radius) * math.Pi, nil
}
