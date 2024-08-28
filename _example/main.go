package main

import (
	"encoding/json"
	"flag"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/motoki317/pip-go"
)

var (
	npoints = flag.Int("n", 30, "npoints")
	width   = flag.Int("w", 500, "width")
	height  = flag.Int("h", 500, "height")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	var points []pip.Point
	for n := 0; n < *npoints; n++ {
		points = append(points, pip.Point{X: rand.Float64() * float64(*width), Y: rand.Float64() * float64(*height)})
	}
	polygon := pip.NewPolygon(points)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("index.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.HandleFunc("/polygon", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&points)
	})
	http.HandleFunc("/hit", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query()
		x, _ := strconv.ParseFloat(param.Get("x"), 64)
		y, _ := strconv.ParseFloat(param.Get("y"), 64)
		res := struct {
			Result bool `json:"result"`
		}{polygon.Contains(pip.Point{X: x, Y: y})}
		json.NewEncoder(w).Encode(&res)
	})
	http.ListenAndServe(":8080", nil)
}
