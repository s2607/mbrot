package main

import "github.com/s2607/mandelbrot"
import "fmt"

func main() {
	//1181
	//559
	//274
	var scale int = 274
	for i := -scale * 2; i < scale*2; i++ {
		for rel := -scale * 2; rel < scale*2; rel++ {
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
