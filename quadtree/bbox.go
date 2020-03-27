package quadtree

// BBox defines a region of either points or regions
// in the quadtree
type BBox struct {
	xMin, xMax float64
	yMin, yMax float64
	// CentreX,CentreY are the centre of the BBox
	CentreX, CentreY float64
	// Width/Height are 1/2 +/- the X/Y coordinate
	Width, Height float64
}

func NewBBox(centreX, centreY, width, height float64) *BBox {
	return &BBox{
		xMin:    centreX - width/2,
		xMax:    centreX + width/2,
		yMax:    centreY + height/2,
		yMin:    centreY - height/2,
		CentreX: centreX,
		CentreY: centreY,
		Width:   width,
		Height:  height,
	}
}

// Contains checks whether the provided point is within
// the bounding box
func (b BBox) Contains(p *Point) bool {
	return p.X <= b.xMax &&
		p.X >= b.xMin &&
		p.Y <= b.yMax &&
		p.Y >= b.yMin
}

// Intersects checks if the bbox intersects with the second
// bounding box
func (b BBox) Intersects(b2 *BBox) bool {
	return b.yMax >= b2.yMin && b2.yMax >= b.yMin &&
		b.xMax >= b2.xMin && b2.xMax >= b.xMin
}

// ContainedBy checks if the current bbox is entirely
// within the second bbox
func (b BBox) ContainedBy(b2 *BBox) bool {
	return b2.CentreX-b2.Width <= b.CentreX-b.Width &&
		b2.CentreX+b2.Width >= b.CentreX+b.Width &&
		b2.CentreY-b2.Height <= b.CentreY-b.Height &&
		b2.CentreY+b2.Height >= b.CentreY+b.Height
}
