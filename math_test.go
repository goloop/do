package do

import (
	"math"
	"reflect"
	"testing"
)

// TestAbs tests the Abs function.
func TestAbs(t *testing.T) {
	// Test positive integer.
	n1 := 5
	expected1 := 5
	result1 := Abs(n1)
	if result1 != expected1 {
		t.Errorf("Abs of positive integer: expected %v, but got %v",
			expected1, result1)
	}

	// Test negative integer.
	n2 := -8
	expected2 := 8
	result2 := Abs(n2)
	if result2 != expected2 {
		t.Errorf("Abs of negative integer: expected %v, but got %v",
			expected2, result2)
	}

	// Test positive floating-point number.
	n3 := 3.14
	expected3 := 3.14
	result3 := Abs(n3)
	if result3 != expected3 {
		t.Errorf("Psitive floating-point number: expected %v, but got %v",
			expected3, result3)
	}

	// Test negative floating-point number.
	n4 := -2.718
	expected4 := 2.718
	result4 := Abs(n4)
	if result4 != expected4 {
		t.Errorf("Negative floating-point number: expected %v, but got %v",
			expected4, result4)
	}

	// Test zero value.
	n5 := 0
	expected5 := 0
	result5 := Abs(n5)
	if result5 != expected5 {
		t.Errorf("Abs of zero value: expected %v, but got %v",
			expected5, result5)
	}
}

func TestAbs_Generic(t *testing.T) {
	// Test positive integer.
	n1 := 5
	expected1 := 5
	result1 := Abs(n1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Abs of positive integer (G): expected %v, but got %v",
			expected1, result1)
	}

	// Test negative integer.
	n2 := -8
	expected2 := 8
	result2 := Abs(n2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Abs of negative integer (G): expected %v, but got %v",
			expected2, result2)
	}

	// Test positive floating-point number.
	n3 := 3.14
	expected3 := 3.14
	result3 := Abs(n3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Positive floating-point number (G): expected %v, but got %v",
			expected3, result3)
	}

	// Test negative floating-point number.
	n4 := -2.718
	expected4 := 2.718
	result4 := Abs(n4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Negative floating-point number (G): expected %v, but got %v",
			expected4, result4)
	}

	// Test zero value.
	n5 := 0
	expected5 := 0
	result5 := Abs(n5)
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf("Abs of zero value (G): expected %v, but got %v",
			expected5, result5)
	}
}

// TestAverage tests the Average function.
func TestAverage(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		want float64
	}{
		{"Numbers including zero", []int{1, 0, 3}, 1.3333333333333333},
		{"Non-zero numbers", []int{2, 2, 6}, 3.3333333333333335},
		{"Numbers including negative", []int{-1, 0, 3}, 0.6666666666666666},
		{"All zeros", []int{0, 0, 0}, 0},
		{"Single element", []int{5}, 5},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Average(tt.v...); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}

	// Empty slice
	if got := Average[int](); got != 0 {
		t.Errorf("Average() = %v, want 0", got)
	}
}

// TestMedian calculates the median of a list of integers.
func TestMedian(t *testing.T) {
	// Test case 1: Odd number of values.
	numbers1 := []int{3, 5, 7, 1, 9, 2}
	median1 := Median(numbers1...)
	expected1 := 4.0
	if median1 != expected1 {
		t.Errorf("Median of %v: expected %f, but got %f",
			numbers1, expected1, median1)
	}

	// Test case 2: Even number of values.
	numbers2 := []int{4, 6, 2, 8}
	median2 := Median(numbers2...)
	expected2 := 5.0
	if median2 != expected2 {
		t.Errorf("Median of %v: expected %f, but got %f",
			numbers2, expected2, median2)
	}

	// Test case 3: Empty slice.
	var numbers3 []int
	median3 := Median(numbers3...)
	expected3 := 0.0
	if median3 != expected3 {
		t.Errorf("Median of %v: expected %f, but got %f",
			numbers3, expected3, median3)
	}

	// Test case with odd number of values
	//valuesOdd := []int{3, 5, 7, 1, 9, 2}
	//expectedOdd := 4.0
	//resultOdd := Median(valuesOdd...)
	//if resultOdd != expectedOdd {
	//	t.Errorf("Expected median: %f, but got: %f", expectedOdd, resultOdd)
	//}
	valuesOdd := []int{1, 3, 5, 7, 9}
	expectedOdd := 5.0
	resultOdd := Median(valuesOdd...)
	if resultOdd != expectedOdd {
		t.Errorf("Expected median: %f, but got: %f", expectedOdd, resultOdd)
	}

	// Test case with even number of values
	valuesEven := []int{1, 2, 3, 4}
	expectedEven := 2.5
	resultEven := Median(valuesEven...)
	if resultEven != expectedEven {
		t.Errorf("Expected median: %f, but got: %f", expectedEven, resultEven)
	}

	// Test case with empty input
	valuesEmpty := []int{}
	expectedEmpty := 0.0
	resultEmpty := Median(valuesEmpty...)
	if resultEmpty != expectedEmpty {
		t.Errorf("Expected median: %f, but got: %f", expectedEmpty, resultEmpty)
	}
}

