# graphlib
Graphlib is a library for golang for working with graphs along with a few basic graph algorithms. Features included are:

* Creating a Graph
* To add links to the Nodes of the graphs
* Finding the shortest distance for all the nodes from a single source Node
* Finding the path and shortest distance from one node to another node

## Install
`go get github.com/drakcoder/graphlib`

## Example Usage

For creating a graph and finding the shortest distance between any 2 nodes that are existing in the graph.

``` package main

import (
	"fmt"

	"github.com/drakcoder/graphlib"
)

func main() {
	g := graphlib.NewGraph("directed")
	g.AddNodes("a", "b", "c", "d", "e")
	g.AddLink("a", "b", 6)
	g.AddLink("d", "a", 1)
	g.AddLink("b", "e", 2)
	g.AddLink("b", "d", 1)
	g.AddLink("c", "e", 5)
	g.AddLink("c", "b", 5)
	g.AddLink("e", "d", 1)
	g.AddLink("e", "c", 4)
	// dist, prev := g.Dijkstra("a")
	path, dist := g.DistBetn("a", "e")
	fmt.Println(dist, path)
} 
```
## Functionalities
**NewGraph()**
is used to create a new graph. this function creates and returns a pointer for the graph.
``` package main

import (
	"github.com/drakcoder/graphlib"
)

func main() {
	g := graphlib.NewGraph("directed")
} 
```
Take Note that the the argument in the function can be either "directed" or "undirected" and must be present.

**AddNodes()**
is used to add new Nodes to the graph before creating any links for them.

**AddLink()**
is used to add links to the existing Nodes along with the cost to traverse.

**DistBetn()** 
is used to find the shortest distance to traverse from a source node to a destination node.
The function returns two values. 
- an array of strings referring to the order of the traversal
- the cost of traversal

**Dijkstra()**
is used to find the shortest distance from a source Node to all other Nodes
the function returns two values
- a map of shortest distances
- a map of previous nodes
