package main

import (
    "fmt"
)

// Graph struct to represent the graph
type Graph struct {
    nodes map[int][]int
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v int) {
    g.nodes[u] = append(g.nodes[u], v)
    g.nodes[v] = append(g.nodes[v], u)
}

// BFS implementation of BFS in Golang
func (g *Graph) BFS(start int) {
    visited := make(map[int]bool)    // Visited nodes
    queue := []int{start}            // Queue initialized with the starting node

    visited[start] = true

    for len(queue) > 0 {
        // Dequeue the first element from the queue
        current := queue[0]
        queue = queue[1:]

        // Process the current node
        fmt.Printf("%d ", current)

        // Traverse all neighbors
        for _, neighbor := range g.nodes[current] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}

func main() {
    g := Graph{nodes: make(map[int][]int)}

    // Add edges to the graph
    g.AddEdge(1, 2)
    g.AddEdge(1, 3)
    g.AddEdge(2, 4)
    g.AddEdge(3, 5)
    g.AddEdge(4, 5)

    fmt.Println("BFS starting from node 1:")
    g.BFS(1)
	fmt.Println("2222")
}