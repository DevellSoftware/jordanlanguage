package parse

type NodeLineEnd struct {
}

func NewNodeLineEnd() *NodeLineEnd {
	return &NodeLineEnd{}
}

func (n *NodeLineEnd) Children() []Node {
	return []Node{}
}

func (n *NodeLineEnd) ToString() *NodeString {
	return NewNodeString("\\n")
}

func (n *NodeLineEnd) Type() string {
	return "lineEnd"
}
