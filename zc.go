package zc

const ProgName = "zc"

type Calc interface {
	Eval(string) error
	Stack() []string
	SetStack([]string)
	Peek(int) (string, bool)
	Pop() (string, bool)
	MustPop() string
	Push(string)
	Info() string
	SetInfo(string)
	Error() error
	SetError(error)
}

type CalcFunc func(Calc)

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
	return func(c Calc) {
		for _, decl := range op.Funcs {
			if isOpMatch(c, decl) {
				decl.Func(c)
				return
			}
		}
		c.SetError(ErrInvalidArgumentTypes(op.Name))
	}
}

func isOpMatch(c Calc, decl FuncDecl) bool {
	for i, param := range decl.Params {
		arg, ok := c.Peek(len(decl.Params) - i - 1)
		if !ok {
			return false
		}
		if !param.Is(arg) {
			return false
		}
	}
	return true
}

func NoOp(c Calc) {}
