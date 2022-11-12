package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Add (%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 3, b: 1, want: 2},
		{a: 2, b: 2, want: 0},
		{a: 5, b: 6, want: -1},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f) : want %f, got %f", tc.a, tc.a, tc.want, got)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 3, b: 3, want: 9},
		{a: 2, b: 2, want: 4},
		{a: 0, b: 5, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)

		}

	}

}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Error("want error for invalid input, got nil")
	}

}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSquareRoot(t *testing.T) {
	t.Parallel()

	type testCase struct {
		num  float64
		want float64
	}

	testCases := []testCase{
		{num: 4, want: 2},
		{num: 16, want: 4},
		{num: 8, want: 2.82},
		{num: 2.5, want: 1.58},
	}

	for _, tc := range testCases {
		got, err := calculator.SquareRoot(tc.num)
		if err != nil {
			t.Fatalf("want no fatal error. Got :%v", err)
		}

		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("SquareRoot(%f), want %f, got %f", tc.num, tc.want, got)
		}
	}

}

func TestSquareRootInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.SquareRoot(-2)
	if err == nil {
		t.Error("Wanted error for invalid input, got nil")

	}
}
