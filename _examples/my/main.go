package main

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
	"os"
)

// A custom lexer for INI files. This illustrates a relatively complex Regexp lexer, as well
// as use of the Unquote filter, which unquotes string tokens.
var (
	iniLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"whitespace", `[ \t]+`},
		{"Other", `[^(\n@)]+`},
		{"FuncName", `^@[a-zA-Z][a-zA-Z_\d]*`},
		{"String", `"(\\.|[^"])*"|'(\\.|[^'])*'|.*`},
		{"Punct", `[)(,]`},
	})
	parser = participle.MustBuild[Entry](
		participle.Lexer(iniLexer),
	)
)

type Entry struct {
	Funcs []*Func `@@*`
}

type Func struct {
	Others *string `@Other`
	Func   *F      `| @@`
}

type F struct {
	FuncName string   `@FuncName`
	Args     []string `"(" @String ("," @String)* ")"`
}

type Map struct {
	Key   string `@Ident ":"`
	Value string `@String`
}

func main() {
	//	lex := lexer.LexString("", `
	//asd(fas)dfsa
	//fsadfa, sdfsadf
	//fasdf"safd
	//@enum("hello:world","foo:bar","say:if", "hello:fsafsadf",
	//"num:1")
	//asdfafasdfasdfasd
	// @copy("fsdafasf:fdsafas", "fsdfa:fasdfa")
	//dsafs"dafasdfasd
	//asdfsa"dfsadf
	//	`)
	//	for {
	//		token, err := lex.Next()
	//		if err != nil {
	//			panic(err)
	//		}
	//		repr.Println(token)
	//		if token.EOF() {
	//			break
	//		}
	//	}
	//
	//	return
	doc, err := parser.ParseString("", `
asd(fas)dfsa
fsadfa, sdfsadf
fasdf"safd
@enum("hello:world","foo:bar","say:if", "hello:fsafsadf",
"num:1")
asdfafasdfasdfasd
 @copy("fsdafasf:fdsafas", "fsdfa:fasdfa")
dsafs"dafasdfasd
asdfsa"dfsadf
	`, participle.Trace(os.Stdout))
	repr.Println(doc, repr.Indent(" "))
	if err != nil {
		panic(err)
	}
}
