package main

import "fmt"

type printTriangle interface {
	triangleLeft() string
	triangleRight() string
}

type triangle struct {
	value int
}

func (t *triangle) triangleLeft() string {
	print := ""
	for i := 0; i < t.value; i++ {
		printLeft := ""
		for j := 0; j <= i; j++ {
			printLeft += "*"
		}
		fmt.Println(printLeft)
	}
	return print
}

func (t *triangle) triangleRight() string {
	print := ""
	for i := 0; i < t.value; i++ {
		printRight := ""
		for j := 0; j < t.value; j++ {
			if j >= i {
				printRight += "*"
			} else {
				printRight += " "
			}
		}
		fmt.Println(printRight)
	}
	return print
}

func main() {
	var printTriangle printTriangle = &triangle{7}

	fmt.Println(printTriangle.triangleLeft())
	fmt.Println(printTriangle.triangleRight())
}
