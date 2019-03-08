package betterpath

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Parse checks the data format and stores points
func (points *Points) Parse() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No file provided\n")
		os.Exit(1)
	}

	csvFile, _ := os.Open(os.Args[1])
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		a, err := strconv.ParseFloat(line[0], 64)
		b, err := strconv.ParseFloat(line[1], 64)
		c, err := strconv.ParseInt(line[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong data format\n")
			os.Exit(1)
		}
		*points = append(*points, Point{
			x: a,
			y: b,
			t: c,
		},
		)
	}
}
