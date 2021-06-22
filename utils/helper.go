package utils

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"sync"
)

// Open CSV File func
func OpenCSVFile(filePath string) (*csv.Reader, *os.File, error)  {
	f, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, errors.New("Please upload file csv")
		}
		return nil, nil, err
	}
	reader := csv.NewReader(f)
	return reader, f, nil
}

// Read CSV File per line func
func ReadCSVFilePerLine(csvReader *csv.Reader, jobs chan<- string, wg *sync.WaitGroup) {
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		var rowOrdered string
		for _, each := range row {
			rowOrdered = each
		}
		wg.Add(1)
		jobs <- rowOrdered
	}
	close(jobs)
}