/* This is a set of test cases for the 'tempconv' package. It tests all
six functions included in the package and returns the results when run.
To run the tests, simply type 'go test' in terminal. If everything works
fine, you will get an output of 'PASS' and 'ok', along with the directory
the test was executed. If it fails, it will print the type of the error
it encountered. */

package tempconv

import (
	"testing"
)

func TestCToF(t *testing.T) {
	c := Celsius(100)
	expected := Fahrenheit(212)
	result := CToF(c)
	if result != expected {
		t.Errorf("CToF(%v) = %v, expected %v", c, result, expected)
	}
}

func TestFToC(t *testing.T) {
	f := Fahrenheit(32)
	expected := Celsius(0)
	result := FToC(f)
	if result != expected {
		t.Errorf("FToC(%v) = %v, expected %v", f, result, expected)
	}
}

func TestKToC(t *testing.T) {
	k := Kelvin(273.15)
	expected := Celsius(0)
	result := KToC(k)
	if result != expected {
		t.Errorf("KToC(%v) = %v, expected %v", k, result, expected)
	}
}

func TestCToK(t *testing.T) {
	c := Celsius(100)
	expected := Kelvin(373.15)
	result := CToK(c)
	if result != expected {
		t.Errorf("CToK(%v) = %v, expected %v", c, result, expected)
	}
}

func TestKToF(t *testing.T) {
	k := Kelvin(273.15)
	expected := Fahrenheit(32)
	result := KToF(k)
	if result != expected {
		t.Errorf("KToF(%v) = %v, expected %v", k, result, expected)
	}
}

func TestFToK(t *testing.T) {
	f := Fahrenheit(32)
	expected := Kelvin(273.15)
	result := FToK(f)
	if result != expected {
		t.Errorf("FToK(%v) = %v, expected %v", f, result, expected)
	}
}
