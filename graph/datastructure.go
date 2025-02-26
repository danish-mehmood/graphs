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

func (g *Graph) doesEdgeExist(from, to string) bool {
	neighbors := g.GetNeighbors(from)
	if _, exists := neighbors[to]; exists {
		return true
	} else {
		return false
	}

}

func (g *Graph) IsPathValid(path []string) bool {
	if len(path) == 0 {
		return false
	} else {
		prevNode := path[0]
		var i int = 1
		for ; i < len(path); i++ {
			nextNode := path[i]
			// fmt.Println(prevNode, path[i])
			if g.doesEdgeExist(prevNode, nextNode) {
				prevNode = nextNode
				continue
			} else {
				return false
			}

		}
		return true
	}

}

func (g *Graph) ComputePathCost(path []string) (float64, error) {
	if !g.weighted {
		return 0, errors.New("the graph is not weighted")
	}
	if len(path) < 2 {
		return 0, errors.New("the path is too short")
	}

	sum := 0
	for i := 0; i < len(path)-1; i++ {
		from, to := path[i], path[i+1]

		// Directly access adjacency list
		weight, exists := g.adjList[from][to]
		if !exists {
			return 0, errors.New("the path does not exist")
		}

		sum += weight
	}

	return float64(sum), nil
}
