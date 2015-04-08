package mandelbrot

import "math/cmplx"
import "fmt"

func S() string {
	var z, c complex128 = 0, 1
	for i := 0; i < 10; i++ {
		z = cmplx.Pow(z, 2) + c
	}

	return fmt.Sprintf("a %v", z)

}
func Four() int {
	return 4
}
