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

	qt := quadtree.NewQuadTree(quadtree.NewBBox(0, 0, bboxSize, bboxSize), capacity)

	for i := 0; i < numPoints; i++ {
		qt.Insert(&quadtree.Point{
			X: randomNumber(),
			Y: randomNumber(),
		})
	}

	quadtree.DrawQuadtree(qt, outputFilename)
}

func randomNumber() float64 {
	val := rand.Float64()
	if rand.Float64() <= 0.5 {
		val = val * -1
	}
	return val * bboxSize / 2
}
