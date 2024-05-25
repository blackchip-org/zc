package zc

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6/types"
)

func TestType(t *testing.T) {
	b := NewCatalogBuilder()
	b.AddType(types.Int)
	cat := b.Build()

	ty, ok := cat.TypeOf(big.NewInt(12))
	if !ok {
		t.Fatalf("expected ok")
	}
	name := "Int"
	if ty.Name() != name {
		t.Errorf("\n have: %v \n want: %v", ty.Name(), name)
	}
}

func TestUnregisteredType(t *testing.T) {
	cat := NewCatalogBuilder().Build()
	_, ok := cat.TypeOf(big.NewInt(12))
	if ok {
		t.Fatalf("expected not ok")
	}
}

func TestDuplicateType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			haveErr := fmt.Sprintf("%v", r)
			wantErr := "duplicate type: Int"
			if haveErr != wantErr {
				t.Fatalf("\n have: %v \n want: %v", haveErr, wantErr)
			}
			return
		}
		t.Fatal("expected panic")
	}()

	b := NewCatalogBuilder()
	b.AddType(types.Int)
	b.AddType(types.Int)
}
