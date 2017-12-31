package cycle

// Node is a v in G = (V, E)
type Node interface {
	Children() []Node
}

type node struct {
	index    int
	lowlink  int
	onStack  bool
	node     Node
	children []*node
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
		children := V[i].node.Children()
		innerChildren := make([]*node, len(children))
		for i, child := range children {
			innerChildren[i] = rec[child]
		}
		V[i].children = innerChildren
	}

	return V
}

// GetNodesCycles get cycles from graph data of adjacency lists with algorithm tarjan
func GetNodesCycles(nodes []Node, n int) (cycles [][]Node) {
	var (
		index         = 0
		stack         = make([]*node, 0, len(nodes))
		strongConnect func(*node)
	)

	strongConnect = func(v *node) {
		index++
		v.index = index
		v.lowlink = index
		stack = append(stack, v) //push
		v.onStack = true

		for _, w := range v.children {
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

	// algorithm tarjan begin
	V := wrapNodes(nodes)
	for _, v := range V {
		if 0 == v.index {
			strongConnect(v)
		}

		if n == len(cycles) {
			return
		}
	}

	return cycles
}
