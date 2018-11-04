package main

import (
	"fmt"
)

func main() {
	dg := Digraph{}

	InputString := "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
	dg.BuildDigraphStruct(InputString)

	outputOne := problemOne(dg)
	println(outputOne)

	outputTwo := problemTwo(dg)
	println(outputTwo)

	outputThree := problemThree(dg)
	println(outputThree)

	outputFour := problemFour(dg)
	println(outputFour)

	outputFive := problemFive(dg)
	println(outputFive)

	outputSix := problemSix(dg)
	println(outputSix)

	outputSeven := problemSeven(dg)
	println(outputSeven)

	outputEight := problemEight(dg)
	println(outputEight)

	outputNine := problemNine(dg)
	println(outputNine)

	outputTen := problemTen(dg)
	println(outputTen)

}

// 1. The distance of the route A-B-C.
func problemOne(dg Digraph) (output string) {
	totalDistance := 0

	trips := []Route{
		{"A", "B"},
		{"B", "C"},
	}

	totalDistance, found := dg.GetRouteWeight(trips)
	if !found {
		return "Output #1: NO SUCH ROUTE"
	}
	return fmt.Sprintf("Output #1: %d", totalDistance)
}

// 2. The distance of the route A-D.
func problemTwo(dg Digraph) (output string) {
	trips := []Route{
		{"A", "D"},
	}
	totalDistance, found := dg.GetRouteWeight(trips)
	if !found {
		return "Output #2: NO SUCH ROUTE"
	}
	return fmt.Sprintf("Output #2: %d", totalDistance)
}

// 3. The distance of the route A-D-C.
func problemThree(dg Digraph) (output string) {
	totalDistance := 0

	trips := []Route{
		{"A", "D"},
		{"D", "C"},
	}

	totalDistance, found := dg.GetRouteWeight(trips)
	if !found {
		return "Output #3: NO SUCH ROUTE"
	}
	return fmt.Sprintf("Output #3: %d", totalDistance)
}

// 4. The distance of the route A-E-B-C-D.
func problemFour(dg Digraph) (output string) {
	totalDistance := 0

	trips := []Route{
		{"A", "E"},
		{"E", "B"},
		{"B", "C"},
		{"C", "D"},
	}

	totalDistance, found := dg.GetRouteWeight(trips)
	if !found {
		return "Output #4: NO SUCH ROUTE"
	}
	return fmt.Sprintf("Output #4: %d", totalDistance)
}

// 5. The distance of the route A-E-D.
func problemFive(dg Digraph) (output string) {
	totalDistance := 0

	trips := []Route{
		{"A", "E"},
		{"E", "D"},
	}

	totalDistance, found := dg.GetRouteWeight(trips)
	if !found {
		return "Output #5: NO SUCH ROUTE"
	}
	return fmt.Sprintf("Output #5: %d", totalDistance)
}

// 6. The number of trips starting at C and ending at C with a maximum of 3
// stops.  In the sample data below, there are two such trips: C-D-C (2
// stops). and C-E-B-C (3 stops).
func problemSix(dg Digraph) (output string) {
	sumRoutes := 0
	dg.GetRoutesNumByDepth("C", "C", 3, &sumRoutes)

	return fmt.Sprintf("Output #6: %d", sumRoutes)
}

// 7. The number of trips starting at A and ending at C with exactly 4 stops.
// In the sample data below, there are three such trips: A to C (via B,C,D); A
// to C (via D,C,D); and A to C (via D,E,B).
func problemSeven(dg Digraph) (output string) {
	sumRoutes := 0
	dg.GetRoutesNumExactLength("A", "C", 4, &sumRoutes)

	return fmt.Sprintf("Output #7: %d", sumRoutes)
}

// 8. The length of the shortest route (in terms of distance to travel) from A
// to C.
func problemEight(dg Digraph) (output string) {
	// _ = dg.CalculateShortestPath("A", "C")
	return fmt.Sprintf("Output #8: %s", "INCOMPLETE")
}

// 9. The length of the shortest route (in terms of distance to travel) from B
// to B.
func problemNine(dg Digraph) (output string) {
	return fmt.Sprintf("Output #9: %s", "INCOMPLETE")
}

// 10. The number of different routes from C to C with a distance of less than
// 30.  In the sample data, the trips are: CDC, CEBC, CEBCDC, CDCEBC, CDEBC,
// CEBCEBC, CEBCEBCEBC.
func problemTen(dg Digraph) (output string) {
	return fmt.Sprintf("Output #10: %s", "INCOMPLETE")
}
