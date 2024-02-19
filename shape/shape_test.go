package shape

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape       Shape
		want        float64
		description string
	}{
		{Rectangle{10.0, 10.0}, 40.0, "perimeter of rectangle"},
		{Circle{10.0}, 62.83185307179586, "perimeter of circle"},
		{Triangle{12, 6}, 24.0, "perimeter of triangle"},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("Perform %s: %#v got %.2f want %.2f", tt.description, tt.shape, got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape       Shape
		want        float64
		description string
	}{
		{Rectangle{12, 6}, 72.0, "area of rectangle"},
		{Circle{10}, 314.1592653589793, "area of circle"},
		{Triangle{12, 6}, 36.0, "area of triangle"},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Perform %s: %#v got %.2f want %.2f", tt.description, tt.shape, got, tt.want)
		}
	}
}
