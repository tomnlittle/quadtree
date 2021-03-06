package main

import (
	"math/rand"
	"quadtree/quadtree"
	"time"
)

const bboxSize = 200
const numPoints = 100
const capacity = 8
const outputFilename string = "output.png"

func main() {
	rand.Seed(time.Now().UnixNano())

	qt := quadtree.NewQuadTree(quadtree.NewBBox(-bboxSize, bboxSize, -bboxSize, bboxSize), capacity)

	for i := 0; i < numPoints; i++ {
		qt.Insert(&quadtree.Point{
			X: randomCoordinate(bboxSize),
			Y: randomCoordinate(bboxSize),
		})
	}

	quadtree.DrawQuadTree(qt, outputFilename)
}

func randomCoordinate(bboxSize float64) float64 {
	val := rand.Float64()
	if rand.Float64() <= 0.5 {
		val = val * -1
	}
	return val * bboxSize
}
