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
	nodes     map[string]*Node
	exists    map[string]bool
	GraphType string
}

func NewGraph(GT string) *Graph {
	if GT != "undirected" && GT != "directed" {
		panic("The argument for graph creation can be \"directed\" or \"undirected\" only")
	}
	return &Graph{nodes: map[string]*Node{}, exists: map[string]bool{}, GraphType: GT}
}

func (g *Graph) AddNodes(names []string) {
	for _, name := range names {
		if _, ok := g.nodes[name]; !ok {
			g.nodes[name] = &Node{Name: name, links: []Edge{}}
			g.exists[name] = true
		}
	}
}

func (g *Graph) AddLink(a, b string, cost int) {
	aNode := g.nodes[a]
	bNode := g.nodes[b]
	if aNode == nil || bNode == nil {
		panic("creating edge for node that does not exist!")
	}
	aNode.links = append(aNode.links, Edge{from: aNode, to: bNode, cost: uint(cost)})
	if g.GraphType == "undirected" {
		bNode.links = append(bNode.links, Edge{from: bNode, to: aNode, cost: uint(cost)})
	}
}

func (g *Graph) DistBetn(source string, destination string) ([]string, uint) {
	dist, prev := map[string]uint{}, map[string]string{}
	var path []string
	if !g.exists[source] || !g.exists[destination] {
		panic("one of the nodes does not exist!")
	}
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
		currDist := dist[u]
		for _, link := range g.nodes[u].links {
			if _, ok := visited[link.to.Name]; ok {
				continue
			}
			alt := currDist + link.cost
			v := link.to.Name
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
		visited[u] = true
	}
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
		currDist := dist[u]
		for _, link := range g.nodes[u].links {
			if _, ok := visited[link.to.Name]; ok {
				continue
			}
			alt := currDist + link.cost
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
	return lowestNode
}

func (g *Graph) TopologicalSort() []string {
	in_degree := make(map[string]int)
	for name := range g.nodes {
		for _, link := range g.nodes[name].links {
			if _, ok := in_degree[name]; !ok {
				in_degree[link.to.Name] = 1
			} else {
				in_degree[link.to.Name]++
			}
		}
	}
	var q []string
	for name := range g.nodes {
		if in_degree[name] == 0 {
			q = append(q, name)
		}
	}
	cnt := 0
	var result []string
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		result = append(result, cur)
		for _, link := range g.nodes[cur].links {
			in_degree[link.to.Name]--
			if in_degree[link.to.Name] == 0 {
				q = append(q, link.to.Name)
			}
		}
		cnt++
	}
	if cnt != len(g.nodes) {
		panic("there exists a cycle in the graph")
	}
	return result
}

//test
