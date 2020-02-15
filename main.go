package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/subchen/go-cli"
)

func main() {
	var width int
	var height int
	var marker string
	var isLine bool

	app := cli.NewApp()
	app.Name = "splot"
	app.Usage = "simple plots"
	app.Version = "1.0.0"
	app.Flags = []*cli.Flag{
		{
			Name:     "w, width",
			Usage:    "width of the plot",
			DefValue: "100",
			Value:    &width,
		},
		{
			Name:     "h, height",
			Usage:    "height of the plot (an extra row will be added to this value)",
			DefValue: "10",
			Value:    &height,
		},
		{
			Name:     "m, marker",
			Usage:    "the character to use in the plot",
			DefValue: ".",
			Value:    &marker,
		},
		{
			Name:   "l, line",
			Usage:  "when present, displays a line plot instead of full bars",
			IsBool: true,
			Value:  &isLine,
		},
	}

	app.Action = func(c *cli.Context) {
		values, _ := readInput()

		min, max, stepY, stepX, width := determineBounds(values, width, height)

		yAxisLabelWidth := len(fmt.Sprintf("%.2f", max))

		cols := make([]float64, width)
		for i := 0; i < width; i++ {
			cols[i] = getValue(values, stepX, i)
		}

		for y := height + 1; y > 0; y-- {
			cap := min + (stepY * float64(y-1))
			bot := min + (stepY * float64(y))

			for x := 0; x < width; x++ {

				if x == 0 {
					format := "%" + strconv.Itoa(yAxisLabelWidth) + ".2f"
					fmt.Printf(format, cap)
					fmt.Print(" ")
				} else {
					fmt.Print(" ")
				}

				if isLine {
					if cols[x] >= cap && cols[x] <= bot {
						fmt.Print(marker)
					} else {
						fmt.Print(" ")
					}
				} else {
					if cols[x] >= cap {
						fmt.Print(marker)
					} else {
						fmt.Print(" ")
					}
				}

				if x == width-1 {
					fmt.Print("\n")
				}
			}
		}
	}

	app.Run(os.Args)
}

func readInput() ([]float64, error) {
	values := []float64{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			continue
		}
		values = append(values, f)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

func determineBounds(values []float64, width int, height int) (float64, float64, float64, int, int) {
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)

	min := sorted[0]
	max := sorted[len(sorted)-1]

	stepY := (max - min) / float64(height)

	var stepX int
	if len(values) < width {
		stepX = 1
		width = len(values)
	} else {
		stepX = int(math.Ceil(float64(len(values)) / float64(width)))
	}

	return min, max, stepY, stepX, width
}

func getValue(values []float64, group int, iter int) float64 {
	if group == 1 {
		return values[iter]
	}

	acc := float64(0)
	num := 0
	for i := 0; i < group; i++ {
		idx := iter*group + i
		if idx < len(values) {
			acc += values[idx]
			num++
		}
	}

	return acc / float64(num)
}
