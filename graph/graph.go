package graph

import "fmt"

type Item interface{}

// 组成图的顶点
type Node struct {
	Value Item
}

// 定义一个图的结构, 图有顶点与边组成 V  E
type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// 添加节点
func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// 添加边
func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	// 无向图
	g.Edges[*n1] = append(g.Edges[*n1], n2) // 设定从节点n1到n2的边
	g.Edges[*n2] = append(g.Edges[*n2], n1) // 设定从节点n2到n1的边
}

func (g *Graph) String() {
	s := ""

	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + "->"
		near := g.Edges[*g.Nodes[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}
