package main

import "github.com/s2607/mandelbrot"
import "fmt"

func main() {
	var scale int = 79
	for i := -scale; i < scale; i++ {
		for rel := -scale; rel < scale; rel++ {
			if mandelbrot.Mbi(float64(rel), float64(i), float64(scale)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("<")
	}
	fmt.Println("done")
}
