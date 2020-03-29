package quadtree

type QuadTree struct {
	bbox     *BBox
	capacity int

	points Points

	topLeft     *QuadTree
	topRight    *QuadTree
	bottomLeft  *QuadTree
	bottomRight *QuadTree
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

// CountPoints returns the total number of points in the tree
func (q QuadTree) CountPoints() int {
	length := len(q.points)
	if q.hasSubdivided() {
		length += q.topLeft.CountPoints()
		length += q.topRight.CountPoints()
		length += q.bottomLeft.CountPoints()
		length += q.bottomRight.CountPoints()
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

	if q.hasSubdivided() {
		found := q.topLeft.GetPointsWithin(r)
		found = append(found, q.topRight.GetPointsWithin(r)...)
		found = append(found, q.bottomLeft.GetPointsWithin(r)...)
		found = append(found, q.bottomRight.GetPointsWithin(r)...)
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

	if !q.hasSubdivided() && len(q.points) < q.capacity {
		q.points = append(q.points, p)
		return true
	}

	// otherwise we should subdivide if we haven't already
	if !q.hasSubdivided() {
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
	return q.topLeft.Insert(p) ||
		q.topRight.Insert(p) ||
		q.bottomLeft.Insert(p) ||
		q.bottomRight.Insert(p)
}

// hasSubdivided evalulates whether the quadtree has been
// subdivided
func (q QuadTree) hasSubdivided() bool {
	return q.topLeft != nil
}

func (q *QuadTree) subdivide() {
	q.topLeft = NewQuadTree(NewBBox(
		q.bbox.xMin,
		q.bbox.xMax-q.bbox.Width()/2,
		q.bbox.yMin+q.bbox.Height()/2,
		q.bbox.yMax,
	), q.capacity)

	q.topRight = NewQuadTree(NewBBox(
		q.bbox.xMin+q.bbox.Width()/2,
		q.bbox.xMax,
		q.bbox.yMin+q.bbox.Height()/2,
		q.bbox.yMax,
	), q.capacity)

	q.bottomLeft = NewQuadTree(NewBBox(
		q.bbox.xMin,
		q.bbox.xMax-q.bbox.Width()/2,
		q.bbox.yMin,
		q.bbox.yMax-q.bbox.Height()/2,
	), q.capacity)

	q.bottomRight = NewQuadTree(NewBBox(
		q.bbox.xMin+q.bbox.Width()/2,
		q.bbox.xMax,
		q.bbox.yMin,
		q.bbox.yMax-q.bbox.Height()/2,
	), q.capacity)
}
