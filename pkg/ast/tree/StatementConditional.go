package tree

type StatementConditional struct {
	condition Expression
	body      []Statement
}
