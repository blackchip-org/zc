package stack

import "testing"

func TestSlicePush(t *testing.T)     { testPush(t, NewSlice[int]()) }
func TestSlicePop(t *testing.T)      { testPop(t, NewSlice[int]()) }
func TestSlicePopEmpty(t *testing.T) { testPopEmpty(t, NewSlice[int]()) }
func TestSliceLen(t *testing.T)      { testLen(t, NewSlice[int]()) }
func TestSliceAt(t *testing.T)       { testAt(t, NewSlice[int]()) }
func TestSliceDup(t *testing.T)      { testDup(t, NewSlice[int]()) }
func TestSliceTop(t *testing.T)      { testTop(t, NewSlice[int]()) }
func TestSliceUp(t *testing.T)       { testUp(t, NewSlice[int]()) }
