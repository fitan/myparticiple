package main

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

var (
	iniLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Ident`, `[a-zA-Z][a-zA-Z_\d]*`},
		{`String`, `".*?"|'.*?'|[-+]?[.0-9]+\b`},
		{"Punct", `[)(,:}{\]\[]`},
		{"whitespace", `\s+`},
		{"FuncName", "^@[a-zA-Z]*"},
		{"Op", `=|!=|=~|!~`},
		{"Time", `[0-9]+[smhdw]`},
	})
	parser = participle.MustBuild[Promql](
		participle.Lexer(iniLexer),
	)
)

type Promql struct {
	Name   string   `@Ident?`
	Labels []*Label `"{" ( @@ ("," @@)* )? "}"`
	Time   string   `( "[" @Time "]" )?`
}
type Label struct {
	Key   string `@Ident`
	Op    string `@Op`
	Value string `@String`
}

func main() {
	doc, err := parser.ParseString("", `
name{project="foo",env="prod"}[5s]
	`)
	repr.Println(doc, repr.Indent(" "))
	if err != nil {
		panic(err)
	}
}
