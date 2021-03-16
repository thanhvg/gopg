package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{0.8, 0.2}
	got := Perimeter(rec)
	wanted := 2.00
	if got != wanted {
		t.Errorf("Expect %.2f got %.2f", wanted, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0}}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
