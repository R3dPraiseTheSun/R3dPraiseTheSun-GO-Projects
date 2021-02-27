package charts

import (
	"ChartsGolang/cmd/ChromeHandle"
	"ChartsGolang/cmd/TemplateHandler"
	"os"
	"sync"
)

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

type Chart struct {
	Cfg
	Dataset ChartDataset
}

func NewChartConfiguration(cType uint, cLegend string, templatePath string, JSLib string, SSPath string) Cfg {
	return Cfg{
		ChartType:         cType,
		ChartLegend:       cLegend,
		TemplatePath:      templatePath,
		JSLibrary:         JSLib,
		ScreenCapturePath: SSPath,
	}
}

func NewChartDataset(name string, OX []interface{}, OY []interface{}) ChartDataset {
	return ChartDataset{
		Series{
			Name:   name,
			DataOX: OX,
			DataOY: OY,
		},
	}
}

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

func (c *Chart) Render(wg *sync.WaitGroup) error {
	chromehandle.ScreenShotSave(c.ScreenCapturePath+c.Dataset[0].Name+".html", c.Dataset[0].Name, wg)
	return nil
}
