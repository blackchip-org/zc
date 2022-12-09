package internal

import "embed"

//go:embed zlib/*.zc test/*.zc
var Scripts embed.FS
