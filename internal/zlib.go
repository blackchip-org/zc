package internal

import "embed"

//go:embed zlib/* test/*
var Files embed.FS
