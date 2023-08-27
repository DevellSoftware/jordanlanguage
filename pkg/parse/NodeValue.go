package parse

type NodeValue interface {
	ToString() *NodeString
	ToBoolean() *NodeBoolean
	ToNumber() *NodeNumber
	ToArray() *NodeArray

	Node
}
