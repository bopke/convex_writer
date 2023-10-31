package main

import "image"

func sign(p1, p2, p3 image.Point) int {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}

func IsInsideTriangle(point, a, b, c image.Point) bool {
	d1 := sign(point, a, b)
	d2 := sign(point, b, c)
	d3 := sign(point, c, a)
	hasNeg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos := (d1 > 0) || (d2 > 0) || (d3 > 0)
	return !(hasNeg && hasPos)
}

func IsInsideConvexShape(point, a, b, c, d image.Point) bool {
	return IsInsideTriangle(point, a, b, c) ||
		IsInsideTriangle(point, a, b, d) ||
		IsInsideTriangle(point, a, c, d) ||
		IsInsideTriangle(point, b, c, d)
}

func IsShapeConvex(vertices []image.Point) bool {
	return !IsInsideTriangle(vertices[0], vertices[1], vertices[2], vertices[3]) &&
		!IsInsideTriangle(vertices[1], vertices[0], vertices[2], vertices[3]) &&
		!IsInsideTriangle(vertices[2], vertices[1], vertices[0], vertices[3]) &&
		!IsInsideTriangle(vertices[3], vertices[1], vertices[2], vertices[0])
}
