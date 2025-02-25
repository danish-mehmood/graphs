package graph

import (
	"errors"
	"fmt"
)

type graphType int

const (
	Undirected graphType = iota
	Directed
)

type Graph struct {
	adjList   map[string]map[string]int
	graphType graphType
	weighted  bool
}

func NewGraph(graphType graphType, weighted bool) *Graph {
	return &Graph{
		adjList:   make(map[string]map[string]int),
		graphType: graphType,
		weighted:  weighted,
	}
}

func (g *Graph) AddNode(node string) {
	if _, exists := g.adjList[node]; !exists {
		g.adjList[node] = make(map[string]int)
	}
}

func (g *Graph) AddEdge(from, to string, weight int) error {
	if !g.weighted {
		weight = 1
	}
	if _, exists := g.adjList[from]; !exists {
		return errors.New("source node does not exist")
	}
	if _, exists := g.adjList[to]; !exists {
		return errors.New("destination node does not exist")
	}
	g.adjList[from][to] = weight
	if g.graphType == Undirected {
		g.adjList[to][from] = weight
	}
	return nil
}

// RemoveEdge removes an edge between two nodes
func (g *Graph) RemoveEdge(from, to string) error {
	if _, exists := g.adjList[from]; !exists {
		return errors.New("source node does not exist")
	}
	if _, exists := g.adjList[to]; !exists {
		return errors.New("destination node does not exist")
	}

	delete(g.adjList[from], to)
	if g.graphType == Undirected {
		delete(g.adjList[to], from)
	}
	return nil
}

// GetNeighbors returns the neighbors of a given node
func (g *Graph) GetNeighbors(node string) map[string]int {
	return g.adjList[node]
}

// PrintGraph prints the adjacency list representation of the graph
func (g *Graph) PrintGraph() {
	for node, edges := range g.adjList {
		fmt.Printf("%s -> ", node)
		for neighbor, weight := range edges {
			fmt.Printf("(%s, %d) ", neighbor, weight)
		}
		fmt.Println()
	}
}
