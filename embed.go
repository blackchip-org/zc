package zc

import "embed"

//go:embed zlib modes test/*.zc test/lang/*.zc locales doc/zlib README.md
var Files embed.FS
