package test

import (
	"path"
	"testing"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/app"
)

func TestModes(t *testing.T) {
	files, err := zc.Files.ReadDir("modes")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		name := path.Base(file.Name())
		name = name[:len(name)-3] // remove .zc suffix
		t.Run(name, func(t *testing.T) {
			c := app.NewDefaultCalc()
			if err := c.SetMode(name); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
