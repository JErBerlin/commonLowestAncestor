// read.go provides a function to read the records of type worker from a csv file
package main

import (
	"encoding/csv"
	"io"
)
// csv reader splits every record line by the separator
type Record = []string

func ReadRecords(f io.Reader) ([]Record, error) {

	r := csv.NewReader(f)
	r.LazyQuotes = true // needed for the field reporters

	_, err := r.Read() // first line is header
	if err != nil {
		return nil, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
