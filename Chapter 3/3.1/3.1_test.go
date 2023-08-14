// This is a simple test to check if the 'isValidFunction' works properly. It
// checks for valid polygon points (1st case), 'NaN' values (2nd case) and
// '±Inf' values (3rd case).

package main

import (
	"math"
	"testing"
)

func TestIsValidPolygon(t *testing.T) {

	if !isValidPolygon(10.0, 20.0) {
		t.Error("Expected (10.0, 20.0) to be a valid polygon point")
	}

	if isValidPolygon(math.NaN(), 5) {
		t.Error("Expected NaN value not to be a valid polygon point")
	}

	if isValidPolygon(10, math.Inf(1)) {
		t.Error("Expected Inf value not to be a valid polygon point")
	}
}
