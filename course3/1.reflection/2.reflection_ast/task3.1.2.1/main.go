package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "as.go", src, 0)
	if err != nil {
		panic(err)
	}

	for _, decl := range file.Decls {
		genDecl, _ := decl.(*ast.GenDecl)
		for _, spec := range genDecl.Specs {
			typeSpec, _ := spec.(*ast.TypeSpec)
			if typeSpec.Name.Name == "MyStruct" {
				typeSpec.Name.Name = "User"
			}
		}
	}

	err = printer.Fprint(os.Stdout, fset, file)
	if err != nil {
		panic(err)
	}
}
