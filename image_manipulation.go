package main

import (
	"errors"
	"image"
	"image/color"
)

const ChunkSize = 20

var VerticesOutOfBoundsError = errors.New("vertices are not within the bounds of picture")

func validateVertices(img *image.RGBA, vertices []image.Point) bool {
	for _, elem := range vertices {
		if elem.X >= img.Bounds().Max.X {
			return false
		}
		if elem.Y >= img.Bounds().Max.Y {
			return false
		}
	}
	return true
}

func processPartial(img *image.RGBA, vertices []image.Point, color color.Color, offsetX, offsetY, limitX, limitY int, channel chan bool) {
	for x := offsetX; x < limitX; x++ {
		for y := offsetY; y < limitY; y++ {
			if IsInsideConvexShape(image.Point{X: x, Y: y}, vertices[0], vertices[1], vertices[2], vertices[3]) {
				img.Set(x, y, color)
			}
		}
	}
	channel <- true
}

func ProcessImage(img *image.RGBA, vertices []image.Point, color color.Color) error {
	if !validateVertices(img, vertices) {
		return VerticesOutOfBoundsError
	}
	var channels []chan bool
	for i := 0; i < img.Bounds().Max.X; i += ChunkSize {
		xLimit := i + ChunkSize
		if xLimit > img.Bounds().Max.X {
			xLimit = img.Bounds().Max.X
		}
		for j := 0; j < img.Bounds().Max.Y; j += ChunkSize {
			yLimit := j + ChunkSize
			if yLimit > img.Bounds().Max.Y {
				yLimit = img.Bounds().Max.Y
			}
			channel := make(chan bool)
			go processPartial(img, vertices, color, i, j, xLimit, yLimit, channel)
			channels = append(channels, channel)
		}
	}
	for _, channel := range channels {
		<-channel
	}
	return nil
}
