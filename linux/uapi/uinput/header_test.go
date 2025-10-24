//go:build linux

package uinput_test

import (
	"testing"

	"github.com/andrieee44/gopkg/linux/uapi/internal/uapicgo"
)

func TestConstants(t *testing.T) {
	t.Parallel()
	uapicgo.TestConstants(t, uapicgo.UinputConstants())
}

func TestIOBindings(t *testing.T) {
	t.Parallel()
	uapicgo.TestIOBindings(t, uapicgo.UinputIOBindings())
}

func TestStructs(t *testing.T) {
	t.Parallel()
	uapicgo.TestLayouts(t, uapicgo.UinputLayouts())
}
