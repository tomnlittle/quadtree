package quadtree

type QuadTree struct {
	bbox     *BBox
	capacity int

	points Points

	TopLeft     *QuadTree
	TopRight    *QuadTree
	BottomLeft  *QuadTree
	BottomRight *QuadTree
}

// NewQuadTree creates a new quadtree
func NewQuadTree(bbox *BBox, capacity int) *QuadTree {
	if bbox == nil || capacity <= 0 {
		return nil
	}

	return &QuadTree{
		bbox:     bbox,
		capacity: capacity,
	}
}

// BBox returns a copy of the quadtree's bounding box
func (q QuadTree) BBox() BBox {
	return *q.bbox
}

// HasSubdivided evalulates whether the quadtree has been
// subdivided
func (q QuadTree) HasSubdivided() bool {
	return q.TopLeft != nil
}

// CountPoints returns the total number of points in the tree
func (q QuadTree) CountPoints() int {
	length := len(q.points)
	if q.HasSubdivided() {
		length += q.TopLeft.CountPoints()
		length += q.TopRight.CountPoints()
		length += q.BottomLeft.CountPoints()
		length += q.BottomRight.CountPoints()
	}

	return length
}

// GetPointsWithin returns the points that are contained by the
// input BBox
func (q QuadTree) GetPointsWithin(r *BBox) Points {
	if r == nil {
		return Points{}
	}

	if !q.bbox.IntersectsBBox(r) {
		return Points{}
	}

	if q.HasSubdivided() {
		found := q.TopLeft.GetPointsWithin(r)
		found = append(found, q.TopRight.GetPointsWithin(r)...)
		found = append(found, q.BottomLeft.GetPointsWithin(r)...)
		found = append(found, q.BottomRight.GetPointsWithin(r)...)
		return found
	}

	if r.ContainsBBox(q.bbox) {
		return q.points
	}

	found := Points{}
	for _, p := range q.points {
		if r.ContainsPoint(p) {
			found = append(found, p)
		}
	}
	return found
}

// Insert adds a point to the quadtree
func (q *QuadTree) Insert(p *Point) bool {
	if p == nil {
		return false
	}

	if !q.bbox.ContainsPoint(p) {
		return false
	}

	if !q.HasSubdivided() && len(q.points) < q.capacity {
		q.points = append(q.points, p)
		return true
	}

	// otherwise we should subdivide if we haven't already
	if !q.HasSubdivided() {
		q.subdivide()

		// re-insert all the points back into the quadtree
		for _, pt := range q.points {
			q.insertIntoSubdivisions(pt)
		}
		q.points = nil
	}

	// finally add the point we tried to add in the first place
	return q.insertIntoSubdivisions(p)
}

func (q QuadTree) insertIntoSubdivisions(p *Point) bool {
	return q.TopLeft.Insert(p) ||
		q.TopRight.Insert(p) ||
		q.BottomLeft.Insert(p) ||
		q.BottomRight.Insert(p)
}

func (q *QuadTree) subdivide() {
	q.TopLeft = NewQuadTree(NewBBox(
		q.bbox.xMin,
		q.bbox.xMax-q.bbox.Width()/2,
		q.bbox.yMin+q.bbox.Height()/2,
		q.bbox.yMax,
	), q.capacity)

	q.TopRight = NewQuadTree(NewBBox(
		q.bbox.xMin+q.bbox.Width()/2,
		q.bbox.xMax,
		q.bbox.yMin+q.bbox.Height()/2,
		q.bbox.yMax,
	), q.capacity)

	q.BottomLeft = NewQuadTree(NewBBox(
		q.bbox.xMin,
		q.bbox.xMax-q.bbox.Width()/2,
		q.bbox.yMin,
		q.bbox.yMax-q.bbox.Height()/2,
	), q.capacity)

	q.BottomRight = NewQuadTree(NewBBox(
		q.bbox.xMin+q.bbox.Width()/2,
		q.bbox.xMax,
		q.bbox.yMin,
		q.bbox.yMax-q.bbox.Height()/2,
	), q.capacity)
}
