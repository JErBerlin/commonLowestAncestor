package main

import "fmt"

func main() {

	directory := NewDag()

	directory.AddEmployee("Markus", 0)

	directory.AddEmployee("Erik", 1)
	directory.AddEmployee("Artashes", 2)
	directory.AddLink("Markus","Erik")
	directory.AddLink("Markus","Artashes")

	directory.AddEmployee("Akshai", 3)
	directory.AddEmployee("Martin", 4)
	directory.AddLink("Artashes","Akshai")
	directory.AddLink("Artashes","Martin")

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

