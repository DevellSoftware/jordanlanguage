package parse

type NodeCollection interface {
	Add(node Node)
	Children() []Node
}
