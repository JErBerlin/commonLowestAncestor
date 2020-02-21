// worker.go provides a data type for the features defining a worker in the organisation structure
// we suppose that the information to comes from csv records
// with header 'id, name, "reporter1; ... ; reporterN"'
package main

import (
	"errors"
	"strconv"
	"strings"
)

type Worker struct {
	name string
	id   int

	reporters []string
}

// FillWorker fills the fields of the worker type with the inputs from a record
func FillWorker(record Record) (Worker, error) {
	if len(record) != len([]string{"name","id","reporters"}) {
		return Worker{}, errors.New("filling of worker failed: record has invalid number of fields")
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return Worker{}, err
	}
	record[1] = strings.Trim(record[1], " ")
	name := record[1]
	record[2] = strings.Trim(record[2], " \"")
	if record[2] != "" {
		reporters := strings.Split(record[2], ";")
		for i, s := range reporters {
			s = strings.Trim(s, " ")
			if s != "" {
				reporters[i] = s
			}
		}
		return Worker{name, id, reporters}, nil
	}
	return Worker{name, id, nil}, nil
}