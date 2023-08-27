package parse

type NodeParentheses struct {
	children []Node
}

func NewNodeParentheses() *NodeParentheses {
	return &NodeParentheses{
		children: make([]Node, 0),
	}
}

func (n *NodeParentheses) Children() []Node {
	return n.children
}

func (n *NodeParentheses) Add(child Node) {
	n.children = append(n.children, child)
}

func (n *NodeParentheses) ToString() *NodeString {
	return NewNodeString("()")
}

func (n *NodeParentheses) Type() string {
	return "parentheses"
}
