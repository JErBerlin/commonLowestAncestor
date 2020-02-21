package main

import (
	"fmt"
	"log"
	"os"
)

const PathToFile = "./directory.csv"

func main() {

	// read csv file
	f, err := os.Open(PathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	records, err := ReadRecords(f)
	if err != nil {
		log.Fatal(err)
	}

	// fill up the array of data structure
	var workers []Worker
	for _,r := range records {
		w, err := FillWorker(r)
		if err != nil {
			log.Fatal(err)
		}
		workers = append(workers, w)
	}

	// assign the workers (allocate the nodes in the gap)
	directory := NewDag()
	for _, w := range workers {
		directory.AddEmployee(w.name,w.id)
	}

	// assign report relationship (build the edges of the dag)
	// we do the loop two times, the second adding the reporters
	// performance is not a priority here but the coherence of the data
	for _, w := range workers {
		for _,r := range w.reporters {
			if !directory.hasEmployee(r) {
				log.Fatal(fmt.Sprintf("some reporter %s of manager %s is not listed in the organization",
					r, w.name ))
			}
			err := directory.AddLink(w.name,r)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println(directory)
	CEO := "Markus"
	employee1 := "Akshai"
	employee2 := "Martin"
	lca, err := FindLCA(directory, CEO, employee1, employee2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("LCA(%s,%s): %s", employee1, employee2, lca)
}

