package main

import "fmt"

const pi = 3.14

func main() {
	printCircleArea()
}

func printCircleArea() {
	circleRadius := 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Println("Radius:", circleRadius)
	fmt.Printf("Formula A=pi*r*r\n")

	fmt.Println("Area = ", circleArea)
}
