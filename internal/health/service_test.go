package health

import "testing"

type testCase struct {
	testName string
}

var testCases = []testCase{
	{
		testName: "TestCheck",
	},
}

func TestCheck(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			service := NewService()
			result := service.Check()

			if result != "OK" {
				t.Errorf("Expected result to be 'OK', got '%s'", result)
			}
		})
	}
}
