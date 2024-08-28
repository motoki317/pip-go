package pip

type Point struct {
	X float64
	Y float64
}

type Polygon struct {
	points []Point
	bb     BoundingBox
}

func NewPolygon(points []Point) *Polygon {
	points = append(points, points[0]) // Add first point
	return &Polygon{
		points: points,
		bb:     GetBoundingBox(points),
	}
}

type BoundingBox struct {
	BottomLeft Point
	TopRight   Point
}

// Contains checks if point is inside polygon
func (p *Polygon) Contains(pt Point) bool {
	// If point not in bounding box return false immediately
	if !p.bb.Contains(pt) {
		return false
	}

	// If the point is in the bounding box then we need to check the polygon
	nverts := len(p.points)
	intersect := false

	verts := p.points
	j := 0

	for i := 1; i < nverts; i++ {
		if ((verts[i].Y > pt.Y) != (verts[j].Y > pt.Y)) &&
			(pt.X < (verts[j].X-verts[i].X)*(pt.Y-verts[i].Y)/(verts[j].Y-verts[i].Y)+verts[i].X) {
			intersect = !intersect
		}

		j = i
	}

	return intersect
}

// Contains checks if point is in bounding box
func (bb BoundingBox) Contains(pt Point) bool {
	// Bottom Left is the smallest and x and y value
	// Top Right is the largest x and y value
	return pt.X < bb.TopRight.X && pt.X > bb.BottomLeft.X &&
		pt.Y < bb.TopRight.Y && pt.Y > bb.BottomLeft.Y
}

func GetBoundingBox(points []Point) BoundingBox {
	var maxX, maxY, minX, minY float64

	for _, side := range points {
		if side.X > maxX || maxX == 0.0 {
			maxX = side.X
		}
		if side.Y > maxY || maxY == 0.0 {
			maxY = side.Y
		}
		if side.X < minX || minX == 0.0 {
			minX = side.X
		}
		if side.Y < minY || minY == 0.0 {
			minY = side.Y
		}
	}

	return BoundingBox{
		BottomLeft: Point{X: minX, Y: minY},
		TopRight:   Point{X: maxX, Y: maxY},
	}
}
