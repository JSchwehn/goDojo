package gcd

import (
	"testing"
)

func TestGcdEuclid(t *testing.T) {
	// setup some integers to feed into the method and the result we would expect
	const v1, v2, result = 84, 30, 6

	if x := GcdEuclid(v1, v2); x != 6 {
		t.Errorf("GcdEuclid(%v, %v) = %v, want %v", v1, v2, x, result)
	}
}
func TestGcdEuclidModernIterativ(t *testing.T) {
	// setup some integers to feed into the method and the result we would expect
	const v1, v2, result = 84, 30, 6

	if x := GcdEuclidModernIterative(v1, v2); x != 6 {
		t.Errorf("GcdEuclid(%v, %v) = %v, want %v", v1, v2, x, result)
	}
}
func TestGcdEuclidModernRecursiv(t *testing.T) {
	// setup some integers to feed into the method and the result we would expect
	const v1, v2, result = 84, 30, 6

	if x := GcdEuclidModernRecursive(v1, v2); x != 6 {
		t.Errorf("GcdEuclid(%v, %v) = %v, want %v", v1, v2, x, result)
	}
}
