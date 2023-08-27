package parse

import "fmt"

type NodeArrayElement struct {
	key   NodeValue
	value NodeValue
}

type NodeArray struct {
	elements []*NodeArrayElement
}

func NewNodeArray() *NodeArray {
	return &NodeArray{
		elements: make([]*NodeArrayElement, 0),
	}
}

func (this *NodeArray) ToString() *NodeString {
	return NewNodeString(fmt.Sprint(len(this.elements)))
}

func (this *NodeArray) ToNumber() *NodeNumber {
	return NewNodeNumber(float64(len(this.elements)))
}

func (this *NodeArray) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(len(this.elements) != 0)
}

func (this *NodeArray) ToArray() *NodeArray {
	return this
}

func (this *NodeArray) Children() []Node {
	return []Node{}
}

func (this *NodeArray) Add(node NodeValue) {
	this.elements = append(this.elements, &NodeArrayElement{
		key:   NewNodeNumber(float64(len(this.elements))),
		value: node,
	})
}

func (this *NodeArray) AddKey(key NodeValue, node NodeValue) {
	this.elements = append(this.elements, &NodeArrayElement{key, node})
}

func (this *NodeArray) Get(index NodeValue) NodeValue {
	if index.ToNumber().value < 0 || index.ToNumber().value >= float64(len(this.elements)) {
		return NewNodeUndefined()
	}

	return this.elements[int(index.ToNumber().value)].value
}

func (this *NodeArray) Type() string {
	return "array"
}

func (this *NodeArray) Value() []*NodeArrayElement {
	return this.elements
}
