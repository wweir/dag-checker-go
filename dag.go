package dag

// Node is a v in G = (V, E)
type Node interface {
	Targets() []Node
}

type node struct {
	index   int
	lowlink int
	onStack bool
	node    Node
	targets []*node
}

// wrapNodes wrap outer Node into node
func wrapNodes(nodes []Node) []*node {
	V := make([]*node, len(nodes))
	rec := map[Node]*node{}

	for i := range nodes {
		V[i] = &node{node: nodes[i]}
		rec[nodes[i]] = V[i]
	}

	for i := range V {
		targets := V[i].node.Targets()
		wrapped := make([]*node, len(targets))
		for i, target := range targets {
			wrapped[i] = rec[target]
		}
		V[i].targets = wrapped
	}

	return V
}

// GetNodesCycles get cycles from graph data of adjacency lists with algorithm tarjan
func GetNodesCycles(nodes []Node, n int) (cycles [][]Node) {
	var (
		index         = 0 // 0 means undefined
		stack         = make([]*node, 0, len(nodes))
		strongConnect func(*node)
	)

	strongConnect = func(v *node) {
		index++
		v.index = index
		v.lowlink = index
		stack = append(stack, v) //push
		v.onStack = true

		for _, w := range v.targets {
			if 0 == w.index {
				strongConnect(w)
				if w.lowlink < v.lowlink {
					v.lowlink = w.lowlink
				}
			}

			if w.onStack {
				if w.lowlink < v.lowlink {
					v.lowlink = w.lowlink
				}
			}
		}

		if v.index == v.lowlink {
			strongComponent := []Node{}
			for i := len(stack) - 1; ; i-- {
				w := stack[i]
				strongComponent = append(strongComponent, w.node)

				w.onStack = false
				if v == stack[i] {
					stack = stack[:i] //pop
					break
				}
			}

			// skip one v strongComponent, do not follow tarjan way
			if 1 != len(strongComponent) {
				cycles = append(cycles, strongComponent)
			}
		}
	}

	// tarjan start
	V := wrapNodes(nodes)
	for _, v := range V {
		if 0 == v.index {
			strongConnect(v)
		}

		if n > 0 && n <= len(cycles) {
			return cycles[:n]
		}
	}

	return cycles
}
