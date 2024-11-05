package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	node, time int
}

type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	// Build the graph as an adjacency list
	graph := make([][]Edge, n+1)
	for _, time := range times {
		source, target, travelTime := time[0], time[1], time[2]
		graph[source] = append(graph[source], Edge{target, travelTime})
	}

	// Initialize distances with a large value (infinity)
	distances := make([]int, n+1)
	for i := 1; i <= n; i++ {
		distances[i] = math.MaxInt32
	}
	// Set the distance to the start node to zero
	distances[k] = 0

	// Priority queue (Min-Heap) to keep track of nodes and their arrival times
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Edge{k, 0})

	// Run Dijkstra's algorithm
	for minHeap.Len() > 0 {
		currentEdge := heap.Pop(minHeap).(Edge)
		node, currentTime := currentEdge.node, currentEdge.time

		// Skip if the current time is greater than the stored time
		if currentTime > distances[node] {
			continue
		}

		// Update the neighbors of the current node
		for _, neighbor := range graph[node] {
			nextNode, travelTime := neighbor.node, neighbor.time
			newTime := currentTime + travelTime
			// If the new time is shorter, update it
			if newTime < distances[nextNode] {
				distances[nextNode] = newTime
				heap.Push(minHeap, Edge{nextNode, newTime})
			}
		}
	}

	// Find the maximum time for the signal to reach all nodes
	maxTime := 0
	for i := 1; i <= n; i++ {
		if distances[i] == math.MaxInt32 {
			// If any node is unreachable, return -1
			return -1
		}
		if distances[i] > maxTime {
			maxTime = distances[i]
		}
	}

	return maxTime
}

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n, k := 4, 2
	result := networkDelayTime(times, n, k)
	fmt.Println(result)
}