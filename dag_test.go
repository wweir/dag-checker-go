package dag

import (
	"testing"
)

type tv struct {
	id      string
	targets []*tv
}

func (v *tv) Targets() []Node {
	nodes := make([]Node, len(v.targets))
	for i := range v.targets {
		nodes[i] = Node(v.targets[i])
	}
	return nodes
}

/*
A+--------->B
*/
func TestGetNodesCycles_NoCycle(t *testing.T) {
	a := &tv{}
	b := &tv{}
	a.targets = []*tv{b}

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

/*
A+-------------+
^              |
|              |
|              v
+-------------+B
*/
func TestGetNodesCycles_Cycle1(t *testing.T) {
	a := &tv{}
	b := &tv{}
	a.targets = []*tv{b}
	b.targets = []*tv{a}

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

/*
+------------+A+-------------+
|             ^              |
|             |              |
v             |              v
B+------------+-------------+C
*/
func TestGetNodesCycles_Cycle2(t *testing.T) {
	a := &tv{}
	b := &tv{}
	c := &tv{}
	a.targets = []*tv{b, c}
	b.targets = []*tv{a}
	c.targets = []*tv{a}

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

/*
+------------+A+-------------+-----------+D
|             ^              |            ^
|             |              |            |
v             |              v            |
B+------------+              C+-----------+
*/
func TestGetNodesCycles_Cycle3(t *testing.T) {
	a := &tv{}
	b := &tv{}
	c := &tv{}
	d := &tv{}
	a.targets = []*tv{b, c}
	b.targets = []*tv{a}
	c.targets = []*tv{d}
	d.targets = []*tv{c}

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
