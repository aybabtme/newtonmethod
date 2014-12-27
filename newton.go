// Package newtonmethod uses Newton's method to approximate
// roots to equations, using their derivative and a starting
// approximation.
package newtonmethod

type ƒ func(x float64) float64

// Apply Newton's method to find a root to ƒ(value), with
// initial approximate solution x, iterating n times.
func Apply(value float64, ƒ, dƒ ƒ, n int, x float64) float64 {

	if n == 0 {
		return x
	}

	return Apply(
		value,
		ƒ,
		dƒ,
		n-1,
		x-(ƒ(x)-value)/dƒ(x),
	)
}
