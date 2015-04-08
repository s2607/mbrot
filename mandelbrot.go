package mandelbrot

import "math/cmplx"

func test(c complex128, depth int) bool {
	var z complex128 = 0
	for i := 0; i < depth; i++ {
		z = cmplx.Pow(z, 2) + c
		if real(z)+imag(z) > 2 {
			return false
		}

	}
	return true
}
func Mbi(x float64, y float64, scale float64) bool {
	x = x - (scale / 2)
	y = y - (scale / 2)
	x = x * (1 / scale)
	y = y * (1 / scale)

	return test(complex(x, y), 10)
}
