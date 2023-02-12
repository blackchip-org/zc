package internal

import "embed"

//go:embed zlib/* modes/* test/* locales/*
var Files embed.FS
