//go:build linux

package uapicgo

// #include <stdlib.h>
import "C"

import (
	"reflect"
	"slices"
	"testing"
	"unsafe"
)

type Constant[T comparable] struct {
	Name  string
	Go, C T
}

type IOBinding struct {
	Name string
	Go   func() (uint32, error)
	C    uint32
}

type StructLayouts struct {
	Name  string
	Go, C []uintptr
}

func TestConstants[T comparable](t *testing.T, constants []Constant[T]) {
	t.Helper()

	var constant Constant[T]

	for _, constant = range constants {
		if constant.Go == constant.C {
			continue
		}

		t.Errorf("%s: got: %v, exp: %v", constant.Name, constant.Go, constant.C)
	}
}

func TestIOBindings(t *testing.T, ioBindings []IOBinding) {
	t.Helper()

	var (
		ioBinding IOBinding
		req       uint32
		err       error
	)

	for _, ioBinding = range ioBindings {
		req, err = ioBinding.Go()
		if err != nil {
			t.Error(err)
		}

		if req == ioBinding.C {
			continue
		}

		t.Errorf("%s: got: %v, exp: %v", ioBinding.Name, req, ioBinding.C)
	}
}

func TestLayouts(t *testing.T, layouts []StructLayouts) {
	t.Helper()

	var layout StructLayouts

	for _, layout = range layouts {
		if slices.Equal(layout.Go, layout.C) {
			continue
		}

		t.Errorf("%s: got: %v, exp: %v", layout.Name, layout.Go, layout.C)
	}
}

func cLayout(fn func(*C.size_t) *C.size_t) []uintptr {
	var (
		count    C.size_t
		layout   *C.size_t
		slice    []C.size_t
		goLayout []uintptr
		idx      int
	)

	layout = fn(&count)
	if layout == nil {
		panic("malloc *size_t for layout failed")
	}

	slice = unsafe.Slice(layout, count)
	goLayout = make([]uintptr, count)

	for idx = range goLayout {
		goLayout[idx] = uintptr(slice[idx])
	}

	C.free(unsafe.Pointer(layout))

	return goLayout
}

func goLayout[T any](a T) []uintptr {
	var (
		typ         reflect.Type
		layout      []uintptr
		fields, idx int
	)

	typ = reflect.TypeOf(a)
	fields = typ.NumField()
	layout = make([]uintptr, fields+1)

	for idx = range fields {
		layout[idx] = typ.Field(idx).Offset
	}

	layout[idx+1] = typ.Size()

	return layout
}
