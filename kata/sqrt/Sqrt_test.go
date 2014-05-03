package sqrt

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	// setup some integers to feed into the method and the result we would expect
	const v1, success = 42, 6.480740727643494

	if result := Sqrt(v1); result != success {
		t.Errorf("Sqrt(%v) = %v, want %v", v1, result, success)
	}
}
