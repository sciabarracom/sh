package main

import (
	"strings"
	"syscall/js"

	"mvdan.cc/sh/v3/syntax"
)

// figParse is of type:
//
// func figParse(src string) struct{tokens []string; completionType string}
func figParse(this js.Value, args []js.Value) interface{} {
	if this.Type() != js.TypeUndefined {
		panic("figParse: unexpected 'this' value")
	}

	if len(args) != 1 {
		panic("figParse: requires one argument")
	}
	if args[0].Type() != js.TypeString {
		panic("figParse: argument 0 must be string")
	}

	src := args[0].String()
	p := syntax.NewParser(syntax.RecoverErrors(5))
	f, err := p.Parse(strings.NewReader(src), "")
	if err != nil {
		panic(err)
	}

	last := lastStmt(f)
	compl := extractCompletion(last, src)
	return js.ValueOf(compl)
}

func lastStmt(node syntax.Node) *syntax.Stmt {
	var last *syntax.Stmt
	// TODO: keep the first incomplete statement, too
	syntax.Walk(node, func(node syntax.Node) bool {
		if stmt, _ := node.(*syntax.Stmt); stmt != nil {
			last = stmt
		}
		return true
	})
	return last
}

func extractCompletion(stmt *syntax.Stmt, src string) map[string]interface{} {
	// syntax.DebugPrint(os.Stderr, stmt)
	if len(stmt.Redirs) > 0 {
		lastRedir := stmt.Redirs[len(stmt.Redirs)-1]
		// fmt.Println(lastRedir.Word.Pos())
		// fmt.Fprintln(os.Stderr, lastRedir.Word.Pos())
		if lastRedir.Word.Pos().IsRecovered() {
			return compResult(nil, "fileOrFolder")
		}
	}
	if call, _ := stmt.Cmd.(*syntax.CallExpr); call != nil {
		flat := flatWords(call.Args, src)
		if stmt.End().Offset() < uint(len(src)) {
			// Starting a new word.
			flat = append(flat, "")
		}
		return compResult(flat, "normal")
	}

	// other commands, e.g. "if foo; then bar; fi"
	return nil
}

func flatWords(words []*syntax.Word, src string) []interface{} {
	var flat []interface{}
	for _, word := range words {
		start := word.Pos().Offset()
		end := word.End().Offset()
		flat = append(flat, src[start:end])
	}
	return flat
}

func compResult(tokens []interface{}, compType string) map[string]interface{} {
	return map[string]interface{}{
		"tokens":         tokens,
		"completionType": compType,
	}
}

func main() {
	js.Global().Set("figParse", js.FuncOf(figParse))

	// Keep Go running, otherwise the wasm stuff shuts down.
	c := make(chan bool)
	<-c
}
