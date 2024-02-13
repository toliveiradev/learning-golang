package main

import "fmt"

// Define a struct for Graph
type Graph struct {
	vertices []*Vertex
}

// Define a struct for Vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Create a new graph
func NewGraph() *Graph {
	return &Graph{}
}

// Add a vertex to the graph
func (g *Graph) AddVertex(key int) {
	vertex := &Vertex{key: key}
	g.vertices = append(g.vertices, vertex)
}

// Add an edge between two vertices
func (g *Graph) AddEdge(k1, k2 int) {
	vertex1 := g.getVertex(k1)
	vertex2 := g.getVertex(k2)
	// add k2 to the adjacency list of k1
	vertex1.adjacent = append(vertex1.adjacent, vertex2)
	// add k1 to the adjacency list of k2
	vertex2.adjacent = append(vertex2.adjacent, vertex1)
}

// Get a vertex from the graph by its key
func (g *Graph) getVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

// Print the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %d -> ", v.key)
		for _, u := range v.adjacent {
			fmt.Printf("%d ", u.key)
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()

	// Add vertices
	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	// Add edges
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// Print the graph
	graph.Print()
}
