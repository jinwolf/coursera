package sortlib

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"sorted slice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"Unsorted slice", []int{5, 2, 3}, []int{2, 3, 5}},
		{"Unsorted slice", []int{5, 2, 3, 10, 25, 12, 33, -1}, []int{-1, 2, 3, 5, 10, 12, 25, 33}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeSort(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("TestMergeSort(%v): got %v, want %v",tc.input,got,tc.expected)
			}
		})
	}

}

func TestSwap(t *testing.T) {
	testCases := []struct {
		name string
		a, b int // initial values
		expA, expB int // expected values after swap
	}{
		{"positive numbers", 1, 2, 2, 1},
		{"Negative numbers", -1, -2, -2, -1},
		{"Zero and positive", 0, 3, 3, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, b := tc.a, tc.b 
			Swap(&a, &b) 

			if a != tc.expA || b != tc.expB {
				t.Errorf("Swap(%d, %d): expected (%d, %d), got (%d, %d)", tc.a, tc.b, tc.expA, tc.expB, a, b)
			}
		})
	}
}

func TestMergeSlices(t *testing.T) {
	testCases := []struct {
		name string
		a, b []int
		merged []int
	}{
		{"Single element", []int{2}, []int{1}, []int{1,2}},
		{"Two element", []int{2,4}, []int{1,5}, []int{1,2,4,5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			merged := MergeSlices(tc.a, tc.b)

			if !reflect.DeepEqual(merged, tc.merged) {
				t.Errorf("TestMergedSlices: expected %v, got %v",tc.merged, merged)
			}
		})
	}
}
/*
func TestMergeInPlace(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		sorted []int
	}{
		{"two elements", []int{3, 2}, []int{2, 3}},
		{"two elements", []int{3, -2}, []int{-2, 3}},
		{"two elements", []int{-1, -2}, []int{-2, -1}},
		{"three elements", []int{3, -1, -2}, []int{-2, -1, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			MergeInPlace(tc.input)

			if !reflect.DeepEqual(tc.input, tc.sorted) {
				t.Errorf("TestMergedInPlace: expected %v, got %v", tc.sorted, tc.input)
			}
		})
	}
}
*/


