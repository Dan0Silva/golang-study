package shape

import "testing"

type testCase struct {
	shape          Shape
	expectedResult float64
}

func TestArea(t *testing.T) {
	t.Run("testing rectangle area", func(t *testing.T) {
		testCases := []testCase{
			{Rectangle{10, 15}, 150.00},
			{Rectangle{10, 10}, 100.00},
			{Rectangle{6, 12}, 72.00},
		}

		for _, thisCase := range testCases {
			if result := thisCase.shape.Area(); result != thisCase.expectedResult {
				t.Errorf("invalid calc [%f, %f]", result, thisCase.expectedResult)
			}
		}
	})
	t.Run("testing circle area", func(t *testing.T) {
		testCases := []testCase{
			{Circle{10}, 314.16},
			{Circle{5}, 78.54},
			{Circle{0}, 0.00},
		}

		for _, thisCase := range testCases {
			if result := thisCase.shape.Area(); result != thisCase.expectedResult {
				t.Errorf("invalid calc [%f, %f]", result, thisCase.expectedResult)
			}
		}
	})
}
