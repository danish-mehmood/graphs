package main

import (
	"fmt"

	"github.com/danish-mehmood/graphs/graph"
)

func main() {
	// Create an undirected weighted graph
	g := graph.NewGraph(graph.Undirected, true)

	// Add nodes
	g.AddNode("0")
	g.AddNode("1")
	g.AddNode("2")
	g.AddNode("3")
	g.AddNode("4")
	g.AddNode("5")
	g.AddNode("6")
	g.AddNode("7")

	// Add edges
	g.AddEdge("0", "1", 5)
	g.AddEdge("0", "3", 3)
	g.AddEdge("0", "5", 2)
	g.AddEdge("3", "1", 5)
	g.AddEdge("3", "0", 3)
	g.AddEdge("3", "5", 2)
	g.AddEdge("3", "2", 5)
	g.AddEdge("3", "4", 3)
	g.AddEdge("1", "2", 2)
	g.AddEdge("1", "3", 5)
	g.AddEdge("1", "0", 3)
	g.AddEdge("5", "3", 2)
	g.AddEdge("5", "0", 5)
	g.AddEdge("5", "4", 3)
	g.AddEdge("5", "6", 2)
	g.AddEdge("2", "1", 5)
	g.AddEdge("2", "3", 3)
	g.AddEdge("2", "4", 2)
	g.AddEdge("2", "7", 5)
	g.AddEdge("4", "2", 3)
	g.AddEdge("4", "3", 2)
	g.AddEdge("4", "5", 5)
	g.AddEdge("4", "6", 3)
	g.AddEdge("4", "7", 2)
	g.AddEdge("6", "4", 5)
	g.AddEdge("6", "5", 3)
	g.AddEdge("6", "7", 2)
	g.AddEdge("7", "2", 5)
	g.AddEdge("7", "4", 3)
	g.AddEdge("7", "6", 2)

	// Print graph
	g.PrintGraph()

	// Get neighbors of a node
	fmt.Println("Neighbors of A:", g.GetNeighbors("A"))

	pathvalid := g.IsPathValid([]string{"0", "1", "3", "2", "4", "7"})
	fmt.Println(pathvalid)
	cost, err := g.ComputePathCost([]string{"0", "1", "3", "2", "4", "7"})
	fmt.Println(cost, err)
}
