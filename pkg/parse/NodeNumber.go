package parse

import "strconv"

type NodeNumber struct {
	value float64
}

func NewNodeNumber(value float64) *NodeNumber {
	return &NodeNumber{value}
}

func (node *NodeNumber) ToNumber() *NodeNumber {
	return node
}

func (node *NodeNumber) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(node.value != 0)
}

func (node *NodeNumber) ToString() *NodeString {
	return NewNodeString(strconv.FormatFloat(node.value, 'f', -1, 64))
}

func (node *NodeNumber) ToArray() *NodeArray {
	array := NewNodeArray()
	array.Add(node)

	return array
}

func (node *NodeNumber) Children() []Node {
	return []Node{}
}

func (node *NodeNumber) Type() string {
	return "number"
}

func (node *NodeNumber) Value() float64 {
	return node.value
}
