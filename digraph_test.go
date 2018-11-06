package main

import (
	"strings"
	"testing"
)

// This test case tests itself and implicitly tests the 'Push()' function
func TestBuildDigraphStruct(t *testing.T) {
	// Test failure when empty string provided
	dg := Digraph{}
	stringOne := ""
	err := dg.BuildDigraphStruct(stringOne)
	if err == nil {
		t.Error("Expected error but found nil. ")
	}
	if !strings.Contains(err.Error(), "Check input string") {
		t.Errorf("Expected error message to be related to input string but it didn't")
		t.Logf("err is %s", err.Error())
	}

	// Test failure when converting to number
	dg = Digraph{}
	stringTwo := "AB3, ACC"
	err = dg.BuildDigraphStruct(stringTwo)
	if err == nil {
		t.Error("Expected error but found nil. ")
	}

	if !strings.Contains(err.Error(), "Expected weight to be integer") {
		t.Error("Failure when confirming conversion to integer")
		t.Logf("err is %s", err.Error())
	}

	// Test a good case
	dg = Digraph{}
	stringThree := "AB3, AC3, BC4"
	err = dg.BuildDigraphStruct(stringThree)
	if err != nil {
		t.Errorf("Expected no errors. Found %s", err.Error())
	}

	node0 := (dg)["A"][0]
	if node0.edge != "B" || node0.weight != 3 {
		t.Errorf("Expected index 0 of 'A' head to be B:3 but found: %v", node0)

	}

	node1 := (dg)["A"][1]
	if node1.edge != "C" || node1.weight != 3 {
		t.Errorf("Expected index 1 of 'A' head to be C:3 but found: %v", node1)
	}
}

func TestGetWeight(t *testing.T) {
	dg := Digraph{}
	stringOne := "AC4, AD255"
	err := dg.BuildDigraphStruct(stringOne)
	if err != nil {
		t.Errorf("Error creating a new Digraph. Error: %s", err.Error())
	}

	// test good cases
	weight, found := dg.GetWeight("A", "C")
	if weight != 4 || found != true {
		t.Errorf("Failure fetching correct weight values but 'found': %t and 'weight': %d", found, weight)
	}

	weight, found = dg.GetWeight("A", "D")
	if weight != 255 || found != true {
		t.Errorf("Failure fetching correct weight values but 'found': %t and 'weight': %d", found, weight)
	}

	// test not found case
	weight, found = dg.GetWeight("A", "E")
	if weight != 0 || found != false {
		t.Errorf("Checking against unprovided node but 'found': %t and 'weight': %d", found, weight)
	}
}

func TestGetRouteWeight(t *testing.T) {
	dg := Digraph{}
	stringOne := "AB3, BC3, CD3, BD1"
	err := dg.BuildDigraphStruct(stringOne)
	if err != nil {
		t.Errorf("Error creating a new Digraph. Error: %s", err.Error())
	}

	// test good case
	trips := []Route{
		{"A", "B"},
		{"A", "D"},
	}
	totalDistance, found := dg.GetRouteWeight(trips)
	if totalDistance != 3 || found != false {
		t.Errorf("Failure fetching correct total route values. 'found': %t and 'totalDistance': %d", found, totalDistance)
	}

	// test not found case
	trips = []Route{
		{"A", "B"},
		{"A", "N"},
	}
	totalDistance, found = dg.GetRouteWeight(trips)
	if found != false {
		t.Errorf("Checking against unprovided node but 'found': %t", found)
	}
}

func TestGetTails(t *testing.T) {
	dg := Digraph{}
	stringOne := "AB3, AC3, CD3, BD1"
	err := dg.BuildDigraphStruct(stringOne)
	if err != nil {
		t.Errorf("Error creating a new Digraph. Error: %s", err.Error())
	}

	tails := dg.GetTails("A")
	containsB := false
	containsC := false
	for _, tail := range tails {
		if tail == "B" {
			containsB = true
		}
		if tail == "C" {
			containsC = true
		}
	}
	if containsB == false {
		t.Errorf("Expected tails to contain 'B' for graph: '%s'", stringOne)
	}
	if containsC == false {
		t.Errorf("Expected tails to contain 'C' for graph: '%s'", stringOne)
	}
}

func TestGetRoutesNumByDepth(t *testing.T) {
	dg := Digraph{}
	stringOne := "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
	err := dg.BuildDigraphStruct(stringOne)
	if err != nil {
		t.Errorf("Error creating a new Digraph. Error: %s", err.Error())
	}

	sumRoutes := 0
	dg.GetRoutesNumByDepth("C", "C", 3, &sumRoutes)
	if sumRoutes != 2 {
		t.Errorf("Expected C->C route with max depth 3 to end up with `2` but got '%d' for graph: '%s'", sumRoutes, stringOne)
	}
}

func TestGetRoutesNumExactLength(t *testing.T) {
	dg := Digraph{}
	stringOne := "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
	err := dg.BuildDigraphStruct(stringOne)
	if err != nil {
		t.Errorf("Error creating a new Digraph. Error: %s", err.Error())
	}

	sumRoutes := 0
	dg.GetRoutesNumExactLength("A", "C", 4, &sumRoutes)
	if sumRoutes != 3 {
		t.Errorf("Expected A->C route with exact depth 4 to end up with `3` but got '%d' for graph: '%s'", sumRoutes, stringOne)
	}
}

func TestGetShortestPath(t *testing.T) {
	dg := Digraph{}
	stringOne := "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
	err := dg.BuildDigraphStruct(stringOne)
	if err != nil {
		t.Errorf("Error creating a new Digraph. Error: %s", err.Error())
	}
	total := dg.GetShortestPath("A", "C")
	if total != 9 {
		t.Errorf("Expected route A->C shortest path distance to equal 9. Found: %d", total)
	}

}
