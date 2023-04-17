package calc

import (
	"fmt"
	"os"
	"path"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
)

type library struct {
	modules map[string]zc.ModuleDef
	scripts map[string]*ast.File
}

func NewLibrary() zc.Library {
	return &library{
		modules: make(map[string]zc.ModuleDef),
		scripts: make(map[string]*ast.File),
	}
}

func (l *library) Load(name string) ([]byte, error) {
	data, err := internal.Files.ReadFile(path.Join("zlib", name))
	if err == nil {
		return data, nil
	}
	return os.ReadFile(name)
}

func (l *library) Define(m zc.ModuleDef) error {
	_, exists := l.modules[m.Name]
	if exists {
		return fmt.Errorf("module already defined: %v", m.Name)
	}
	script, err := l.Load(m.Script)
	if err != nil {
		return fmt.Errorf("unable to load %v: %v", m.Script, err)
	}
	file, err := parser.Parse(m.Name, script)
	if err != nil {
		return fmt.Errorf("unable to load %v: %v", m.Script, err)
	}
	l.modules[m.Name] = m
	l.scripts[m.Name] = file
	return nil
}

func (l *library) Module(name string) (zc.ModuleDef, bool) {
	mod, ok := l.modules[name]
	return mod, ok
}

func (l *library) Parse(name string) (*ast.File, bool) {
	f, ok := l.scripts[name]
	return f, ok
}
