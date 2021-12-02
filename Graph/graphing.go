package Graph

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"net/http"
)

var Labels *[]string
var Categories *[]string
var Inputs *[]float64

// generate random data for bar chart
func generateRandomBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(*Categories); i++ {
		items = append(items, opts.BarData{Value: rand.Intn(999)})
	}
	return items
}

func generateBarItems(r []float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(*Inputs); i++ {
		for _,v := range r{
			items = append(items, opts.BarData{Value: v})
		}
	}
	return items
}

func generatePieItems(data []float64) []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < len(*Labels); i++ {
		items = append(items, opts.PieData{Value: data})
	}
	return items
}

func generateLineValues(inputs []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(inputs); i++ {
		items = append(items, opts.LineData{Value: inputs[i]})
	}
	return items
}

func CreatePieGraph(w http.ResponseWriter) {
	pie := charts.NewPie()

	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Today's Profit",
		Subtitle: "It's extremely easy to use, right?",
	}))
}

func CreateBarGraph(w http.ResponseWriter){
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))


	bar.SetXAxis(Labels)
	for _, v := range *Categories{
		bar.AddSeries(v, generateBarItems(*Inputs))
	}
	// Put data into instance

	// Where the magic happens
	bar.Render(w)
}

func CreateLineOverTime(labels, categories []string, inputs []float64){
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{PageTitle: "Bronze Hermes Data",Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Graph",
			Subtitle: "Data",
		}))

	// Put data into instance
		line.SetXAxis("labels").
			AddSeries("categories[0]", generateLineValues([]float64{0})).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	//line.Render(render)
}