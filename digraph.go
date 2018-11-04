package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Digraph struct defines a map: string array -> node array
// Ex: [
//  	"A" -> [{B,4}, {C,3},
//		"B" -> {A,2}, {C,4},
// 		...
// ]
// This implies that A vertex has an edge leading to B with weight of 4, as well as edge
// leading to C with weight of 3. As so on..
type Digraph map[string][]Node

// Node struct defines a single connecting vertex with a weighted edge
type Node struct {
	edge   string
	weight int
}

// Route struct defines a route between two vertices
type Route struct {
	from string
	to   string
}

// BuildDigraphStruct of Digraph uses string input (ex: AB5, BC4, CD8 ...) to build a directed graph
// stored in memory as map of string -> node
func (dg *Digraph) BuildDigraphStruct(input string) (err error) {
	// Validate input and parse to data struct
	nodesSlice := strings.Split(input, ",")
	if len(nodesSlice) <= 1 {
		err = fmt.Errorf("Failed to parse train Directed graph input. Check input string")
		return
	}
	for _, v := range nodesSlice {
		saneNode := strings.TrimSpace(v)
		weight, e := strconv.Atoi(saneNode[2:])
		if e != nil {
			err = fmt.Errorf("Failed to parse Directed graph input. Expected weight to be integer. Error: %s", e.Error())
			return
		}

		head := string(saneNode[0])
		tail := string(saneNode[1])
		node := Node{edge: tail, weight: weight}
		dg.Push(head, node)
	}
	return
}

// Push function of Digraph pushes Node n to data struct under a head
func (dg *Digraph) Push(head string, n Node) {
	var nSlice []Node
	existingSlice := (*dg)[head]
	nSlice = append(existingSlice, n)
	(*dg)[head] = nSlice
}

// GetWeight of Digraph fetches immediate weight between two edges
// It returns false on the second argument if the node doesn't contain the vertex as a tail
func (dg *Digraph) GetWeight(head string, tail string) (weight int, found bool) {
	for _, m := range (*dg)[head] {
		if m.edge == tail {
			weight = m.weight
			found = true
			break
		}
	}
	return
}

// GetRouteWeight of Digraph fetches total weight of a given route (Ex: A-B-D-E)
// If the route doesn't exist, the second argument is returned as false
func (dg *Digraph) GetRouteWeight(trips []Route) (totalWeight int, found bool) {
	for _, trip := range trips {
		weight, f := dg.GetWeight(trip.from, trip.to)
		if !f {
			return
		}
		totalWeight += weight
	}
	found = true
	return
}

// GetTails of Digraph fetches all the tails of a given vertex
// Ex: C,D of A if A->C and A->D
func (dg *Digraph) GetTails(head string) (tails []string) {
	for _, v := range (*dg)[head] {
		tails = append(tails, string(v.edge))
	}
	return tails
}

// GetRoutesNumByDepth recursively fetches the number of possible routes
// from starting to end node in a defined max depth range
func (dg *Digraph) GetRoutesNumByDepth(start string, end string, depth int, sumRoutes *int) {
	if depth == 0 {
		return
	}
	towns := dg.GetTails(start)
	for _, town := range towns {
		if town == end {
			*sumRoutes++
		}
		dg.GetRoutesNumByDepth(town, end, depth-1, sumRoutes)
	}
}

// GetRoutesNumExactLength recursively fetches the number of possible routes
// from starting to end node matching the exact depth defined
func (dg *Digraph) GetRoutesNumExactLength(start string, end string, depth int, sumRoutes *int) {
	if depth == 0 {
		return
	}
	towns := dg.GetTails(start)
	for _, town := range towns {
		if town == end && depth == 1 {
			*sumRoutes++
		}
		dg.GetRoutesNumExactLength(town, end, depth-1, sumRoutes)
	}
}

// GetShortestPath is an attempt at apply Dijkstra on the Digraph implementation of
// weighted directed graphs
// This function is incomplete and needs some work to fix
func (dg *Digraph) GetShortestPath(start string, end string) (path []string) {

	// create distance map
	dist := map[string]int{}
	nodes := []string{}

	// set start node as 0 distance and infinity for the other nodes
	for k := range *dg {
		dist[k] = math.MaxInt32
		nodes = append(nodes, k)
	}
	dist[start] = 0

	// Go to each node from the starting vertex and discover it's neighbours
	// For every step - keep holding queues in `path` and pop vertices if an edge with a smaller weight is discovered
	var min int
	var nodesIndex int
	for len(nodes) > 0 {
		min = math.MaxInt32
		for k, elem := range nodes {
			print("elem: ", elem)
			if dist[elem] < min {
				min = dist[elem]
				nodesIndex = k
			}
		}
		nodes = append(nodes[:nodesIndex], nodes[nodesIndex+1:]...)

	}
	return
}

// DumpGraph is a debug function that dumps all key/values of map []string -> []Node
func (dg *Digraph) DumpGraph() {
	keys := reflect.ValueOf(*dg).MapKeys()

	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	for _, v := range strkeys {
		fmt.Printf("key %s, node: %v\n", v, (*dg)[v])
	}

}
