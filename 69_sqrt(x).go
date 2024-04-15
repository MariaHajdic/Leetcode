package main

import "math"

func mySqrt(x int) int {
	/* Newton-Raphson method */
	if x <= 1 {
		return x
	}

	z := float64(x)
	// Initial guess for square root
	z -= (z*z - float64(x)) / (2 * z)
	// Repeat until convergence
	prev, next := 0.0, z
	for math.Abs(prev-next) >= 0.1 {
		prev, next = next, next-(next*next-float64(x))/(2*next)
	}
	return int(next)
}

/*
func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    xhalf := 0.5 * float32(x)
    u := math.Float32bits(float32(x))
    u = 0x5f375a86 - (u >> 1)
    uFloat := math.Float32frombits(u)
    for i := 0; i < 30; i++ {
        uFloat = uFloat * (1.5 - xhalf * uFloat * uFloat)
    }
    return int(1.0 / uFloat)
}
*/