// TestMax tests the Max function.
func TestMax(t *testing.T) {
	tests := []struct {
		name string
		v    int
		more []int
		want int
	}{
		{
			name: "All positive numbers",
			v:    5,
			more: []int{7, 2, 4, 9},
			want: 9,
		},
		{
			name: "Includes negative numbers",
			v:    0,
			more: []int{-7, 2, -3, 9},
			want: 9,
		},
		{
			name: "Single number",
			v:    1,
			more: []int{},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.v, tt.more...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMaxList tests the MaxList function.
func TestMaxList(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name     string
		list     []int
		defaults []int
		want     int
	}{
		{
			name:     "Non-empty list, no defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: nil,
			want:     9,
		},
		{
			name:     "Empty list, no defaults",
			list:     []int{},
			defaults: nil,
			want:     0,
		},
		{
			name:     "Empty list, with defaults",
			list:     []int{},
			defaults: []int{20, 10},
			want:     20,
		},
		{
			name:     "Non-empty list with defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: []int{20, 10},
			want:     9,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the MaxList function.
			got := MaxList(tt.list, tt.defaults...)

			// Check if the result matches the expected value.
			if got != tt.want {
				t.Errorf("MaxList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMin tests the Min function.
func TestMin(t *testing.T) {
	tests := []struct {
		name string
		v    int
		more []int
		want int
	}{
		{
			name: "All positive numbers",
			v:    5,
			more: []int{7, 2, 4, 9},
			want: 2,
		},
		{
			name: "Includes negative numbers",
			v:    0,
			more: []int{-7, 2, -3, 9},
			want: -7,
		},
		{
			name: "Single number",
			v:    1,
			more: []int{},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.v, tt.more...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMinList tests the MinList function.
func TestMinList(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name     string
		list     []int
		defaults []int
		want     int
	}{
		{
			name:     "Non-empty list, no defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: nil,
			want:     1,
		},
		{
			name:     "Empty list, no defaults",
			list:     []int{},
			defaults: nil,
			want:     0,
		},
		{
			name:     "Empty list, with defaults",
			list:     []int{},
			defaults: []int{20, 10},
			want:     10,
		},
		{
			name:     "Non-empty list with defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: []int{20, 10},
			want:     1,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the MinList function.
			got := MinList(tt.list, tt.defaults...)

			// Check if the result matches the expected value.
			if got != tt.want {
				t.Errorf("MinList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name string
		v    []int
		want int
	}{
		{"Numbers including zero", []int{1, 0, 3}, 4},
		{"Non-zero numbers", []int{2, 2, 6}, 10},
		{"Numbers including negative", []int{-1, 0, 3}, 2},
		{"All zeros", []int{0, 0, 0}, 0},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Sum(tt.v...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}

	// Empty slice.
	if got := Sum[int](); got != 0 {
		t.Errorf("Sum() = %v, want %v", got, 0)
	}
}

// TestIsEven tests the IsEven function.
func TestIsEven(t *testing.T) {
	// Test cases for integer values.
	// Expected results: true if the value is even, false otherwise.
	tests := []struct {
		value    float64
		expected bool
	}{
		{4, true},
		{3, false},
		{-6, true},
		{-5, false},
	}

	for _, test := range tests {
		result := IsEven(test.value)
		if result != test.expected {
			t.Errorf("IsEven(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values.
	// Expected results: true if the integer part is even, false otherwise.
	testsFloat := []struct {
		value    float64
		expected bool
		f        []bool
	}{
		{4.2, false, []bool{}},
		{4.2, true, []bool{true}},
		{3.8, false, []bool{true, false}},
		{-5.5, false, []bool{}},
		{-6.5, true, []bool{true}},
	}

	for _, test := range testsFloat {
		result := IsEven(test.value, test.f...)
		if result != test.expected {
			t.Errorf("IsEven(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

func TestIsOdd(t *testing.T) {
	// Test cases for integer values.
	// Expected results: true if the value is odd, false otherwise.
	tests := []struct {
		value    int
		expected bool
	}{
		{4, false},   // even number
		{3, true},    // odd number
		{-6, false},  // negative even number
		{-5, true},   // negative odd number
		{-10, false}, // negative even number
	}

	for _, test := range tests {
		result := IsOdd(test.value)
		if result != test.expected {
			t.Errorf("IsOdd(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values.
	// Expected results: true if the integer part is odd, false otherwise.
	testsFloat := []struct {
		value    float64
		expected bool
		f        []bool
	}{
		{3.2, false, []bool{}},
		{3.2, true, []bool{true}},
		{3.0, true, []bool{}},
		{-5.5, true, []bool{true}},
		{-6.5, false, []bool{true, false}},
		{-11.0, true, []bool{false}},
	}

	for _, test := range testsFloat {
		result := IsOdd(test.value, test.f...)
		if result != test.expected {
			t.Errorf("IsOdd(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

// TestIsWhole tests the IsWhole function.
func TestIsWhole(t *testing.T) {
	// Test cases for integer values
	// Expected results: true if the value has no fractional part,
	// false otherwise
	tests := []struct {
		value    int
		expected bool
	}{
		{4, true},
		{3, true},
		{-6, true},
		{-5, true},
		{-10, true},
		{0, true},
		{100, true},
		{5, true},
		{5, true},
		{-2, true},
	}

	for _, test := range tests {
		result := IsWhole(test.value)
		if result != test.expected {
			t.Errorf("IsWhole(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values
	// Expected results: true if the value has no fractional part,
	// false otherwise
	testsFloat := []struct {
		value    float64
		expected bool
	}{
		{4.2, false},
		{3.8, false},
		{-5.5, false},
		{-11.0, true},
		{0.0, true},
		{100.0, true},
		{5.0, true},
		{-2.0, true},
		{-2.1, false},
	}

	for _, test := range testsFloat {
		result := IsWhole(test.value)
		if result != test.expected {
			t.Errorf("IsWhole(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}
