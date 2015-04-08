package main

import "github.com/s2607/mandelbrot"
import "fmt"

func main() {
	var scale int = 70
	for i := 0; i < scale/1; i++ {
		for rel := 0; rel < scale/1; rel++ {
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
