package parse

type NodeUndefined struct {
}

func NewNodeUndefined() *NodeUndefined {
	return &NodeUndefined{}
}

func (node *NodeUndefined) ToNumber() *NodeNumber {
	return NewNodeNumber(0)
}

func (node *NodeUndefined) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(false)
}

func (node *NodeUndefined) ToString() *NodeString {
	return NewNodeString("undefined")
}

func (node *NodeUndefined) ToArray() *NodeArray {
	return NewNodeArray()
}

func (node *NodeUndefined) Children() []Node {
	return []Node{}
}

func (node *NodeUndefined) Type() string {
  return "undefined"
}
