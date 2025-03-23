package test

import (
	"fmt"
	"math"
	"testing"
)

func TestDayOfWeek(t *testing.T) {
	type testSet struct {
		input int8
		want  string
	}

	testSets := []testSet{
		{input: 1, want: "Monday"},
		{input: 2, want: "Tuesday"},
		{input: 3, want: "Wednesday"},
		{input: 4, want: "Thursday"},
		{input: 5, want: "Friday"},
		{input: 6, want: "Saturday"},
		{input: 7, want: "Sunday"},
	}

	for _, set := range testSets {
		r := DayOfWeek(set.input)
		if r != set.want {
			t.Errorf("DayOfWeek(%d) = %s, want %s", set.input, r, set.want)
		}

	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(float64(i))
	}
}

var sum float64

func BenchmarkSquare2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum += math.Pow(float64(i), 2)
	}
	fmt.Println(sum)
}
