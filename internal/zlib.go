package internal

import "embed"

//go:embed zlib/* modes/* test/*
var Files embed.FS
