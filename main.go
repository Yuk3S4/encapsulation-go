package main

import (
	"github.com/Yuk3S4/encapsulation-go/course"
)

func main() {
	Go := &course.Course{
		Name:    "Go desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
	Go.ChangePrice(24)
}
