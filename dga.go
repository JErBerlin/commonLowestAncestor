// dga.go defines the type dag (from directed acyclic graph) representing a hierarchical organization with a CEO at
// the root node and where every contains pointers to its children
// it provides methods to build the graph and a function to check if the graph is indeed acyclic
// https://godoc.org/github.com/funkygao/golib/dag
package main

import (
	"errors"
	"fmt"
)

type Node struct {
	name string
	id   int

	nManagers int
	reporters []*Node
}

func (n Node) String() string {
	return fmt.Sprintf("%s (%d) --> %v ", n.name, n.nManagers, n.reporters)
}

type MapIdNode map[int]*Node
type MapNameId map[string]int

func (nodes MapIdNode) String() string {
	var line string
	for id, node := range nodes {
		line += fmt.Sprintf("%d: %v\n", id, node)
	}
	return line
}

type Dag struct {
	root      *Node
	nodes     MapIdNode
	contacts  MapNameId
	nContacts int
}

func (dag Dag) String() string {
	if dag.isCyclic() {
		return "cyclic graph"
	}
	return fmt.Sprintf("%v", dag.nodes)
}

func NewDag() *Dag {
	this := new(Dag)
	this.nodes = make(map[int]*Node)
	this.contacts = make(map[string]int)
	return this
}

func (dag *Dag) AddEmployee(name string, id int) *Node {
	node := &Node{name: name, id: id}
	if dag.nContacts == 0 {
		dag.root = node
	}
	dag.nodes[id] = node
	dag.contacts[name] = id
	dag.nContacts++
	return node
}

func (dag *Dag) AddLink(from, to string) error {
	fromId, ok := dag.contacts[from]
	if !ok {
		return errors.New(fmt.Sprintf("Manager %s unknown to this organisation", from))
	}
	toId, ok := dag.contacts[to]
	if !ok {
		return errors.New(fmt.Sprintf("Employee %s unknown to this organisation", to))
	}
	fromNode := dag.nodes[fromId]
	toNode := dag.nodes[toId]
	fromNode.reporters = append(fromNode.reporters, toNode)
	toNode.nManagers++
	return nil
}

func (dag *Dag) hasEmployee(name string) bool {
	_, ok := dag.contacts[name]
	return ok
}

func (dag *Dag) isCyclicUtil(v int, visited []bool, recStack []bool) bool {
	if visited[v] == false {
		// Mark the current node as visited, part of recursion stack
		visited[v] = true
		recStack[v] = true

		// Recur for all the vertices adjacent to this vertex
		for i := 0; i < dag.nodes[v].nManagers; i++ {
			if !visited[i] && dag.isCyclicUtil(i, visited, recStack) {
				return true
			} else if recStack[i] {
				return true
			}
		}

	}
	recStack[v] = false // remove the vertex from recursion stack
	return false
}

func (dag *Dag) isCyclic() bool {
	// Mark all the vertices as not visited, not part of recursion stack
	var v int = len(dag.nodes)
	visited := make([]bool, v)
	recStack := make([]bool, v)

	for i := 0; i < v; i++ {
		visited[i] = false
		recStack[i] = false
	}

	// Call the recursive helper function to detect cycle in different DFS trees
	for i := 0; i < v; i++ {
		if dag.isCyclicUtil(i, visited, recStack) {
			return true
		}
	}
	return false
}
