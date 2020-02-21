package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// parse arguments from cmd line, first pos is path to program
	args := os.Args[1:]

	// pathToFile string, employee1, employee2 string
	if len(args) < 3 {
		fmt.Println(`usage: 
    ccm path_to_file employee_1 employee_2

    where - path_to_file is the path to the csv file, including its name
          - employee_1, employee_2 are the names of the employees for which we want to find the common manager`)
		os.Exit(1)
	}

	pathToFile := args[0]
	employee1 := args[1]
	employee2 := args[2]

	// read csv file
	f, err := os.Open(pathToFile)
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
	for _, r := range records {
		w, err := FillWorker(r)
		if err != nil {
			log.Fatal(err)
		}
		workers = append(workers, w)
	}

	// assign the workers (allocate the nodes in the gap)
	directory := NewDag()
	for _, w := range workers {
		directory.AddEmployee(w.name, w.id)
	}

	// assign report relationship (build the edges of the dag)
	// we do the loop two times, the second adding the reporters
	// performance is not a priority here but the coherence of the data
	for _, w := range workers {
		for _, r := range w.reporters {
			if !directory.hasEmployee(r) {
				log.Fatal(fmt.Sprintf("some reporter %s of manager %s is not listed in the organization",
					r, w.name))
			}
			err := directory.AddLink(w.name, r)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// print the graph and the solution
	fmt.Println(directory)
	lca, err := FindLCA(directory, employee1, employee2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("CCM(%s,%s): %s", employee1, employee2, lca) // ccm stands for closest common manager
	// (in graphs' jargon: lca)
}
