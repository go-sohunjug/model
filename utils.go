package model

import "math"

func Trunc(f float64, n int) float64 {
	p := math.Pow(10, float64(n))
	return math.Trunc(f*p) / p
}
