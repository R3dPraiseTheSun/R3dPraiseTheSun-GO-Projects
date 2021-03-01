package charts

import (
	"ChartsGolang/cmd/ChromeHandle"
	"ChartsGolang/cmd/TemplateHandler"
	"os"
	"sync"
)

//Chart configuration: Type(0 (line) or 1 (bars)), Legend, TemplatePath(the path to the template you want to use(template.gohtml from resources for now), JSLibrary(uplot stuff), ScreenCapturePath(the path where you want to save the html file).
type Cfg struct {
	ChartType         uint
	ChartLegend       string
	TemplatePath      string
	JSLibrary         string
	ScreenCapturePath string
}

//Chart series
type Series struct {
	Name   string
	DataOX []interface{}
	DataOY []interface{}
}
type ChartDataset []Series

type Chart struct {
	Cfg
	Dataset ChartDataset
}

//NewChartConfiguration is used to create a new configuration and returns the Cfg struct
func NewChartConfiguration(cType uint, cLegend string, templatePath string, JSLib string, SSPath string) Cfg {
	return Cfg{
		ChartType:         cType,
		ChartLegend:       cLegend,
		TemplatePath:      templatePath,
		JSLibrary:         JSLib,
		ScreenCapturePath: SSPath,
	}
}

//NewChartDataset is used to create a new Dataset from a chart and return the ChartDataset struct
func NewChartDataset(name string, OX []interface{}, OY []interface{}) ChartDataset {
	return ChartDataset{
		Series{
			Name:   name,
			DataOX: OX,
			DataOY: OY,
		},
	}
}

//NewChart function is used to create the chart object, also Execute function from template package is creating a new html file to the path set in Cfg
func NewChart(cfg Cfg, data ChartDataset) *Chart {
	//create html file
	tpl, err := templateshandle.GetTemplate(cfg.TemplatePath)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(cfg.ScreenCapturePath + data[0].Name + ".html")
	if err != nil {
		panic(err)
	}
	execerr := tpl.Execute(f, &Chart{Cfg: cfg, Dataset: data})
	if execerr != nil {
		panic(execerr)
	}
	//return chart for render
	return &Chart{
		Cfg:     cfg,
		Dataset: data,
	}
}

//Render function is saving a screenshot of the page opened in headless chrome
func (c *Chart) Render(wg *sync.WaitGroup) error {
	chromehandle.ScreenShotSave(c.ScreenCapturePath+c.Dataset[0].Name+".html", c.Dataset[0].Name, wg)
	return nil
}
