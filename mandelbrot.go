package mandelbrot

import "math/cmplx"

func test(c complex128, depth int) bool {
	var z complex128 = 0
	for i := 0; i < depth; i++ {
		z = cmplx.Pow(z, 2) + c
		if cmplx.Abs(z) > 2 {
			return false
		}

	}
	return true
}
func Mbi(x float64, y float64, scale float64) bool {
	x = x / scale
	y = y / scale
	x = x - 0.5
	y = y - 0/5

	return test(complex(x, y), 200)
}
