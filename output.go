package betterpath

import (
	"fmt"
	"os"
)

// Output prints the updated dataset in a csv file
func (points Points) Output() error {

	f, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	for _, point := range points {
		s := fmt.Sprintf("%.8f,%.9f,%d", point.x, point.y, point.t)
		_, err := f.WriteString(s)
		if err != nil {
			f.Close()
			return err
		}
	}
	err = f.Close()
	return err
}
