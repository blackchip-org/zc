package unicode_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "unicode",
		ScriptPath: "zc:zlib/char/unicode_/unicode.zc",
		Natives: map[string]zc.CalcFunc{
			"decode":       Decode,
			"encode":       Encode,
			"lower":        Lower,
			"lower=":       LowerIs,
			"title":        Title,
			"title=":       TitleIs,
			"upper":        Upper,
			"upper=":       UpperIs,
			"utf-8-decode": UTF8Decode,
			"utf-8-encode": UTF8Encode,
		},
	}
)
