package betterpath

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

// Parse checks the data format and stores points
func (points *Points) Parse() error {

	if len(os.Args) < 2 {
		return errors.New("No file provided")
	}

	csvFile, err := os.Open(os.Args[1])
	if err != nil {
		return errors.New("Open error")
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New("Read error")
		}

		a, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			return errors.New("Wrong data format")
		}
		b, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return errors.New("Wrong data format")
		}
		c, err := strconv.ParseInt(line[2], 10, 32)
		if err != nil {
			return errors.New("Wrong data format")
		}

		*points = append(*points, Point{
			x: a,
			y: b,
			t: c,
		},
		)
	}
	return nil
}
