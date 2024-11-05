# Network Delay Time

This project implements the solution for the 'Network Delay Time' problem using Golang. It uses Dijkstra's algorithm to find the minimum time required for a signal to be transmitted from a given source node (`k`) to all other nodes in a directed, weighted graph.

## Problem Description

You are given a network of `n` nodes labeled from `1` to `n`. You are also given `times`, a list of travel times represented as directed edges `times[i] = (ui, vi, wi)`, where `ui` is the source node, `vi` is the target node, and `wi` is the time it takes for a signal to travel from `ui` to `vi`.

We will send a signal from a given node `k`. The goal is to determine the minimum time it takes for all `n` nodes to receive the signal. If it is impossible for all nodes to receive the signal, return `-1`.

### Example

Input:
```
times = [[2,1,1],[2,3,1],[3,4,1]]
n = 4
k = 2
```

Output:
```
2
```

### Constraints

- `1 <= k <= n <= 100`
- `1 <= times.length <= 6000`
- `times[i].length == 3`
- `1 <= ui, vi <= n`
- `ui != vi`
- `0 <= wi <= 100`

## Solution Approach

The solution uses Dijkstra's algorithm to find the shortest path from the source node to all other nodes in the graph. It maintains a min-heap (priority queue) to efficiently find the next node with the shortest distance to process.

### How to Run the Code

1. Make sure you have Golang installed on your machine.
2. Clone this repository or copy the code into a `.go` file.
3. Run the program using the following command:
   ```
   go run main.go
   ```

### Explanation of the Code

- **Graph Representation**: The graph is represented using an adjacency list where each node points to a list of its neighbors along with the time required to reach them.
- **Min-Heap**: A min-heap (implemented using Go's `container/heap` package) is used to always expand the node with the current minimum time.
- **Dijkstra's Algorithm**: We initialize the distance to the starting node (`k`) as `0` and all other nodes as infinity (`math.MaxInt32`). The algorithm then proceeds by visiting each node's neighbors and updating the shortest distance if a shorter path is found.

### Optimization Improvements

- **Graph Representation Optimization**: Changed the graph representation from a `map` to a slice of slices (`[][]Edge`) to slightly improve lookup times and reduce memory overhead.
- **Code Simplification**: The code was cleaned up for better readability, and unnecessary checks were removed to improve clarity.

### Example Usage

In the given example, the signal starts from node `2` and takes the following paths:
- From `2` to `1` with time `1`
- From `2` to `3` with time `1`
- From `3` to `4` with additional time `1`

Thus, the signal reaches all nodes in a minimum of `2` units of time.

### License

This project is licensed under the MIT License.
