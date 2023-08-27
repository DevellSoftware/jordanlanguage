package parse

import "strconv"

type NodeString struct {
	value string
}

func NewNodeString(value string) *NodeString {
	return &NodeString{value}
}

func (node *NodeString) ToString() *NodeString {
	return node
}

func (node *NodeString) ToNumber() *NodeNumber {
	num, err := strconv.ParseFloat(node.value, 64)

	if err != nil {
		return NewNodeNumber(0)
	}

	return NewNodeNumber(num)
}

func (node *NodeString) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(node.value != "")
}

func (node *NodeString) ToArray() *NodeArray {
	array := NewNodeArray()

	array.Add(node)

	return array
}

func (node *NodeString) Children() []Node {
	return []Node{}
}

func (node *NodeString) Value() string {
	return node.value
}

func (node *NodeString) Type() string {
	return "string"
}
