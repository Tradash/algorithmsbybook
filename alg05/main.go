package main

// алгоритме Дейкстры, стр 52
// Код подсмотрен https://medium.com/@rishabhmishra131/golang-dijkstra-algorithm-7bf2722ba0c8

import (
	"algoritmes/alg05/types"
	"fmt"
	"math"
	"sync"
)

type ItemGraph struct {
	Nodes []*types.Node
	Edges map[types.Node][]*types.Edge
	Lock  sync.RWMutex
}

func (g *ItemGraph) AddNode(n *types.Node) {
	g.Lock.Lock()
	g.Nodes = append(g.Nodes, n)
	g.Lock.Unlock()
}

func (g *ItemGraph) AddEdge(n1, n2 *types.Node, weight int) {
	g.Lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[types.Node][]*types.Edge)
	}
	ed1 := types.Edge{
		Node:   n2,
		Weight: weight,
	}
	ed2 := types.Edge{
		Node:   n1,
		Weight: weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
	g.Lock.Unlock()
}

type InputGraph struct {
	weight int
	from   string
	to     string
}

func createGraph(data []InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]*types.Node)
	for _, v := range data {
		if _, found := nodes[v.from]; !found {
			nA := types.Node{v.from}
			nodes[v.from] = &nA
			g.AddNode(&nA)

		}
		if _, found := nodes[v.to]; !found {
			nA := types.Node{v.to}
			nodes[v.to] = &nA
			g.AddNode(&nA)
		}
		g.AddEdge(nodes[v.from], nodes[v.to], v.weight)
	}
	return &g
}

func d(startNode *types.Node, endNode *types.Node, g *ItemGraph) ([]string, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)

	q := types.NodeQueue{}

	pq := q.NewQ()

	start := types.Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)

	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]
		for _, val := range near {
			if !visited[val.Node.Value] {
				if dist[v.Node.Value]+val.Weight < dist[val.Node.Value] {
					store := types.Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Value] + val.Weight,
					}
					dist[val.Node.Value] = dist[v.Node.Value] + val.Weight
					prev[val.Node.Value] = v.Node.Value
					pq.Enqueue(store)
				}
			}

		}
	}
	fmt.Println(dist)
	fmt.Println(prev)
	pathVal := prev[endNode.Value]
	var finalArr, rFA []string
	for pathVal != startNode.Value {
		finalArr = append(finalArr, pathVal)
		pathVal = prev[pathVal]
	}
	finalArr = append(finalArr, pathVal)
	fmt.Println(finalArr)

	for i := len(finalArr); i > 0; i-- {
		rFA = append(rFA, finalArr[i-1])
	}

	return append(rFA, endNode.Value), dist[endNode.Value]

}

func main() {
	inputGraph := []InputGraph{{3, "A", "B"}, {1, "A", "C"}, {6, "B", "F"}, {3, "B", "D"}, {1, "D", "F"}, {4, "C", "D"}, {2, "C", "E"}, {5, "E", "F"}, {2, "E", "F"}}
	z := createGraph(inputGraph)
	fmt.Println(z)
	short, i := d(&types.Node{"A"}, &types.Node{"F"}, z)
	fmt.Println(short, i)

}
