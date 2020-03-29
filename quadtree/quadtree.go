package quadtree

type QuadTree struct {
	rootRegion *BBox
	capacity   int

	points Points

	subdivided bool

	topLeft     *QuadTree
	topRight    *QuadTree
	bottomLeft  *QuadTree
	bottomRight *QuadTree
}

func NewQuadTree(region *BBox, capacity int) *QuadTree {
	return &QuadTree{
		rootRegion: region,
		capacity:   capacity,
		subdivided: false,
	}
}

func (q QuadTree) CountAll() int {
	var length int
	if q.subdivided {
		length += q.topLeft.CountAll()
		length += q.topRight.CountAll()
		length += q.bottomLeft.CountAll()
		length += q.bottomRight.CountAll()
	}

	return len(q.points) + length
}

func (q QuadTree) GetPointsWithin(r *BBox) Points {
	if !q.rootRegion.IntersectsBBox(r) {
		return Points{}
	}

	if q.subdivided {
		found := q.topLeft.GetPointsWithin(r)
		found = append(found, q.topRight.GetPointsWithin(r)...)
		found = append(found, q.bottomLeft.GetPointsWithin(r)...)
		found = append(found, q.bottomRight.GetPointsWithin(r)...)
		return found
	}

	if r.ContainsBBox(q.rootRegion) {
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
	if !q.rootRegion.ContainsPoint(p) {
		return false
	}

	if !q.subdivided && len(q.points) < q.capacity {
		q.points = append(q.points, p)
		return true
	}

	// otherwise we should subdivide if we haven't already
	if !q.subdivided {
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

func (q *QuadTree) insertIntoSubdivisions(p *Point) bool {
	return q.topLeft.Insert(p) ||
		q.topRight.Insert(p) ||
		q.bottomLeft.Insert(p) ||
		q.bottomRight.Insert(p)
}

func (q *QuadTree) subdivide() {
	q.topLeft = NewQuadTree(NewBBox(
		q.rootRegion.CentreX-q.rootRegion.Width/4,
		q.rootRegion.CentreY+q.rootRegion.Height/4,
		q.rootRegion.Width/2,
		q.rootRegion.Height/2,
	), q.capacity)

	q.topRight = NewQuadTree(NewBBox(
		q.rootRegion.CentreX+q.rootRegion.Width/4,
		q.rootRegion.CentreY+q.rootRegion.Height/4,
		q.rootRegion.Width/2,
		q.rootRegion.Height/2,
	), q.capacity)

	q.bottomLeft = NewQuadTree(NewBBox(
		q.rootRegion.CentreX-q.rootRegion.Width/4,
		q.rootRegion.CentreY-q.rootRegion.Height/4,
		q.rootRegion.Width/2,
		q.rootRegion.Height/2,
	), q.capacity)

	q.bottomRight = NewQuadTree(NewBBox(
		q.rootRegion.CentreX+q.rootRegion.Width/4,
		q.rootRegion.CentreY-q.rootRegion.Height/4,
		q.rootRegion.Width/2,
		q.rootRegion.Height/2,
	), q.capacity)

	q.subdivided = true
}
