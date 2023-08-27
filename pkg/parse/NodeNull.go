package parse

type NodeNull struct {
}

func NewNodeNull() *NodeNull {
	return &NodeNull{}
}

func (node *NodeNull) ToNumber() *NodeNumber {
	return NewNodeNumber(0)
}

func (node *NodeNull) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(false)
}

func (node *NodeNull) ToString() *NodeString {
	return NewNodeString("")
}

func (node *NodeNull) ToArray() *NodeArray {
	return NewNodeArray()
}

func (node *NodeNull) Children() []Node {
	return []Node{}
}

func (node *NodeNull) Type() string {
	return "null"
}
