package passagepathing

import (
	"fmt"
	"strings"
)

type Node struct {
	Edges   []string
	Visited bool
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(value string) bool {
	if _, ok := g.Nodes[value]; ok {
		return false
	}
	g.Nodes[value] = &Node{
		Edges:   []string{},
		Visited: false,
	}
	return true
}

func (g *Graph) AddEdge(value, edge string) {
	if _, ok := g.Nodes[value]; !ok {
		panic("Can't insert edge, missing node")
	}
	g.Nodes[value].Edges = append(g.Nodes[value].Edges, edge)
}

func (t *Graph) PrintGraph() {
	for k, v := range t.Nodes {
		fmt.Printf("Node: %s, Edges: %s\n", k, v.Edges)
	}
}

// Modified DFS
func (g *Graph) PathsWithOneSmallCave(n1, n2 string, counter *int) {
	// Allow cycles for uppercase letters
	if n1 != strings.ToUpper(n1) {
		g.Nodes[n1].Visited = true
	}

	if n1 == n2 {
		*counter += 1
	}

	for _, edge := range g.Nodes[n1].Edges {
		if !g.Nodes[edge].Visited {
			g.PathsWithOneSmallCave(edge, n2, counter)
		}
	}
	g.Nodes[n1].Visited = false
}

func (g *Graph) PathsWithTwoSmallCaves(n1, n2 string, paths []string, smallCaveCounter map[string]int, counter *int) {

	if n1 != strings.ToUpper(n1) {
		smallCaveCounter[n1] += 1
		visitedSmallCaves := 0
		// Can't visit these more then twice
		if n1 == "start" || n1 == "end" {
			g.Nodes[n1].Visited = true
		} else {
			for _, v := range smallCaveCounter {
				if v > 1 {
					visitedSmallCaves += 1
				}
				if v > 2 {
					smallCaveCounter[n1] -= 1
					return
				}
			}
			if visitedSmallCaves > 1 {
				smallCaveCounter[n1] -= 1
				return
			}
		}
	}

	if n1 == n2 {
		*counter += 1
	}

	for _, edge := range g.Nodes[n1].Edges {
		if !g.Nodes[edge].Visited {
			g.PathsWithTwoSmallCaves(edge, n2, paths, smallCaveCounter, counter)
		}
	}
	if n1 != strings.ToUpper(n1) {
		smallCaveCounter[n1] -= 1
	}
	// Backtrack start/end nodes
	g.Nodes[n1].Visited = false
}
