package quadtree_test

import (
	"math/rand"
	"quadtree/quadtree"
	"time"

	"testing"
)

func TestNewQuadTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var capacity int = 10
	var bboxSize float64 = 200
	var pointsToInsert int = 1000

	r := quadtree.NewBBox(0, 0, bboxSize, bboxSize)
	qt := quadtree.NewQuadTree(r, capacity)

	for i := 0; i < pointsToInsert; i++ {
		qt.Insert(&quadtree.Point{
			X: randomNumber(bboxSize / 2),
			Y: randomNumber(bboxSize / 2),
		})
	}

	rs := qt.GetPointsWithin(r)

	if len(rs) != pointsToInsert {
		t.Log(qt.CountPoints())
		t.Errorf("Expected there to be %v points but found %v", pointsToInsert, len(rs))
	}
}

func randomNumber(bboxSize float64) float64 {
	val := rand.Float64()
	if rand.Float64() <= 0.5 {
		val = val * -1
	}
	return val * bboxSize
}
