package updater

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"

	"github.com/infiniteloopcloud/jsongen/pkg/logger"
)

func ParseAndModify(file string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "isEmptyValue" {
			assignStmt := assignStmt()
			ifStmt := ifStmt()

			funcDecl.Body.List = append([]ast.Stmt{assignStmt, ifStmt}, funcDecl.Body.List...)
		}
	}

	return node, fset, nil
}

func PersistChanges(node *ast.File, fset *token.FileSet, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	return printer.Fprint(file, fset, node)
}

func assignStmt() ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{Name: "z"},
			&ast.Ident{Name: "ok"},
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.TypeAssertExpr{
				X: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "v"},
					Sel: &ast.Ident{Name: "Interface()"},
				},
				Type:   &ast.Ident{Name: "IsZeroer"},
				Rparen: 0,
			},
		},
	}
}

func ifStmt() ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.Ident{
			Name: "ok",
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "z",
								},
								Sel: &ast.Ident{
									Name: "IsZero",
								},
							},
						},
					},
				},
			},
		},
	}
}
