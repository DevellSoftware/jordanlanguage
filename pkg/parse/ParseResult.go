package parse

import "fmt"

type ParseResult struct {
	errors  []*ParseError
	success bool
}

func NewParseResult() *ParseResult {
	return &ParseResult{
		errors:  make([]*ParseError, 0),
		success: true,
	}
}

func (p *ParseResult) setSuccess(success bool) {
	p.success = success
}

func (p *ParseResult) Errors() []*ParseError {
	return p.errors
}

func (p *ParseResult) Success() bool {
	return p.success
}

func (p *ParseResult) AddError(err *ParseError) {
	p.success = false
	p.errors = append(p.errors, err)
}

func (p *ParseResult) PrintErrors() {
	fmt.Println("Errors count: ", len(p.errors))
	for _, err := range p.errors {
		fmt.Println(err.Message())
	}
}
