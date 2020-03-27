package quadtree

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func DrawQuadtree(qt *QuadTree, outputFilename string) {
	width := qt.rootRegion.Width
	height := qt.rootRegion.Height
	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	result := drawQt(dest, qt, width, height)

	draw2dimg.SaveToPngFile(outputFilename, result)
}

func drawQt(img draw.Image, qt *QuadTree, width, height float64) draw.Image {
	img = drawBBox(img, qt.rootRegion, width, height)

	if qt.subdivided {
		img = drawQt(img, qt.topLeft, width, height)
		img = drawQt(img, qt.topRight, width, height)
		img = drawQt(img, qt.bottomLeft, width, height)
		img = drawQt(img, qt.bottomRight, width, height)
	}

	for _, pt := range qt.points {
		img = drawPoint(img, pt, width, height)
	}

	return img
}

func drawBBox(img draw.Image, bbox *BBox, width, height float64) draw.Image {
	newX, newY := toCanvasCoordinates(bbox.CentreX, bbox.CentreY, width, height)

	// Shift the centre of the bbox to the top left which is where
	// we will start drawing the rectangle
	shiftedX := newX - bbox.Width/2
	shiftedY := newY - bbox.Height/2

	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFillColor(color.Black)
	gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetLineWidth(1)

	draw2dkit.Rectangle(gc, shiftedX, shiftedY, shiftedX+bbox.Width, shiftedY+bbox.Height)
	gc.FillStroke()
	gc.Fill()

	// Set the centre of the bbox to green
	img.Set(int(newX), int(newY), color.RGBA{0x00, 0xff, 0x00, 0xff})

	return img
}

func drawPoint(img draw.Image, pt *Point, width, height float64) draw.Image {
	x, y := toCanvasCoordinates(pt.X, pt.Y, width, height)
	img.Set(int(x), int(y), color.RGBA{0xff, 0x00, 0x00, 0xff})
	return img
}

func toCanvasCoordinates(x, y, width, height float64) (float64, float64) {
	x = x + width/2
	y = y - height/2

	if y < 0 {
		y = -y
	}

	return x, y
}
