package quadtree_test

import (
	"math/rand"
	"quadtree/quadtree"
	"time"

	"testing"
)

func TestQuadTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Generate a random set of points for the base case
	var baseCaseBoundingBoxSize float64 = 200
	points := quadtree.Points{}
	for i := 0; i < 5000; i++ {
		points = append(points, &quadtree.Point{
			X: randomCoordinate(baseCaseBoundingBoxSize),
			Y: randomCoordinate(baseCaseBoundingBoxSize),
		})
	}

	for _, tc := range []struct {
		Description            string
		Capacity               int
		BBox                   *quadtree.BBox
		TestBBox               *quadtree.BBox
		Points                 *quadtree.Points
		ExpectedNumberOfPoints int
	}{
		{
			Description:            "Base Case",
			Capacity:               1,
			BBox:                   quadtree.NewBBox(-baseCaseBoundingBoxSize, baseCaseBoundingBoxSize, -baseCaseBoundingBoxSize, baseCaseBoundingBoxSize),
			TestBBox:               quadtree.NewBBox(-baseCaseBoundingBoxSize, baseCaseBoundingBoxSize, -baseCaseBoundingBoxSize, baseCaseBoundingBoxSize),
			Points:                 &points,
			ExpectedNumberOfPoints: len(points),
		},
		{
			Description: "One Point outside of the queried BBox",
			Capacity:    1,
			BBox:        quadtree.NewBBox(-200, 200, -200, 200),
			TestBBox:    quadtree.NewBBox(-20, 20, -20, 20),
			Points: &quadtree.Points{
				&quadtree.Point{10, 10},
				&quadtree.Point{10, 11},
				&quadtree.Point{10, 12},
				&quadtree.Point{10, 13},
				&quadtree.Point{-150, -150},
			},
			ExpectedNumberOfPoints: 4,
		},
	} {
		qt := quadtree.NewQuadTree(tc.BBox, tc.Capacity)

		insertedPoints := make(map[quadtree.Point]bool)
		for _, p := range *tc.Points {
			insertedPoints[*p] = false
			if qt.Insert(p) == false {
				t.Log(tc.Description)
				t.Errorf("Unable to insert point")
			}
		}

		// If we query using the original BBox that was used to create the
		// Quadtree then we should get all the points that we originally inserted
		r := qt.GetPointsWithin(tc.TestBBox)

		if len(r) != tc.ExpectedNumberOfPoints {
			t.Log(tc.Description)
			t.Errorf("Expected there to be %v points but found %v", tc.ExpectedNumberOfPoints, len(r))
		}

		for _, pt := range r {
			if value, has := insertedPoints[*pt]; !has {
				t.Log(tc.Description)
				t.Errorf("Inserted Point doesn't appear in the list of inserted points")
			} else if value {
				t.Log(tc.Description)
				t.Errorf("Point is duplicated in quadtree")
			}

			insertedPoints[*pt] = true
		}
	}
}

func randomCoordinate(bboxSize float64) float64 {
	val := rand.Float64()
	if rand.Float64() <= 0.5 {
		val = val * -1
	}
	return val * bboxSize
}
