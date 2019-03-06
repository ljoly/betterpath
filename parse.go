package betterpath

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Parse parse csv file
func Parse() {

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
		fmt.Println(line)
	}
}
