package helpers

import (
	"encoding/csv"
	"os"
)

func CreateCSV(path string) (*csv.Writer, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)

	return writer, nil
}
