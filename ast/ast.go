package ast

import (
	"bytes"

	"github.com/nicetacker/monkey/token"
)

// Interfaces

// Node is node
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is statement
type Statement interface {
	Node
	statementNode()
}

// Expression is expression
type Expression interface {
	Node
	expressionNode()
}

// AST nodes

// Program represents program
type Program struct {
	Statements []Statement
}

// TokenLiteral for Program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// ExpressionStatement represents expression in program
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral for expression statement.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String expression statement
func (es ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// LetStatement is let
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral for let statement.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) statementNode() {}

// String let statement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ReturnStatement is return
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// String let statement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// TokenLiteral for return statement.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) statementNode() {}

// Identifier is ident
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral for identifier statement.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) expressionNode() {}

// String stringify Identifier
func (i *Identifier) String() string {
	return i.Value
}

// IntegralLiteral represents int literal
type IntegralLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegralLiteral) expressionNode() {}

// TokenLiteral for integral literals.
func (i *IntegralLiteral) TokenLiteral() string { return i.Token.Literal }

// String stringfy IntegralLiteral
func (i *IntegralLiteral) String() string {
	return i.Token.Literal
}
