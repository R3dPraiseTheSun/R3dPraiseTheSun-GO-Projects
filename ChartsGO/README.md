# ChartsGolang
## Cfg and Dataset struct for charts
```
type Cfg struct {
	ChartType         uint
	ChartLegend       string
	TemplatePath      string
	JSLibrary         string
	ScreenCapturePath string
}

type Series struct {
	Name   string
	DataOX []interface{}
	DataOY []interface{}
}
type ChartDataset []Series
```
### Cfg
**ChartType** can be 0 or 1 where 0 is **Line chart type** and 1 is **Bars chart type**.\
**ChartLegend** A simple description.\
**TemplatePath** This saves **the path to the template** used (in Example TemplatePath is "../resources/templates/template.gohtml".\
**JSLibrary** This saves **the path to the JSLibrary** used (in Example we use [uPlots](https://github.com/leeoniya/uPlot)).\
**ScreenCapturePath** This saves the path to where we want to save the HTML file and the screenshot.\

### Dataset
**Name** is saving the name of the chart. And this is going to be the title of the chart used in _Render()_ function.\
**DataOX** and **DataOY** is the data of **OX** and respectively **OY** axes.\

## HOW TO USE
**From _main.go_ we use this example:**\
How to assign data for a chart and render it:\
**- Create a Cfg and Dataset type separately not using functions, and then render the chart:**
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
OR\
**- Create Cfg using "NewChartConfiguration" and Dataset type using "NewChartDataset" in "NewChart" function, and then render the chart:**
```
charts.NewChart(
		charts.NewChartConfiguration(0, "LineExample2", "../resources/templates/template.gohtml", "../../resources/libs/", "../Examples/HTMLFile/"),
		charts.NewChartDataset("LineExample2", []interface{}{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}, []interface{}{325, 420, 200, 111, 450, 555, 666, 455, 980, 100}),
	).Render(&wg)
```
