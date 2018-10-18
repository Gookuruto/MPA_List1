package main

import (
	"github.com/Arafatk/glot"
)

func plot(xes,ys,title string,dim int,firstgroupName string,secondGroupName string,x []float64,y1 []float64,y2 []float64,output string){
	dimensions := dim
	// The dimensions supported by the plot
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	pointGroupName := firstgroupName
	style := "lines"
	points := [][]float64{x, y1}
	// Adding a point group
	plot.AddPointGroup(pointGroupName, style, points)
	plot.AddPointGroup(secondGroupName,"line",[][]float64{x,y2})
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle(title)
	// Optional: Setting the title of the plot
	plot.SetXLabel(xes)
	plot.SetYLabel(ys)
	// Optional: Setting label for X and Y axis
	// Optional: Setting axis ranges
	plot.SavePlot(output)
}
