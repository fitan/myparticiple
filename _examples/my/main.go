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
	//enumLexer = lexer.MustSimple([]lexer.SimpleRule{
	//	{"whitespace", `\s+`},
	//	{"Punct", `[,)(]`},
	//	{"FuncName", `^@[a-zA-Z][a-zA-Z_\d]*`},
	//	//{"FuncTag", `^@[a-zA-Z][a-zA-Z_\d]*\(`},
	//	{"String", `"(\\.|[^"])*"|'(\\.|[^'])*'`},
	//	{"Ident", `[^ \f\n\r\t\v,)]+`},
	//})
	iniLexer = lexer.MustSimple([]lexer.SimpleRule{
		//{"Whitespace", `\s+`},
		//{"Punct", `[,)(]`},
		//{"Other", `[^@]+`},
		//{"FuncName", `@[a-zA-Z][a-zA-Z_\d]*`},
		//{"String", `"(\\.|[^"])*"|'(\\.|[^'])*'`},
		//{"SQ", `(\\.|[^'])*`},
		//{"DQ", `(\\.|[^"])*`},
		//{"Ident", `[^ \f\n\r\t\v]+`},
		{"whitespace", `\s+`},
		{"Punct", `[,()]`},
		{"FuncName", `^@[a-zA-Z][a-zA-Z_\d]*`},
		//{"FuncTag", `^@[a-zA-Z][a-zA-Z_\d]*\(`},
		{"String", `"(\\.|[^"])*"|'(\\.|[^'])*'`},
		//{"DQ", `(\\.|[^"])+`},
		//{"SQ", `'(\\.|[^'])+'`},
		{"Ident", `[^ \f\n\r\t\v,)]+`},
	})

	parser = participle.MustBuild[Entry](
		participle.Lexer(iniLexer),
	)
)

type Entry struct {
	Funcs []*Func `@@*`
}

type Func struct {
	Others *string `@~FuncName`
	Func   *F      `| @@`
}

//type F struct {
//	FuncName string   `@FuncName`
//	Args     []string `( "(" ( "\"" @DQ "\"" | @SQ | @Ident) ("," ( "\"" @DQ "\"" | @SQ | @Ident))* ")" )?`
//	//Args []string `( "(" ( ( "\"" @Ident "\"") | ( "'" @Ident "'" ) | @Ident ) ("," ( ( "\"" @Ident "\"") | ( "'" @Ident "'" ) | @Ident ) )* ")" )?`
//	//Args []string `"(" ( ( "\"" @Ident "\"") |  @Ident ) ("," ( ( "\"" @Ident "\"") |  @Ident ) )* ")"`
//}

//	type Func struct {
//		Others *string `@Other`
//		//Others *string `@~FuncName`
//		Func *F `| @@`
//	}
type F struct {
	FuncName string   `@FuncName`
	Args     []string `( "(" (@String | @Ident) ("," (@String | @Ident))* ")" )?`
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
@say hello work
@enum("hello:world","foo:bar",'say:if',"hello:fsafsadf", boweian,"num:1")
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
