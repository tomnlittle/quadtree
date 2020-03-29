package main

import (
	"image"
	"image/color"
	"image/draw"
	"quadtree/quadtree"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func DrawQuadtree(qt *quadtree.QuadTree, outputFilename string) {
	bbox := qt.BBox()
	width := bbox.Width()
	height := bbox.Height()
	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	result := drawQt(dest, qt, width, height)

	for _, pt := range qt.GetPointsWithin(&bbox) {
		result = drawPoint(result, *pt, width, height)
	}

	draw2dimg.SaveToPngFile(outputFilename, result)
}

func drawQt(img draw.Image, qt *quadtree.QuadTree, width, height float64) draw.Image {
	img = drawBBox(img, qt.BBox(), width, height)

	if qt.HasSubdivided() {
		img = drawQt(img, qt.TopLeft, width, height)
		img = drawQt(img, qt.TopRight, width, height)
		img = drawQt(img, qt.BottomLeft, width, height)
		img = drawQt(img, qt.BottomRight, width, height)
	}

	return img
}

func drawBBox(img draw.Image, bbox quadtree.BBox, width, height float64) draw.Image {
	p := bbox.Centre()
	newPt := getCanvasPoint(p, width, height)

	// Shift the centre of the bbox to the top left which is where
	// we will start drawing the rectangle
	shiftedX := newPt.X - bbox.Width()/2
	shiftedY := newPt.Y - bbox.Height()/2

	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFillColor(color.Black)
	gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetLineWidth(1)

	draw2dkit.Rectangle(gc, shiftedX, shiftedY, shiftedX+bbox.Width(), shiftedY+bbox.Height())
	gc.FillStroke()
	gc.Fill()

	// Set the centre of the bbox to green
	img.Set(int(newPt.X), int(newPt.Y), color.RGBA{0x00, 0xff, 0x00, 0xff})

	return img
}

func drawPoint(img draw.Image, pt quadtree.Point, width, height float64) draw.Image {
	newPt := getCanvasPoint(pt, width, height)
	img.Set(int(newPt.X), int(newPt.Y), color.RGBA{0xff, 0x00, 0x00, 0xff})
	return img
}

func getCanvasPoint(p quadtree.Point, width, height float64) quadtree.Point {
	x := p.X + width/2
	y := p.Y - height/2

	if y < 0 {
		y = -y
	}

	return quadtree.Point{x, y}
}
