fastfunc
========

Eg.

	package main
	
	import (
		"fmt"
		"math"
	
		"github.com/JessonChan/fastfunc"
	)
	
	func main() {
		fastfunc.SetRunTimes(1e7)
		fmt.Println("func 1", fastfunc.Run(func() {
			for i := 0; i < 3; i++ {
				math.Pow(float64(i), 2)
			}
		})/1e6, "microsecond")
		fmt.Println("func 2", fastfunc.Run(func() {
		}), "nanosecond")
	}
