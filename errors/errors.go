package errors

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	StackEmpty = errors.New("stack empty")
)

type InvalidConversion struct {
	Val  any
	From string
	To   string
}

func (e InvalidConversion) Error() string {
	return fmt.Sprintf("cannot convert %v from %v to %v", e.Val, e.From, e.To)
}

func NewInvalidConversion(from string, to string, a any) error {
	return InvalidConversion{
		Val:  a,
		From: from,
		To:   to,
	}
}

type UnexpectedType struct {
	Expected string
	Actual   reflect.Type
	Val      any
}

func (e UnexpectedType) Error() string {
	return fmt.Sprintf("expected type %v but got %v with value %v", e.Expected, e.Actual, e.Val)
}

func NewUnexpectedType(expected string, v any) error {
	return UnexpectedType{
		Expected: expected,
		Actual:   reflect.TypeOf(v),
		Val:      v,
	}
}

type UnexpectedArgKind struct {
	Have  string
	Want  string
	Val   any
	Index int
}

func (e UnexpectedArgKind) Error() string {
	return fmt.Sprintf("expected %v as argument #%v but got %v with value %v", e.Have, e.Index, e.Want, e.Val)
}

func NewUnexpectedArgKind(have, want string, val any, idx int) error {
	return UnexpectedArgKind{
		Have:  have,
		Want:  want,
		Val:   val,
		Index: idx,
	}
}

type UnexpectedRetKind struct {
	Expected string
	Actual   reflect.Type
	Val      any
	Index    int
}

func (e UnexpectedRetKind) Error() string {
	return fmt.Sprintf("expected %v as return #%v but got %v with value %v", e.Expected, e.Index, e.Actual, e.Val)
}

func NewUnexpectedRetKind(expected string, v any, i int) error {
	return UnexpectedRetKind{
		Expected: expected,
		Actual:   reflect.TypeOf(v),
		Val:      v,
		Index:    i,
	}
}

type UnregisteredType struct {
	Type reflect.Type
}

func (e UnregisteredType) Error() string {
	return fmt.Sprintf("unregistered type %v", e.Type)
}

func NewUnregisteredType(v any) error {
	return UnregisteredType{
		Type: reflect.TypeOf(v),
	}
}
