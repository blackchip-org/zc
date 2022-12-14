package zc

// type Library struct {
// 	defs    map[string]ModuleDef
// 	mods    map[string]*Calc
// 	prelude map[string]*Calc
// }

// func (l *Library) Register(def ModuleDef) {
// 	_, ok := l.mods[def.Name]
// 	if ok {
// 		panic(fmt.Sprintf("cannot register module %v, already loaded", def.Name))
// 	}
// 	l.defs[def.Name] = def
// }

// func (l *Library) LoadPrelude(name string) (*Calc, error) {
// 	p, err := l.Load(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	l.prelude[name] = p
// 	return p, nil
// }

// func (l *Library) Load(name string) (*Calc, error) {
// 	mod, ok := l.mods[name]
// 	if ok {
// 		return mod, nil
// 	}

// 	def, ok := l.defs[name]
// 	if !ok {
// 		return nil, errors.ModuleNotFound{Name: name}
// 	}

// 	c := NewCalc()

// 	root := ast.Node(&ast.FileNode{})
// 	if def.ScriptPath != "" {
// 		src, err := LoadFile(def.ScriptPath)
// 		if err != nil {
// 			return nil, err
// 		}
// 		root, err = parser.Parse(def.ScriptPath, src)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	l.mods[name] = root
// 	return root, nil
// }
