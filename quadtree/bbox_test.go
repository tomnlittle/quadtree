package quadtree_test

import (
	"quadtree/quadtree"
	"testing"
)

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		Description   string
		CorrectResult bool
		Point         quadtree.Point
		BBox          quadtree.BBox
	}{
		{
			Description:   "Base Case - Point is the centre of the BBox",
			CorrectResult: true,
			Point:         quadtree.Point{0, 0},
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
		},
		{
			Description:   "Edge Case - Point is at (0,0) and BBox has no size",
			CorrectResult: true,
			Point:         quadtree.Point{0, 0},
			BBox:          *quadtree.NewBBox(0, 0, 0, 0),
		},
		{
			Description:   "Edge Case - Point is at the top right of the BBox",
			CorrectResult: true,
			Point:         quadtree.Point{50, 50},
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
		},
		{
			Description:   "Edge Case - Point is at the top left of the BBox",
			CorrectResult: true,
			Point:         quadtree.Point{-50, 50},
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
		},
		{
			Description:   "Edge Case - Point is at the bottom left of the BBox",
			CorrectResult: true,
			Point:         quadtree.Point{-50, -50},
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
		},
		{
			Description:   "Edge Case - Point is at the bottom right of the BBox",
			CorrectResult: true,
			Point:         quadtree.Point{50, -50},
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
		},
		{
			Description:   "Edge Case - Point slightly outside of BBox",
			CorrectResult: false,
			Point:         quadtree.Point{50, -50.1},
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
		},
		{
			Description:   "Edge Case - BBox is not at (0,0)",
			CorrectResult: false,
			Point:         quadtree.Point{-100, -100},
			BBox:          *quadtree.NewBBox(100, 100, 100, 100),
		},
		{
			Description:   "Edge Case - BBox shape is not square",
			CorrectResult: true,
			Point:         quadtree.Point{100, 5},
			BBox:          *quadtree.NewBBox(100, 100, 10, 190.6),
		},
	} {
		if tc.BBox.ContainsPoint(&tc.Point) != tc.CorrectResult {
			t.Log(tc.Description)
			t.Errorf("Expected Contains to return %v but got %v instead", tc.CorrectResult, tc.BBox.ContainsPoint(&tc.Point))
		}
	}
}

func TestIntersects(t *testing.T) {
	for _, tc := range []struct {
		Description   string
		CorrectResult bool
		BBox          quadtree.BBox
		InputBBox     quadtree.BBox
	}{
		{
			Description:   "Base Case - BBOX's are equal",
			CorrectResult: true,
			BBox:          *quadtree.NewBBox(0, 0, 0, 0),
			InputBBox:     *quadtree.NewBBox(0, 0, 0, 0),
		},
		{
			Description:   "Base Case - BBOX's do not intersect",
			CorrectResult: false,
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
			InputBBox:     *quadtree.NewBBox(300, 300, 100, 100),
		},
		{
			Description:   "Edge Case - BBOX's have overlapping boundaries",
			CorrectResult: true,
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
			InputBBox:     *quadtree.NewBBox(100, 0, 100, 100),
		},
		{
			Description:   "Edge Case - BBOX's do not have overlapping boundaries",
			CorrectResult: false,
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
			InputBBox:     *quadtree.NewBBox(101, 0, 100, 100),
		},
		{
			Description:   "Edge Case - BBOX's overlap at a corner",
			CorrectResult: true,
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
			InputBBox:     *quadtree.NewBBox(100, 100, 100, 100),
		},
	} {
		if tc.BBox.IntersectsBBox(&tc.InputBBox) != tc.CorrectResult {
			t.Log(tc.Description)
			t.Errorf("Expected Intersects to return %v but got %v instead", tc.CorrectResult, tc.BBox.IntersectsBBox(&tc.InputBBox))
		}
	}
}

func TestContainedBy(t *testing.T) {
	for _, tc := range []struct {
		Description   string
		CorrectResult bool
		BBox          quadtree.BBox
		InputBBox     quadtree.BBox
	}{
		{
			Description:   "Base Case - BBOX's are equal",
			CorrectResult: true,
			BBox:          *quadtree.NewBBox(0, 0, 0, 0),
			InputBBox:     *quadtree.NewBBox(0, 0, 0, 0),
		},
		{
			Description:   "Base Case - Input BBox is smaller than BBox",
			CorrectResult: true,
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
			InputBBox:     *quadtree.NewBBox(0, 0, 99, 99),
		},
		{
			Description:   "Base Case - Input BBox is larger than BBox",
			CorrectResult: false,
			BBox:          *quadtree.NewBBox(0, 0, 100, 100),
			InputBBox:     *quadtree.NewBBox(0, 0, 101, 101),
		},
	} {
		if tc.BBox.ContainsBBox(&tc.InputBBox) != tc.CorrectResult {
			t.Log(tc.Description)
			t.Errorf("Expected ContainedBy to return %v but got %v instead", tc.CorrectResult, tc.BBox.ContainsBBox(&tc.InputBBox))
		}
	}
}
