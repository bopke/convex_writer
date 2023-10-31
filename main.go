package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
)

var addr = "127.0.0.1:8000"

var WrongDataLengthError = errors.New("wrong data length")
var NotConvexShapeError = errors.New("shape is not convex")

func getColorFromRequest(r *http.Request) (color.Color, error) {
	colorRaw := r.FormValue("color")
	var c []uint8
	err := json.Unmarshal([]byte(colorRaw), &c)
	if err != nil {
		return color.RGBA{}, err
	}
	if len(c) != 4 {
		return color.RGBA{}, WrongDataLengthError
	}
	return color.RGBA{R: c[0], G: c[1], B: c[2], A: c[3]}, nil
}

func getVerticesFromRequest(r *http.Request) ([]image.Point, error) {
	verticesRaw := r.FormValue("vertices")
	var vertices [][]int
	err := json.Unmarshal([]byte(verticesRaw), &vertices)
	if err != nil {
		return nil, err
	}
	if len(vertices) != 4 {
		return nil, WrongDataLengthError
	}

	var points []image.Point
	for _, elem := range vertices {
		if len(elem) != 2 {
			return nil, WrongDataLengthError
		}
		points = append(points, image.Point{
			X: elem[0],
			Y: elem[1],
		})
	}
	if !IsShapeConvex(points) {
		return nil, NotConvexShapeError
	}
	return points, nil
}

func getImageFromRequest(r *http.Request) (*image.RGBA, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	jpg, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	img := image.NewRGBA(jpg.Bounds())
	draw.Draw(img, img.Bounds(), jpg, image.ZP, draw.Src)
	return img, nil
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	color, err := getColorFromRequest(r)
	if err != nil {
		log.Println("Processing color failed", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	vertices, err := getVerticesFromRequest(r)
	if err != nil {
		log.Println("Processing vertices failed", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	img, err := getImageFromRequest(r)
	if err != nil {
		log.Println("Processing image failed", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	err = ProcessImage(img, vertices, color)
	if err != nil {
		log.Println("Processing image failed", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	w.WriteHeader(200)
	err = jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Println("Encoding/sending image failed", err)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", requestHandler).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	log.Panic(srv.ListenAndServe())
}
