package main

import (
	"html/template"
	"sync"

	"ChartsGolang/cmd/Charts"
)

var tpl *template.Template

func main() {
	Run()
}

func Run() {
	ConfigChart := charts.Cfg{
		ChartType:         0,
		ChartLegend:       "Line Example 1",
		TemplatePath:      "../resources/templates/template.gohtml",
		JSLibrary:         "../../resources/libs/",
		ScreenCapturePath: "../Examples/HTMLFile/",
	}
	Series := charts.Series{
		Name:   "LineExample1",
		DataOX: []interface{}{100, 200, 210, 400, 500},
		DataOY: []interface{}{325, 420, 102, 300, 444},
	}
	Dataset := charts.ChartDataset{
		Series,
	}
	Chart := charts.NewChart(ConfigChart, Dataset)
	var wg sync.WaitGroup
	wg.Add(1)
	Chart.Render(&wg)

	wg.Add(1)
	charts.NewChart(
		charts.NewChartConfiguration(0, "LineExample2", "../resources/templates/template.gohtml", "../../resources/libs/", "../Examples/HTMLFile/"),
		charts.NewChartDataset("LineExample2", []interface{}{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}, []interface{}{325, 420, 200, 111, 450, 555, 666, 455, 980, 100}),
	).Render(&wg)
	wg.Add(1)
	charts.NewChart(
		charts.NewChartConfiguration(1, "Bars Example 1", "../resources/templates/template.gohtml", "../../resources/libs/", "../Examples/HTMLFile/"),
		charts.NewChartDataset("BarsExample1", []interface{}{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []interface{}{325, 420, 200, 111, 450, 555, 666, 455, 980, 100}),
	).Render(&wg)
	wg.Add(1)
	charts.NewChart(
		charts.NewChartConfiguration(1, "Bars Example 2", "../resources/templates/template.gohtml", "../../resources/libs/", "../Examples/HTMLFile/"),
		charts.NewChartDataset("BarsExample2", []interface{}{100, 200, 300, 400, 500, 600, 700}, []interface{}{980, 450, 200, 111, 420, 555, 666, 455, 325, 100}),
	).Render(&wg)

	wg.Wait()
}
