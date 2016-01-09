package main

import "github.com/s2607/mandelbrot"
import "fmt"
import "math"

func main() {
	//1181
	//559
	//274
	var scale int = 284
	//var scale int = 18
	//var scale int = 30
	//	for i := -scale * 2; i < scale*2; i++ {
	for x := -scale * 2; x < scale*2; x = x + 2 {
		if scale < 50 {
			fmt.Print(x / 10)
		} else {
			fmt.Print(x / 100)
		}
		if x > -10 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
	for x := -scale * 2; x < scale*2; x = x + 1 {
		if scale < 50 {
			fmt.Print(int(math.Abs(float64(x))) % 10)
		} else {
			fmt.Print(int(math.Abs(float64(x))) % 100)
		}
	}
	fmt.Print("\n")
	for i := -scale * 2; i < scale*2; i = i + 2 {
		for rel := -scale * 2; rel < scale*2; rel++ {
			if mandelbrot.Mbi(float64(rel), float64(i), float64(scale)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("|%d\n", i)
	}
	fmt.Println("done")
}
