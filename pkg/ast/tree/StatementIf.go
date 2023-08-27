package tree

type StatementIf struct {
	condition Expression
	body      []Statement
	elseBody  []Statement
}
