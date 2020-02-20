// lca.go provides a function to find the least common ancestor of two nodes in a directed acyclic graph (dga)
// it is adapted for a dga representing a hierarchical organization with a CEO at the root node
// https://godoc.org/github.com/funkygao/golib/dag
package main

import (
	"errors"
)

// FindLCA returns the name of the common manager (ancestor) of two employees (nodes)
// if one of the employees is manager of the other, the earlier is returned
// returns an error if the dg is cyclic
func FindLCA(dag *Dag, CEOName string, employeeOne, employeeTwo string) (string, error) {
	root := dag.nodes[dag.contacts[CEOName]]
	keyOne := dag.contacts[employeeOne]
	keyTwo := dag.contacts[employeeTwo]

	var keyOnePath []int
	var keyTwoPath []int

	if dag.isCyclic() {
		return "", errors.New("not a DAG") // lca not defined if not a dag
	}
	if !findPath(dag, root, &keyOnePath, keyOne) || !findPath(dag, root, &keyTwoPath, keyTwo) {
		return "", nil
	}
	var i int
	for i = 0; i < len(keyOnePath) && i < len(keyTwoPath); i++ {
		if keyOnePath[i] != keyTwoPath[i] {
			break
		}
	}
	return dag.nodes[keyOnePath[i-1]].name, nil
}

// findPath is a recursive helper function for FindLCA()
func findPath(dag *Dag, node *Node, path *[]int, keyToFind int) bool {
	if dag == nil {
		return false
	}
	*path = append(*path, node.id)
	if node.id == keyToFind {
		return true
	}
	for v := range node.reporters {
		if findPath(dag, node.reporters[v], path, keyToFind) {
			return true
		}
	}
	n := len(*path) - 1
	*path = (*path)[:n]
	return false
}