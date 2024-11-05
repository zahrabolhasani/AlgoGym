# Breadth-First Search (BFS) in Golang

This project contains an implementation of the Breadth-First Search (BFS) algorithm in Golang. BFS is a fundamental graph traversal algorithm that can be used to explore the nodes of a graph layer by layer. This implementation is designed to work with an undirected graph, and it demonstrates how BFS can be used to traverse and explore graph data structures.

## Project Structure

- **main.go**: This is the main file that contains the implementation of the `Graph` struct and the `BFS` function.

## How to Run

1. **Install Go**: Ensure that Go (Golang) is installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

2. **Clone the Repository**: Clone this repository to your local machine using the following command:
   ```sh
   git clone <repository-url>
   ```

3. **Navigate to the Project Directory**:
   ```sh
   cd bfs-golang
   ```

4. **Run the Program**: Use the `go run` command to execute the program:
   ```sh
   go run main.go
   ```
   You should see the output of the BFS traversal starting from node 1.

## Example Output

The program runs a BFS traversal starting from node 1, with the following graph structure:

- Nodes: 1, 2, 3, 4, 5
- Edges: (1-2), (1-3), (2-4), (3-5), (4-5)

The expected output is:
```
BFS starting from node 1:
1 2 3 4 5
```
This output indicates that BFS successfully traversed the graph starting from node 1, visiting nodes in a breadth-wise manner.

## Code Overview

- **Graph Struct**: The `Graph` struct contains a `nodes` map that represents the adjacency list of the graph.
- **AddEdge Function**: The `AddEdge` function is used to add edges between nodes in the graph.
- **BFS Function**: The `BFS` function implements the Breadth-First Search algorithm using a queue to traverse nodes level by level. It prints each node as it is visited.

## Concepts Illustrated

- **Graph Representation**: The graph is represented as an adjacency list using a map in Go.
- **BFS Algorithm**: The Breadth-First Search algorithm is implemented using a queue to explore nodes, which ensures all neighbors are visited before moving to the next level.

## Requirements

- Go 1.16 or higher

## Running Tests

You can add more test cases to the main function to test different graph configurations and verify the BFS behavior.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributions

Feel free to submit pull requests or open issues to improve the code or add more features. Contributions are welcome!

## Contact

If you have any questions or suggestions regarding this project, please feel free to reach out.

---
Thank you for checking out this implementation of BFS in Golang!

