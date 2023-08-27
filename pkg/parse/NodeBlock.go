package parse

import "strings"

type NodeBlock struct {
	children []Node
}

func NewNodeBlock() *NodeBlock {
	return &NodeBlock{
		children: make([]Node, 0),
	}
}

func (this *NodeBlock) Children() []Node {
	return this.children
}

func (this *NodeBlock) Add(node Node) {
	this.children = append(this.children, node)
}

func printNode(node Node) string {
	result := make([]string, 0)
	result = append(result, "{")

	for _, node := range node.Children() {
		if node.Type() == "block" {
			result = append(result, "{")
			result = append(result, printNode(node))
			result = append(result, "}")
		} else {
			result = append(result, node.ToString().Value())
		}
	}

	result = append(result, "}")

	return strings.Join(result, " ")
}

func (this *NodeBlock) ToString() *NodeString {
	return NewNodeString(printNode(this))
}

func (this *NodeBlock) ToNumber() *NodeNumber {
	return NewNodeNumber(float64(len(this.children)))
}

func (this *NodeBlock) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(len(this.children) != 0)
}

func (this *NodeBlock) ToArray() *NodeArray {
	return NewNodeArray()
}

func (this *NodeBlock) Type() string {
	return "block"
}
