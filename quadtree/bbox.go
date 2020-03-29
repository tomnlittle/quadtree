package quadtree

// BBox defines a bounding box
type BBox struct {
	xMin, xMax float64
	yMin, yMax float64
}

// NewBBox takes xMin,xMax,yMin,yMax and returns a new bounding box
func NewBBox(xMin, xMax, yMin, yMax float64) *BBox {
	if xMax < xMin || yMax < yMin {
		return nil
	}

	return &BBox{
		xMin: xMin,
		xMax: xMax,
		yMin: yMin,
		yMax: yMax,
	}
}

// Width - returns the width of the bbox
func (b BBox) Width() float64 {
	return b.xMax - b.xMin
}

// Height returns the height of the bbox
func (b BBox) Height() float64 {
	return b.yMax - b.yMin
}

// Centre returns the (x,y) centre of the bbox
func (b BBox) Centre() Point {
	return Point{b.xMin + b.Width()/2, b.yMin + b.Height()/2}
}

// ContainsPoint checks whether the provided point is within
// the bounding box
func (b BBox) ContainsPoint(p *Point) bool {
	return p != nil &&
		p.X <= b.xMax &&
		p.X >= b.xMin &&
		p.Y <= b.yMax &&
		p.Y >= b.yMin
}

// ContainsBBox checks if the current bbox is entirely
// within the second bbox
func (b BBox) ContainsBBox(b2 *BBox) bool {
	return b2 != nil &&
		b.xMin <= b2.xMin &&
		b.xMax >= b2.xMax &&
		b.yMin <= b2.yMin &&
		b.yMax >= b2.yMax
}

// IntersectsBBox checks if the bbox intersects with the second
// bounding box
func (b BBox) IntersectsBBox(b2 *BBox) bool {
	return b2 != nil &&
		b.yMax >= b2.yMin && b2.yMax >= b.yMin &&
		b.xMax >= b2.xMin && b2.xMax >= b.xMin
}
