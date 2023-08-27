package parse

type NodeBoolean struct {
	value bool
}

func NewNodeBoolean(value bool) *NodeBoolean {
	return &NodeBoolean{value}
}

func (node *NodeBoolean) ToNumber() *NodeNumber {
	if node.value {
		return NewNodeNumber(1)
	}

	return NewNodeNumber(0)
}

func (node *NodeBoolean) ToBoolean() *NodeBoolean {
	return node
}

func (node *NodeBoolean) ToString() *NodeString {
	if node.value {
		return NewNodeString("true")
	}

	return NewNodeString("false")
}

func (node *NodeBoolean) ToArray() *NodeArray {
	array := NewNodeArray()
	array.Add(node)

	return array
}

func (node *NodeBoolean) Children() []Node {
	return []Node{}
}

func (node *NodeBoolean) Type() string {
	return "boolean"
}

func (node *NodeBoolean) Value() bool {
	return node.value
}
