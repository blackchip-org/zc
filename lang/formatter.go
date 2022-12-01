package lang

// import (
// 	"fmt"
// 	"reflect"
// 	"strings"
// )

// type FormatOptions struct {
// 	Indent int
// }

// type formatter struct {
// 	out    strings.Builder
// 	opts   FormatOptions
// 	prefix string
// }

// func DefaultFormatOptions() FormatOptions {
// 	return FormatOptions{Indent: 4}
// }

// func Format(root NodeAST) string {
// 	return FormatCustom(root, DefaultFormatOptions())
// }

// func FormatCustom(root NodeAST, opts FormatOptions) string {
// 	var f formatter
// 	f.formatNode(root)
// 	return f.out.String()
// }

// func (f *formatter) formatFile(file *FileNode) {
// 	for _, line := range file.Lines {
// 		f.formatNode(line)
// 	}
// }

// func (f *formatter) formatNode(node NodeAST) {
// 	if node == nil {
// 		panic("UNEXPECTED NIL NODE")
// 	}
// 	switch n := node.(type) {
// 	case *FileNode:
// 		f.formatFile(n)
// 	case *ExprNode:
// 		f.formatLine(n)
// 	case *ValueNode:
// 		f.formatValue(n)
// 	default:
// 		panic(fmt.Sprintf("unknown node: [%v] %+v", reflect.TypeOf(node).Name(), node))
// 	}
// }

// func (f *formatter) formatLine(line *ExprNode) {
// 	f.out.WriteString(f.prefix)
// 	for i, node := range line.Nodes {
// 		if i != 0 {
// 			f.out.WriteString(" ")
// 		}
// 		if node == nil {
// 			panic("UNEXPECTED NIL LINE")
// 		}
// 		f.formatNode(node)
// 	}
// 	f.out.WriteRune('\n')
// }

// func (f *formatter) formatValue(node *ValueNode) {
// 	f.out.WriteString(node.Value)
// }
