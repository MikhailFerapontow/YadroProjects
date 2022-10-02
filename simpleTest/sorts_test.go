package simpletest

import "testing"

func TestBubbleSort(t *testing.T) {
	// Test values
	arr := []int{5, 1, 2, 3, 0, 4}
	expected := []int{0, 1, 2, 3, 4, 5}

	// Act
	result := BubbleSort(arr)

	if !Equal(result, expected) {
		t.Error("Sorting went wrong")
	}
}

// First test is not very informative
// and does not cover different situations

func TestMergeSort(t *testing.T) {
	// Test values
	testTable := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{0, 5, 4, 3, 2, 1, 9, 7, 8, 6},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			arr:      []int{0, 1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1, 0},
		},
	}

	// Act
	for _, arrays := range testTable {
		result := MergeSort(arrays.arr)
		t.Logf("Calling (%v), result %v", arrays.arr, result)

		if !Equal(result, arrays.expected) {
			t.Errorf("Error. Expected %v, get %v", arrays.expected, result)
		}
	}
}
