package parse

type NodeOperator struct {
	operator string
}

func NewNodeOperator(operator string) *NodeOperator {
	return &NodeOperator{
		operator: operator,
	}
}

func (n *NodeOperator) Operator() string {
	return n.operator
}

func (n *NodeOperator) ToString() *NodeString {
	return NewNodeString(n.operator)
}

func (n *NodeOperator) Children() []Node {
	return []Node{}
}

func (n *NodeOperator) Type() string {
	return "operator"
}
