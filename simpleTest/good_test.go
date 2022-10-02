package simpletest

import "testing"

func Test_IsBad(t *testing.T) {
	// Test values
	str := "bad day"
	expected := true

	// Act
	result := isBad(str)

	//Output
	if result != expected {
		t.Error("Error. Function gives unexpected result")
	}
}

// First test is not very informative
// and does not cover different situations

func Test_IsGood(t *testing.T) {
	// Test values
	testTable := []struct {
		str      string
		expected bool
	}{
		{
			str:      "good Day",
			expected: true,
		},
		{
			str:      "very bad Day",
			expected: false,
		},
		{
			str:      "It would be good",
			expected: true,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := IsGood(testCase.str)
		t.Logf("Calling (%s), result %t", testCase.str, result)

		// Output
		if result != testCase.expected {
			t.Errorf("Error. Expected %t, get %t", testCase.expected, result)
		}
	}

}
