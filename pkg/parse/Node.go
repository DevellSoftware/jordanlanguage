package parse

type Node interface {
	Children() []Node
	ToString() *NodeString
	Type() string
}
