package parse

type NodeKeyword struct {
	keyword string
}

func NewNodeKeyword(keyword string) *NodeKeyword {
	return &NodeKeyword{
		keyword: keyword,
	}
}

func (n *NodeKeyword) Keyword() string {
	return n.keyword
}

func (n *NodeKeyword) ToString() *NodeString {
	return NewNodeString(n.keyword)
}

func (n *NodeKeyword) Children() []Node {
	return []Node{}
}

func (n *NodeKeyword) Type() string {
	return "keyword"
}
