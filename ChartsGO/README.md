# ChartsGolang
## HOW TO USE
**From main.go example:**
How to assign data for a chart and render it:__
**- Create a Cfg and Dataset type separately not using functions, and then render the chart:**__
```
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
  ```
OR__
**- Create Cfg using "NewChartConfiguration" and Dataset type using "NewChartDataset" in "NewChart" function, and then render the chart:**
```
charts.NewChart(
		charts.NewChartConfiguration(0, "LineExample2", "../resources/templates/template.gohtml", "../../resources/libs/", "../Examples/HTMLFile/"),
		charts.NewChartDataset("LineExample2", []interface{}{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}, []interface{}{325, 420, 200, 111, 450, 555, 666, 455, 980, 100}),
	).Render(&wg)
```
