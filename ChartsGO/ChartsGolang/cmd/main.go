package main

import (
	"html/template"
	"sync"

	"ChartsGolang/cmd/Charts"
)

var tpl *template.Template

func main() {
	ConfigChart := charts.Cfg{
		ChartType:         1,
		ChartLegend:       "line2",
		TemplatePath:      "../resources/templates/line.gohtml",
		JSLibrary:         "../resources/libs/uPlot.iife.js",
		ScreenCapturePath: "resouces2",
	}
	Series := charts.Series{
		Name:   "LineChartExample",
		DataOX: []interface{}{100, 200, 210, 400, 500},
		DataOY: []interface{}{325, 420, 102, 300, 444},
	}
	Dataset := charts.ChartDataset{
		Series,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	charts.NewChart(
		charts.NewChartConfiguration(1, "line", "../resources/templates/line.gohtml", "..resources/libs/uplot.iife.js", "resources"),
		charts.NewChartDataset("Line", []interface{}{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}, []interface{}{325, 420, 200, 111, 450, 555, 666, 455, 980, 100}),
	).Render(&wg)

	Chart := charts.NewChart(ConfigChart, Dataset)
	wg.Add(1)
	Chart.Render(&wg)
	wg.Wait()
}
