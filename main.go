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
	app.Usage = "Simple plots"
	app.Version = "1.0.0"
	app.Flags = []*cli.Flag{
		{
			Name:     "w, width",
			Usage:    "",
			DefValue: "100",
			Value:    &width,
		},
		{
			Name:     "h, height",
			Usage:    "",
			DefValue: "10",
			Value:    &height,
		},
		{
			Name:     "m, marker",
			Usage:    "",
			DefValue: ".",
			Value:    &marker,
		},
		{
			Name:   "l, line",
			Usage:  "",
			IsBool: true,
			Value:  &isLine,
		},
	}

	app.Action = func(c *cli.Context) {
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
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		sorted := make([]float64, len(values))
		copy(sorted, values)

		sort.Float64s(sorted)

		min := sorted[0]
		max := sorted[len(sorted)-1]
		stepY := int(math.Ceil((max - min) / float64(height)))

		// move the bottom of the range a bit lower so all x values show something
		// on the y axis
		min = min - float64(stepY)*2.0
		stepY = int(math.Ceil((max - min) / float64(height)))

		// if we have less values than the specified width,
		// make the width match the number of values we have
		if len(values) < width {
			width = len(values)
		}

		// TODO - if len(values) > width, remove every nth value? IDK

		for y := 0; y < height; y++ {
			cap := int(math.Ceil(max)) - (stepY * y)
			bot := int(math.Ceil(max)) - (stepY * (y - 1))
			for x := 0; x < width; x++ {
				fmt.Print(" ")
				if x == 0 {
					fmt.Printf("%3d", cap) // TODO - make this work in all cases
					fmt.Print(" ")
				}

				if isLine {
					if values[x] >= float64(cap) && values[x] <= float64(bot) {
						fmt.Print(marker)
					} else {
						fmt.Print(" ")
					}
				} else {
					if values[x] >= float64(cap) {
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
