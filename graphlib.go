package graphlib

type Node struct {
	Name  string
	links []Edge
}

type Edge struct {
	from *Node
	to   *Node
	cost uint
}

type Graph struct {
	nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{nodes: map[string]*Node{}}
}

func (g *Graph) AddNodes(names ...string) {
	for _, name := range names {
		if _, ok := g.nodes[name]; !ok {
			g.nodes[name] = &Node{Name: name, links: []Edge{}}
		}
	}
}

func (g *Graph) AddLink(a, b string, cost int) {
	aNode := g.nodes[a]
	bNode := g.nodes[b]
	aNode.links = append(aNode.links, Edge{from: aNode, to: bNode, cost: uint(cost)})
}

func (g *Graph) DistBetn(source string, destination string) ([]string, uint) {
	dist, prev := map[string]uint{}, map[string]string{}

	for _, node := range g.nodes {
		dist[node.Name] = INFINITY
		prev[node.Name] = ""
	}
	visited := map[string]bool{}
	dist[source] = 0
	for u := source; u != ""; u = getClosestNonVisitedNode(dist, visited) {
		if source == destination {
			break
		}
		// fmt.Println(u)
		uDist := dist[u]
		for _, link := range g.nodes[u].links {
			if _, ok := visited[link.to.Name]; ok {
				continue
			}
			alt := uDist + link.cost
			v := link.to.Name
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
		visited[u] = true
	}
	var path []string
	cur := destination
	for cur != "" {
		path = append(path, cur)
		cur = prev[cur]
	}
	path = reverse(path)
	return path, dist[destination]
}

const INFINITY = ^uint(0)

func (g *Graph) Dijkstra(source string) (map[string]uint, map[string]string) {
	dist, prev := map[string]uint{}, map[string]string{}

	for _, node := range g.nodes {
		dist[node.Name] = INFINITY
		prev[node.Name] = ""
	}
	visited := map[string]bool{}
	dist[source] = 0
	for u := source; u != ""; u = getClosestNonVisitedNode(dist, visited) {
		// fmt.Println(u)
		uDist := dist[u]
		for _, link := range g.nodes[u].links {
			if _, ok := visited[link.to.Name]; ok {
				continue
			}
			alt := uDist + link.cost
			v := link.to.Name
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
		visited[u] = true
	}
	return dist, prev
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func getClosestNonVisitedNode(dist map[string]uint, visited map[string]bool) string {
	lowestCost := INFINITY
	lowestNode := ""
	for key, dis := range dist {
		if _, ok := visited[key]; dis == INFINITY || ok {
			continue
		}
		if dis < lowestCost {
			lowestCost = dis
			lowestNode = key
		}
	}
	// fmt.Println(lowestNode)
	return lowestNode
}