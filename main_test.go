package main

import "testing"

func TestBoundsOne(t *testing.T) {
	values := []float64{100, 0, 23, 54, 16, 34}
	min, max, stepY, stepX, width := determineBounds(values, 10, 10)

	ExpectEqual(t, float64(0), min)
	ExpectEqual(t, float64(100), max)
	ExpectEqual(t, float64(10), stepY)
	ExpectEqual(t, 1, stepX)
	ExpectEqual(t, 6, width)
}

func TestBoundsTwo(t *testing.T) {
	values := []float64{99, 0, 23, 54, 16, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	min, max, stepY, stepX, width := determineBounds(values, 10, 10)

	ExpectEqual(t, float64(0), min)
	ExpectEqual(t, float64(99), max)
	ExpectEqual(t, float64(9.9), stepY)
	ExpectEqual(t, 2, stepX)
	ExpectEqual(t, 10, width)
}

func TestGetValue(t *testing.T) {
	values := []float64{99, 0, 23, 54, 16, 34, 12, 16, 73, 82, 53, 32, 78, 11, 1}

	testParams := []struct {
		arg1 int
		arg2 int
		res  float64
	}{
		{1, 0, float64(99)},
		{1, 4, float64(16)},
		{2, 0, float64(49.5)},
		{2, 1, float64(38.5)},
		{4, 1, float64(19.5)},
		{4, 3, float64(30)},
	}

	for _, p := range testParams {
		actual := getValue(values, p.arg1, p.arg2)
		ExpectEqual(t, p.res, actual)
	}
}

func ExpectEqual(t *testing.T, e interface{}, a interface{}) {
	if e != a {
		FailWith(t, e, a)
	}
}

func FailWith(t *testing.T, e interface{}, a interface{}) {
	t.Errorf("Expected %v, but got %v", e, a)
}
