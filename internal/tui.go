package internal

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

//prop proportion towards terminal dimensions
const prop = 4

//title title of plot
const title = "Number of fd for"

//Chart type of plot style
type Chart struct {
	m int //m max value
	c *widgets.BarChart
}

//Line type of plot style
type Line struct {
	m  int //m max value
	l  *widgets.Sparkline
	lg *widgets.SparklineGroup
}

//mkWidgetSize calculate size of plot
func mkWidgetSize() (int, int) {
	x, y := ui.TerminalDimensions()
	return x / prop, y / prop
}

//newFdArray create new slice for fd
func newFdArray() []float64 {
	x, _ := ui.TerminalDimensions()

	arrMax := int(math.Abs(float64(x/prop)-2)) + 1
	arr := make([]float64, arrMax)

	return arr
}

//newChart create new Chart
//x, y are size of new plot table
func newChart(pidd pid, x, y int) Chart {
	chart := Chart{c: nil}

	c := widgets.NewBarChart()
	c.Title = title + strconv.Itoa(pidd)
	c.Data = newFdArray()
	c.SetRect(0, 0, x, y)
	c.BarColors = []ui.Color{ui.ColorGreen}
	c.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	chart.c = c

	return chart
}

//draw draw an updated plot with new fd added
func (c *Chart) draw(fdNum int) {
	c.c.Data = append(c.c.Data[1:], float64(fdNum>>2))
	ui.Render(c.c)
}

//newLine create new Line
//x, y are size of new plot table
func newLine(pidd pid, x, y int) Line {
	line := Line{m: 0, l: nil, lg: nil}

	l := widgets.NewSparkline()
	l.Data = newFdArray()
	l.LineColor = ui.ColorGreen

	lg := widgets.NewSparklineGroup(l)
	lg.Title = title + strconv.Itoa(pidd)
	lg.SetRect(0, 0, x, y)

	line.l = l
	line.lg = lg

	return line
}

//draw draw an updated plot with new fd added
func (l *Line) draw(fdNum int) {
	if len(l.l.Data) == 0 {
		log.Fatalln("`impossible` happened, consider opening an issue " +
			"with \"EDataZero\"")
	}

	if l.m < fdNum {
		l.m = fdNum
	}
	l.lg.Title = fmt.Sprintf("%s %d (max: %d) (cur: %d)", title, fdNum, l.m, fdNum)
	l.l.Data = append(l.l.Data[1:], float64(fdNum))

	ui.Render(l.lg)
}

//RunUI run tui, entrypoint to start tui
func RunUI(pidd pid) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	fd, err := GetFdNum(pidd)
	if err != nil {
		fd = 0
	}

	x, y := mkWidgetSize()
	line := newLine(pidd, x, y)
	line.draw(fd)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			fd, err := GetFdNum(pidd)
			if err != nil {
				fd = 0
			}

			line.draw(fd)
		}
	}
}
