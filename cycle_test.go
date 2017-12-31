package cycle

import (
	"testing"
)

type tv struct {
	id       string
	children []*tv
}

func (v *tv) Children() []Node {
	nodes := make([]Node, len(v.children))
	for i := range v.children {
		nodes[i] = Node(v.children[i])
	}
	return nodes
}
func TestGetNodesCycles_NoCycle(t *testing.T) {
	a := &tv{}
	b := &tv{}
	a.children = []*tv{b}

	type args struct {
		nodes []Node
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{
			args:    args{[]Node{a, b}, -1},
			wantNum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCycles := GetNodesCycles(tt.args.nodes, tt.args.n); len(gotCycles) != tt.wantNum {
				t.Errorf("GetNodesCycles() = %v, want count %v", gotCycles, tt.wantNum)
			}
		})
	}
}

func TestGetNodesCycles_Cycle1(t *testing.T) {
	a := &tv{}
	b := &tv{}
	a.children = []*tv{b}
	b.children = []*tv{a}

	type args struct {
		nodes []Node
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{
			args:    args{[]Node{a, b}, -1},
			wantNum: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCycles := GetNodesCycles(tt.args.nodes, tt.args.n); len(gotCycles) != tt.wantNum {
				t.Errorf("GetNodesCycles() = %v, want count %v", gotCycles, tt.wantNum)
			}
		})
	}
}

func TestGetNodesCycles_Cycle2(t *testing.T) {
	a := &tv{}
	b := &tv{}
	c := &tv{}
	a.children = []*tv{b, c}
	b.children = []*tv{a}
	c.children = []*tv{a}

	type args struct {
		nodes []Node
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{
			args:    args{[]Node{a, b, c}, -1},
			wantNum: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCycles := GetNodesCycles(tt.args.nodes, tt.args.n); len(gotCycles) != tt.wantNum {
				t.Errorf("GetNodesCycles() = %v, want count %v", gotCycles, tt.wantNum)
			}
		})
	}
}

func TestGetNodesCycles_Cycle3(t *testing.T) {
	a := &tv{}
	b := &tv{}
	c := &tv{}
	d := &tv{}
	a.children = []*tv{b, c}
	b.children = []*tv{a}
	c.children = []*tv{d}
	d.children = []*tv{c}

	type args struct {
		nodes []Node
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{
			args:    args{[]Node{a, b, c, d}, -1},
			wantNum: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCycles := GetNodesCycles(tt.args.nodes, tt.args.n); len(gotCycles) != tt.wantNum {
				t.Errorf("GetNodesCycles() = %v, want count %v", gotCycles, tt.wantNum)
			}
		})
	}
}
