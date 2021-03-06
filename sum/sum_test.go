package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int
	n2       int
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{118, 1, 119},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_uint32 = []struct {
	n1		 uint32
	n2 		 uint32
	expected uint32
}{
	{1, 2, 3},
	{4, 5, 9},
	{4294967300, 220, 4294967520},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_int32 = []struct {
	n1		 int32
	n2		 int32
	expected int32
}{
	{1, 2, 3},
	{4, 5, 9},
	{3000000000, 1, 3000000001},
}

var sum_tests_int64 = []struct {
	n1		 int64
	n2		 int64
	expected int64
}{
	{1, 2, 3},
	{4, 5, 9},
	{50000000000000000001, 1, 50000000000000000002},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_float64 = []struct {
	n1		 float64
	n2		 float64
	expected float64
}{
	{1, 2, 3},
	{4, 5, 9},
	{120, 1, 119},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%v, %v) returned %v, expected %v", v.n1, v.n2, val, v.expected)
		}
	}
}