package main

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/goast"
	east "gitlab.com/ptflp/goast"
	"go/token"
	"log"
	"os"
)

func main() {
	// читаем файл в []byte
	data, err := os.ReadFile("course3/1.reflection/2.reflection_ast/task3.1.2.2/test/ast/main.go")
	if err != nil {
		fmt.Println(err, "r")
		return
	}
	fset := token.NewFileSet()

	// создаем декоратор decorator.NewDecoratorWithImports
	dec := decorator.NewDecoratorWithImports(fset, "main.go", goast.New())

	// создаем файл *dest.File с помощью decorator.Parse
	dfile, err := dec.Parse(src)
	if err != nil {
		log.Fatalf("err create file *dest.File %v", err)
	}

	// создаем метод TableName, который возвращает строку, используя east.Method
	tableNameU := east.Method{
		Name:         "TableName",
		Receiver:     "u",
		ReceiverType: "User",
		Arguments:    nil,
		Return:       []east.Param{{Type: "string"}},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{

					Results: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: "\"users\"",
							Decs:  dst.BasicLitDecorations{},
						},
					},
					Decs: dst.ReturnStmtDecorations{
						NodeDecs: dst.NodeDecs{
							Before: dst.NewLine,
						},
					},
				},
			},
			RbraceHasNoPos: false,
			Decs:           dst.BlockStmtDecorations{},
		},
	}
	tableNameA := east.Method{
		Name:         "TableName",
		Receiver:     "a",
		ReceiverType: "Address",
		Arguments:    nil,
		Return:       []east.Param{{Type: "string"}},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: "\"address\"",
							Decs: dst.BasicLitDecorations{NodeDecs: dst.NodeDecs{
								Before: dst.NewLine,
							}},
						},
					},
					Decs: dst.ReturnStmtDecorations{},
				},
			},
			RbraceHasNoPos: false,
			Decs:           dst.BlockStmtDecorations{},
		},
	}

	// добавляем метод в структуру User и Address с помощью east.AddMethod
	_, err = east.AddMethod(dfile, "User", tableNameU)
	_, err = east.AddMethod(dfile, "Address", tableNameA)
	// получаем структуры из файла с помощью east.GetStructs
	gS := east.GetStructs(dfile)
	fmt.Println(data)
	// добавляем теги в структуры с помощью east.ModifyStructs, east.AddDBTags, east.AddDBTypeTags
	decorator.Print(dfile)

	east.ModifyStructs(gS).Error()

	// синхронизируем код с измененными структурами с помощью east.SyncStructs
	east.SyncStructs(dfile, gS).Error()
	// сохранить результат в файл с помощью

	east.WriteASTToFile(dfile, "").Error()
	g, err := east.PrintAST(dfile)
	decorator.Print(dfile)
	fmt.Println(g)
}
