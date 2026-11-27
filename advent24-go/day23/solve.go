package day23

import (
	"fmt"
	"sort"
	"strings"
)

type Graph struct {
	adj   map[string]map[string]struct{}
	nodes []string
}

func parseInput(input string) Graph {
	adj := make(map[string]map[string]struct{})
	nodeSet := make(map[string]struct{})
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		if adj[a] == nil {
			adj[a] = make(map[string]struct{})
		}
		if adj[b] == nil {
			adj[b] = make(map[string]struct{})
		}
		adj[a][b] = struct{}{}
		adj[b][a] = struct{}{}
		nodeSet[a] = struct{}{}
		nodeSet[b] = struct{}{}
	}
	nodes := make([]string, 0, len(nodeSet))
	for n := range nodeSet {
		nodes = append(nodes, n)
	}
	sort.Strings(nodes)
	return Graph{adj: adj, nodes: nodes}
}

func hasEdge(adj map[string]map[string]struct{}, u, v string) bool {
	_, ok := adj[u][v]
	return ok
}

func SolvePartOne(input string) int {
	g := parseInput(input)
	count := 0
	for i := 0; i < len(g.nodes); i++ {
		a := g.nodes[i]
		for j := i + 1; j < len(g.nodes); j++ {
			b := g.nodes[j]
			if !hasEdge(g.adj, a, b) {
				continue
			}
			for k := j + 1; k < len(g.nodes); k++ {
				c := g.nodes[k]
				if hasEdge(g.adj, a, c) && hasEdge(g.adj, b, c) {
					if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t") {
						count++
					}
				}
			}
		}
	}
	return count
}

func bronKerbosch(
	adj map[string]map[string]struct{},
	r []string,
	p []string,
	x []string,
	best *[]string,
) {
	if len(p) == 0 && len(x) == 0 {
		if len(r) > len(*best) {
			cp := make([]string, len(r))
			copy(cp, r)
			*best = cp
		}
		return
	}

	if len(r)+len(p) <= len(*best) {
		return
	}

	var pivot string
	maxNeighbors := -1
	union := make([]string, 0, len(p)+len(x))
	union = append(union, p...)
	union = append(union, x...)

	for _, u := range union {
		deg := 0
		for _, v := range p {
			if hasEdge(adj, u, v) {
				deg++
			}
		}
		if deg > maxNeighbors {
			maxNeighbors = deg
			pivot = u
		}
	}

	candidates := make([]string, 0, len(p))
	for _, v := range p {
		if !hasEdge(adj, pivot, v) {
			candidates = append(candidates, v)
		}
	}

	pSet := make(map[string]struct{}, len(p))
	for _, v := range p {
		pSet[v] = struct{}{}
	}
	xSet := make(map[string]struct{}, len(x))
	for _, v := range x {
		xSet[v] = struct{}{}
	}

	for _, v := range candidates {
		var nextR []string
		nextR = append(nextR, r...)
		nextR = append(nextR, v)

		var nextP []string
		for _, w := range p {
			if hasEdge(adj, v, w) {
				nextP = append(nextP, w)
			}
		}

		var nextX []string
		for _, w := range x {
			if hasEdge(adj, v, w) {
				nextX = append(nextX, w)
			}
		}

		bronKerbosch(adj, nextR, nextP, nextX, best)

		delete(pSet, v)
		xSet[v] = struct{}{}

		p = p[:0]
		for w := range pSet {
			p = append(p, w)
		}
		x = x[:0]
		for w := range xSet {
			x = append(x, w)
		}
	}
}

func SolvePartTwo(input string) int {
	g := parseInput(input)
	var best []string
	bronKerbosch(g.adj, nil, g.nodes, nil, &best)
	sort.Strings(best)
	fmt.Println("    Password:", strings.Join(best, ","))
	return len(best)
}
