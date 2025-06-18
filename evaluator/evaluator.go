package evaluator

import (
	"github.com/ezParth/interpretor/ast"
	"github.com/ezParth/interpretor/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
