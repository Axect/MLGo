package main

import "github.com/Arafatk/glot"

func main() {
	dim := 2
	persist, debug := false, false
	plot, _ := glot.NewPlot(dim, persist, debug)

	pointGroupName := "Simple Circles"
	style := "points"
	points := [][]float64{{7, 3, 13, 5.6, 11.1}, {12, 13, 11, 1, 7}}

	plot.AddPointGroup(pointGroupName, style, points)
	plot.SetTitle("Example Plot")
	plot.SetXLabel("X")
	plot.SetYLabel("Y")
	plot.SetXrange(-2, 18)
	plot.SetYrange(-2, 18)

	plot.SavePlot("2.png")
}
