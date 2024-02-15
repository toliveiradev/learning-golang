package main

import "fmt"

// Define a struct for Graph
type Graph struct {
	vertices int
	edges    [][]int
}

// Create a new graph
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([][]int, vertices),
	}
}

// Add an undirected edge between u and v
func (g *Graph) AddEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

// Check if the graph is bipartite using BFS
func (g *Graph) IsBipartite() bool {
	colors := make([]int, g.vertices)
	for i := range colors {
		colors[i] = -1 // initialize with no color
	}

	queue := []int{}
	for i := 0; i < g.vertices; i++ {
		if colors[i] == -1 {
			colors[i] = 0
			queue = append(queue, i)

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				for _, neighbor := range g.edges[current] {
					if colors[neighbor] == -1 {
						colors[neighbor] = 1 - colors[current]
						queue = append(queue, neighbor)
					} else if colors[neighbor] == colors[current] {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	graph := NewGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	if graph.IsBipartite() {
		fmt.Println("Graph is bipartite")
	} else {
		fmt.Println("Graph is not bipartite")
	}
}
