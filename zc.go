package zc

type Calc interface {
	Eval(string) error
	Stack() []string
}

type Env interface {
	Peek(int) (string, bool)
	Pop() (string, bool)
	MustPop() string
	Push(string)
	Error(error)
	//Eval(string)
}

type CalcFunc func(Env)

type OpDecl struct {
	Name  string
	Funcs []FuncDecl
}

type FuncDecl struct {
	Func   CalcFunc
	Params []Type
}

func Op(name string, fn CalcFunc, param ...Type) OpDecl {
	return OpDecl{
		Name: name,
		Funcs: []FuncDecl{
			{Func: fn, Params: param},
		},
	}
}

func GenOp(name string, funcs ...FuncDecl) OpDecl {
	return OpDecl{Name: name, Funcs: funcs}
}

func Func(fn CalcFunc, params ...Type) FuncDecl {
	return FuncDecl{Func: fn, Params: params}
}

func EvalOp(op OpDecl) CalcFunc {
	return func(e Env) {
		for _, decl := range op.Funcs {
			if isOpMatch(e, decl) {
				decl.Func(e)
				return
			}
		}
		e.Error(ErrInvalidArgumentTypes(op.Name))
	}
}

func isOpMatch(e Env, decl FuncDecl) bool {
	for i, param := range decl.Params {
		arg, ok := e.Peek(len(decl.Params) - i)
		if !ok {
			return false
		}
		if !param.Is(arg) {
			return false
		}
	}
	return true
}

func NoOp(e Env) {}
