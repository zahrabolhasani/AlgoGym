# Bipartite Graph Checker

## Overview
This Go program checks if an undirected graph is **bipartite**. A graph is **bipartite** if the nodes can be divided into two independent sets **A** and **B** such that every edge in the graph connects a node in set **A** to a node in set **B**. The solution uses a **graph coloring approach** to determine if the graph can be divided in this way.

## Problem Statement
You are given an undirected graph with **n** nodes, where each node is numbered between **0** and **n - 1**. The graph is represented by a 2D array, `graph`, where `graph[u]` is an array of nodes adjacent to node **u**. The graph has the following properties:

- There are no self-edges (i.e., `graph[u]` does not contain **u**).
- There are no parallel edges (i.e., `graph[u]` does not contain duplicate values).
- If **v** is in `graph[u]`, then **u** is also in `graph[v]` (the graph is undirected).
- The graph may not be connected, meaning there may be nodes **u** and **v** such that there is no path between them.

The task is to determine if the graph is **bipartite**.

### Example

**Input:**
```plaintext
graph = [[1, 2, 3], [0, 2], [0, 1, 3], [0, 2]]
```
**Output:**
```plaintext
false
```
**Explanation:**
There is no way to partition the nodes into two independent sets such that every edge connects a node in one set to a node in the other.

**Input:**
```plaintext
graph = [[1, 3], [0, 2], [1, 3], [0, 2]]
```
**Output:**
```plaintext
true
```
**Explanation:**
We can partition the nodes into two sets: `{0, 2}` and `{1, 3}`.

## Solution Approach
The solution uses a **graph coloring technique** to determine if the graph is bipartite:

1. **Initialization**: We start by creating an array `colors` to store the color of each node. All nodes are initially unvisited (represented by **-1**).

2. **BFS Traversal**: For each unvisited node, we use **Breadth-First Search (BFS)** to traverse the graph:
   - Assign the starting node a color (e.g., color **0**).
   - For each neighboring node, assign the opposite color.
   - If we encounter a neighbor with the same color as the current node, then the graph is **not bipartite**.

3. **Multiple Components**: Since the graph may be disconnected, we repeat the BFS for each unvisited node to ensure all components are checked.

### Complexity
- **Time Complexity**: \( O(V + E) \), where \( V \) is the number of nodes and \( E \) is the number of edges. This is because we traverse all nodes and edges using BFS.
- **Space Complexity**: \( O(V) \), to store the color of each node and the queue for BFS.

## Code
```go
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
```

## Getting Started

### Prerequisites
- Go (Golang) installed on your machine.
- Basic knowledge of Go programming and data structures.

### Running the Code
1. Clone this repository or copy the code into a local directory.
2. Compile and run the code using the following commands:

```bash
go run main.go
```

### Example Usage
Modify the `main` function in `main.go` to test different graph inputs.

## How It Works
1. The program initializes an array `colors` with all values set to **-1**, indicating that no node has been visited.
2. It iterates over all nodes, starting a **BFS** for each unvisited node to ensure that disconnected components are also checked.
3. During the BFS, nodes are colored alternately, and any conflict indicates that the graph is not bipartite.

## Contributing
Feel free to submit issues or pull requests to improve the code or add new features.

## License
This project is open-source and available under the MIT License.

