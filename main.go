package main

import (
	"fmt"

	"github.com/danish-mehmood/graphs/graph"
)

func main() {
	// Create an undirected weighted graph
	g := graph.NewGraph(graph.Undirected, true)

	// Add nodes
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")

	// Add edges
	g.AddEdge("A", "B", 5)
	g.AddEdge("B", "C", 3)
	g.AddEdge("C", "A", 2)

	// Print graph
	g.PrintGraph()

	// Get neighbors of a node
	fmt.Println("Neighbors of A:", g.GetNeighbors("A"))
}
