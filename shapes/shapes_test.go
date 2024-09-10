package shapes

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Testing Perimeter of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		if got != want {
			t.Error(genErrorMessage(got, want))
		}
	})
	t.Run("Testing Perimeter of Circle", func(t *testing.T) {
		circle := Circle{5.0}
		got := circle.Perimeter()
		want := 31.42
		if got != want {
			t.Error(genErrorMessage(got, want))
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{5.0}, want: 78.54},
		{name: "Triangle", shape: Triangle{10.0, 10.0}, want: 50.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			checkArea(test.name, t, test.shape, want)
		})
	}
}

func checkArea(name string, t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("%s %#v got %v want %v", name, shape, got, want)
	}
}

func genErrorMessage(got, want float64) string {
	return fmt.Sprintf("got %v, want %v", got, want)
}
