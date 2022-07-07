package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	expected := 40.0
	length := 10.0
	breadth := 10.0
	rectangle := Rectangle{length, breadth}
	result := Perimeter(rectangle)

	if expected != result {
		// * "f" is for our float64
		// * ".2" means print 2 decimal places
		t.Errorf("expected %.2f got %.2f", expected, result)
	}
}

// func TestArea(t *testing.T) {

// 	t.Run("area of Rectangle", func(t *testing.T) {
// 		expected := 100.0
// 		length := 10.0
// 		breadth := 10.0
// 		rectangle := Rectangle{length, breadth}
// 		result := rectangle.Area()

// 		if expected != result {
// 			t.Errorf("expected %.2f got %.2f", expected, result)
// 		}
// 	})

// 	t.Run("area of Circle", func(t *testing.T) {
// 		expected := 314.1592653589793
// 		radius := 10.0
// 		circle := Circle{radius}
// 		result := circle.Area()

// 		if expected != result {
// 			// * "g" will print a more precise decimal number in the error message
// 			t.Errorf("expected %g got %g", expected, result)
// 		}
// 	})
// }

// // * refactored
// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, expected float64) {
// 		t.Helper()
// 		result := shape.Area()
// 		if expected != result {
// 			// * "g" will print a more precise decimal number in the error message
// 			t.Errorf("expected %g got %g", expected, result)
// 		}
// 	}

// 	t.Run("area of Rectangle", func(t *testing.T) {
// 		expected := 100.0
// 		length := 10.0
// 		breadth := 10.0
// 		rectangle := Rectangle{length, breadth}

// 		checkArea(t, rectangle, expected)
// 	})

// 	t.Run("area of Circle", func(t *testing.T) {
// 		expected := 314.1592653589793
// 		radius := 10.0
// 		circle := Circle{radius}

// 		checkArea(t, circle, expected)
// 	})
// }

// * refactored using table driven test
// * useful when you want to build a list of test cases that can be tested in the same manner
func TestArea(t *testing.T) {
	expectedRectangle := 100.0
	length := 10.0
	breadth := 10.0
	rectangle := Rectangle{length, breadth}

	expectedCircle := 314.1592653589793
	radius := 10.0
	circle := Circle{radius}

	expectedTriangle := 25.0
	base := 5.0
	height := 10.0
	triangle := Triangle{base, height}

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{rectangle, expectedRectangle},
		{circle, expectedCircle},
		{triangle, expectedTriangle},
	}

	for _, test := range areaTests {
		result := test.shape.Area()
		expected := test.expected
		if expected != result {
			// * %#v format string will print out our struct with the values in its field
			t.Errorf("%#v expected %g got %g", test.shape, expected, result)
		}
	}
}
