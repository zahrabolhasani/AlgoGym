package main

import "fmt"

func isBipartite(graph [][]int) bool {
    n := len(graph)
    colors := make([]int, n)
    for i := range colors {
        colors[i] = -1 // Initial color, meaning unvisited
    }

    for i := 0; i < n; i++ {
        if colors[i] == -1 { // Start BFS/DFS from an unvisited node
            queue := []int{i}
            colors[i] = 0 // Start coloring with color 0

            for len(queue) > 0 {
                node := queue[0]
                queue = queue[1:]

                for _, neighbor := range graph[node] {
                    if colors[neighbor] == -1 {
                        colors[neighbor] = 1 - colors[node] // Assign the opposite color
                        queue = append(queue, neighbor)
                    } else if colors[neighbor] == colors[node] {
                        return false // Same color neighbors found, not bipartite
                    }
                }
            }
        }
    }
    return true
}

func main() {
    graph1 := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
    graph2 := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}

    fmt.Println(isBipartite(graph1)) // Output: false
    fmt.Println(isBipartite(graph2)) // Output: true
}
