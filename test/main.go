package main

import (
	"fmt"

	"github.com/drakcoder/graphlib"
)

func main() {
	g := graphlib.NewGraph()
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
