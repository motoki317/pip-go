package pip_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/motoki317/pip-go"
)

func TestPip(t *testing.T) {
	rectangle := pip.NewPolygon(
		[]pip.Point{
			{X: 1.0, Y: 1.0},
			{X: 1.0, Y: 2.0},
			{X: 2.0, Y: 2.0},
			{X: 2.0, Y: 1.0},
		},
	)

	pt1 := pip.Point{X: 1.1, Y: 1.1} // Should be true
	pt2 := pip.Point{X: 1.2, Y: 1.2} // Should be true
	pt3 := pip.Point{X: 1.3, Y: 1.3} // Should be true
	pt4 := pip.Point{X: 1.4, Y: 1.4} // Should be true
	pt5 := pip.Point{X: 1.5, Y: 1.5} // Should be true
	pt6 := pip.Point{X: 1.6, Y: 1.6} // Should be true
	pt7 := pip.Point{X: 1.7, Y: 1.7} // Should be true
	pt8 := pip.Point{X: 1.8, Y: 1.8} // Should be true

	pt9 := pip.Point{X: -4.9, Y: 1.2}    // Should be false
	pt10 := pip.Point{X: 10.0, Y: 10.0}  // Should be false
	pt11 := pip.Point{X: -5.0, Y: -6.0}  // Should be false
	pt12 := pip.Point{X: -13.0, Y: 1.0}  // Should be false
	pt13 := pip.Point{X: 4.9, Y: -1.2}   // Should be false
	pt14 := pip.Point{X: 10.0, Y: -10.0} // Should be false
	pt15 := pip.Point{X: 5.0, Y: 6.0}    // Should be false
	pt16 := pip.Point{X: -13.0, Y: 1.0}  // Should be false

	assert(rectangle.Contains(pt1), true, t)
	assert(rectangle.Contains(pt2), true, t)
	assert(rectangle.Contains(pt3), true, t)
	assert(rectangle.Contains(pt4), true, t)
	assert(rectangle.Contains(pt5), true, t)
	assert(rectangle.Contains(pt6), true, t)
	assert(rectangle.Contains(pt7), true, t)
	assert(rectangle.Contains(pt8), true, t)

	assert(rectangle.Contains(pt9), false, t)
	assert(rectangle.Contains(pt10), false, t)
	assert(rectangle.Contains(pt11), false, t)
	assert(rectangle.Contains(pt12), false, t)
	assert(rectangle.Contains(pt13), false, t)
	assert(rectangle.Contains(pt14), false, t)
	assert(rectangle.Contains(pt15), false, t)
	assert(rectangle.Contains(pt16), false, t)

	t.Log("Finished")
}

func BenchmarkPip(b *testing.B) {
	rectangle := pip.NewPolygon(
		[]pip.Point{
			{X: 1.0, Y: 1.0},
			{X: 1.0, Y: 2.0},
			{X: 2.0, Y: 2.0},
			{X: 2.0, Y: 1.0},
		},
	)

	pt1 := pip.Point{X: 1.1, Y: 1.1} // Should be true
	pt2 := pip.Point{X: 1.2, Y: 1.2} // Should be true
	pt3 := pip.Point{X: 1.3, Y: 1.3} // Should be true
	pt4 := pip.Point{X: 1.4, Y: 1.4} // Should be true
	pt5 := pip.Point{X: 1.5, Y: 1.5} // Should be true
	pt6 := pip.Point{X: 1.6, Y: 1.6} // Should be true
	pt7 := pip.Point{X: 1.7, Y: 1.7} // Should be true
	pt8 := pip.Point{X: 1.8, Y: 1.8} // Should be true

	pt9 := pip.Point{X: -4.9, Y: 1.2}    // Should be false
	pt10 := pip.Point{X: 10.0, Y: 10.0}  // Should be false
	pt11 := pip.Point{X: -5.0, Y: -6.0}  // Should be false
	pt12 := pip.Point{X: -13.0, Y: 1.0}  // Should be false
	pt13 := pip.Point{X: 4.9, Y: -1.2}   // Should be false
	pt14 := pip.Point{X: 10.0, Y: -10.0} // Should be false
	pt15 := pip.Point{X: 5.0, Y: 6.0}    // Should be false
	pt16 := pip.Point{X: -13.0, Y: 1.0}  // Should be false

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rectangle.Contains(pt1)
		rectangle.Contains(pt2)
		rectangle.Contains(pt3)
		rectangle.Contains(pt4)
		rectangle.Contains(pt5)
		rectangle.Contains(pt6)
		rectangle.Contains(pt7)
		rectangle.Contains(pt8)

		rectangle.Contains(pt9)
		rectangle.Contains(pt10)
		rectangle.Contains(pt11)
		rectangle.Contains(pt12)
		rectangle.Contains(pt13)
		rectangle.Contains(pt14)
		rectangle.Contains(pt15)
		rectangle.Contains(pt16)
	}
}

func BenchmarkPipOneMillion(b *testing.B) {
	polygon := pip.NewPolygon(
		[]pip.Point{
			{X: 0.0, Y: 0.0},
			{X: 30.0, Y: 50.0},
			{X: 0.0, Y: 100.0},
			{X: 50.0, Y: 70.0},
			{X: 100.0, Y: 100.0},
			{X: 70.0, Y: 50.0},
			{X: 100.0, Y: 0.0},
			{X: 50.0, Y: 30.0},
		},
	)

	var x float64
	var y float64
	var pts []pip.Point

	for i := 0; i < 1000000; i++ {
		x = 100.0 * rand.Float64()
		y = 100.0 * rand.Float64()
		pts = append(pts, pip.Point{X: x, Y: y})
	}

	b.ResetTimer()
	// Actually test the function
	for n := 0; n < b.N; n++ {
		for j := 0; j < len(pts); j++ {
			polygon.Contains(pts[j])
		}
	}
}

func BenchmarkPipParallelOneMillion(b *testing.B) {
	polygon := pip.NewPolygon(
		[]pip.Point{
			{X: 0.0, Y: 0.0},
			{X: 30.0, Y: 50.0},
			{X: 0.0, Y: 100.0},
			{X: 50.0, Y: 70.0},
			{X: 100.0, Y: 100.0},
			{X: 70.0, Y: 50.0},
			{X: 100.0, Y: 0.0},
			{X: 50.0, Y: 30.0},
		},
	)

	var x float64
	var y float64
	var pts []pip.Point

	for i := 0; i < 1000000; i++ {
		x = 100.0 * rand.Float64()
		y = 100.0 * rand.Float64()
		pts = append(pts, pip.Point{X: x, Y: y})
	}

	b.ResetTimer()
	// Actually test the function
	for n := 0; n < b.N; n++ {
		pip.PointInPolygonParallel(pts, polygon, 7)
	}
}

func assert(a bool, b bool, t *testing.T) bool {
	test := a == b
	t.Log("Was the point was correctly identified? " + strconv.FormatBool(test))
	return a == b
}
