package parse

type NodeIdentifier struct {
	identifier string
}

func NewNodeIdentifier(identifier string) *NodeIdentifier {
	return &NodeIdentifier{
		identifier: identifier,
	}
}

func (n *NodeIdentifier) Identifier() string {
	return n.identifier
}

func (n *NodeIdentifier) ToString() *NodeString {
	return NewNodeString(n.identifier)
}

func (n *NodeIdentifier) ToNumber() *NodeNumber {
	return NewNodeNumber(1)
}

func (n *NodeIdentifier) Children() []Node {
	return []Node{}
}

func (n *NodeIdentifier) ToBoolean() *NodeBoolean {
	return NewNodeBoolean(true)
}

func (n *NodeIdentifier) ToArray() *NodeArray {
	array := NewNodeArray()
	array.Add(n)

	return array
}

func (n *NodeIdentifier) Type() string {
	return "identifier"
}
