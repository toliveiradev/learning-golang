package main

import "fmt"

// Define a struct for Edge
type Edge struct {
	start, end int
	weight     int
}

// Define a struct for Graph
type Graph struct {
	edges []*Edge
}

// Create a new graph
func NewGraph() *Graph {
	return &Graph{}
}

// Add an edge to the graph
func (g *Graph) AddEdge(start, end, weight int) {
	edge := &Edge{start: start, end: end, weight: weight}
	g.edges = append(g.edges, edge)
}

func main() {
	graph := NewGraph()

	// Add weighted edges
	graph.AddEdge(0, 1, 5)
	graph.AddEdge(0, 2, 7)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(2, 3, 8)

	// Print the graph
	for _, edge := range graph.edges {
		fmt.Printf("Edge %d-%d with weight %d\n", edge.start, edge.end, edge.weight)
	}
}
