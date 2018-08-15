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
// which contains some statements
type Program struct {
	Statements []Statement
}

// TokenLiteral for Program
// Return: First TokenLiteral() of statements
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns all string of statements in program
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// ExpressionStatement represents Expression in program
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral for expression statement.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String expression statement
func (es ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// ExpressionStatement is Statement
func (es *ExpressionStatement) statementNode() {}

// LetStatement represents let. (binding name)
//  Token  should be "let"
//  Name   should be name of binded value
//  Value  shoule be binded value
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral for let statement.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

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

// LetStatement is Statement
func (ls *LetStatement) statementNode() {}

// ReturnStatement represents return
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

// String stringify Identifier
func (i *Identifier) String() string {
	return i.Value
}

// Identifier is Expression
func (i *Identifier) expressionNode() {}

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

// PrefixExpression represents prefix operator.
//  Token    : Token of prefix operator (ex: !(bang), -(minus))
//  Operator : Operator
//  Right    : Expression to be evaluated with operator
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// PrefixExpression is Expression
func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns string literal of operator token
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String of prefix expression (for debug)
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
