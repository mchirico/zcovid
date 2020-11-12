package guage

import (
	"fmt"
	"io"


	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Gauge struct {
	Title string
	Name string
	NameGdata string
	Value int
}

func (g Gauge) gaugeBase() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: g.Title}),
	)


	gauge.AddSeries(g.Name, []opts.GaugeData{{Name: g.NameGdata, Value: g.Value}})
	return gauge
}

func (g Gauge)gaugeTimer() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: g.Title}),
	)

	gauge.AddSeries(g.Title, []opts.GaugeData{{Name: g.NameGdata, Value: g.Value}})

	fn := fmt.Sprintf(`setInterval(function () {
			option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
			myChart_%s.setOption(option_%s, true);
		}, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
	gauge.AddJSFuncs(fn)
	return gauge
}

type GaugeExamples struct{}

func (GaugeExamples) Examples() {

	g0 := Gauge{
		Title:     "Title G0",
		Name:      "Name G0",
		NameGdata: "Name G0 Data",
		Value:     60,
	}

	g1 := Gauge{
		Title:     "Title G1",
		Name:      "Name G1",
		NameGdata: "Name G1 Data",
		Value:     25,
	}

	page := components.NewPage()
	page.AddCharts(
		g0.gaugeBase(),
		g1.gaugeTimer(),
	)

	f, err := os.Create("gauge.html")
	if err != nil {
		panic(err)
	}

	_ = f
	page.Render(io.MultiWriter(f))

}
