# Find Cheapest Flight with K Stops

## Overview
This Go program finds the cheapest flight from a source city to a destination city with at most `k` stops. It is based on the Bellman-Ford algorithm, which iteratively relaxes the edges of the graph to find the shortest path. The solution is designed for scenarios with up to 100 cities and various flight paths between them.

## Problem Statement
Given:
- `n` cities connected by flights, represented as an array `flights` where each flight is defined as `[from, to, price]`.
- Three integers `src` (starting city), `dst` (destination city), and `k` (maximum number of stops allowed).

The task is to return the cheapest price from `src` to `dst` with at most `k` stops. If no such route exists, return `-1`.

### Example

**Input:**
```plaintext
n = 4
flights = [[0, 1, 100], [1, 2, 100], [2, 0, 100], [1, 3, 600], [2, 3, 200]]
src = 0
dst = 3
k = 1
```

**Output:**
```plaintext
700
```

**Explanation:**
The optimal path from city `0` to city `3` with at most `1` stop is `0 -> 1 -> 3` with a cost of `100 + 600 = 700`.

## Solution Approach
- The program initializes an array to track the minimum cost to reach each city.
- It iteratively updates the costs for `k + 1` times to simulate the number of allowed stops.
- Each iteration relaxes all edges to find a cheaper path.

### Complexity
- **Time Complexity**: \( O(k 	imes E) \), where \( E \) is the number of flights.
- **Space Complexity**: \( O(n) \), for storing the cost array.

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
Modify the `main` function in `main.go` to change the inputs for `n`, `flights`, `src`, `dst`, and `k`.

## Code Structure
- `main.go`: Contains the core logic for finding the cheapest price from `src` to `dst` with at most `k` stops.
- `findCheapestPrice`: The main function implementing the algorithm.

## How It Works
1. The program initializes an array with maximum integer values to represent the minimum cost to each city.
2. It iterates up to `k + 1` times to update the cost using a temporary array.
3. At the end of the iterations, the minimum cost to reach the destination city `dst` is checked. If it remains the initial maximum value, `-1` is returned to indicate no path was found.
