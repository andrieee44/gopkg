package uapicgo

// #include <linux/uinput.h>
// #include "uinput.h"
import "C"
import "github.com/andrieee44/gopkg/linux/uapi/uinput"

func UinputConstants() []Constant[uint32] {
	return []Constant[uint32]{
		{"uinput.UINPUT_VERSION", uinput.UINPUT_VERSION, C.UINPUT_VERSION},
		{"uinput.UINPUT_MAX_NAME_SIZE", uinput.UINPUT_MAX_NAME_SIZE, C.UINPUT_MAX_NAME_SIZE},
		{"uinput.UINPUT_IOCTL_BASE", uinput.UINPUT_IOCTL_BASE, C.UINPUT_IOCTL_BASE},
		{"uinput.EV_UINPUT", uinput.EV_UINPUT, C.EV_UINPUT},
		{"uinput.UI_FF_UPLOAD", uinput.UI_FF_UPLOAD, C.UI_FF_UPLOAD},
		{"uinput.UI_FF_ERASE", uinput.UI_FF_ERASE, C.UI_FF_ERASE},
	}
}

func UinputIOBindings() []IOBinding {
	return []IOBinding{
		{"uinput.UI_DEV_CREATE()", uinput.UI_DEV_CREATE, C.UI_DEV_CREATE},
		{"uinput.UI_DEV_DESTROY()", uinput.UI_DEV_DESTROY, C.UI_DEV_DESTROY},
		{"uinput.UI_DEV_SETUP()", uinput.UI_DEV_SETUP, C.UI_DEV_SETUP},
		{"uinput.UI_ABS_SETUP()", uinput.UI_ABS_SETUP, C.UI_ABS_SETUP},
		{"uinput.UI_SET_EVBIT()", uinput.UI_SET_EVBIT, C.UI_SET_EVBIT},
		{"uinput.UI_SET_KEYBIT()", uinput.UI_SET_KEYBIT, C.UI_SET_KEYBIT},
		{"uinput.UI_SET_RELBIT()", uinput.UI_SET_RELBIT, C.UI_SET_RELBIT},
		{"uinput.UI_SET_ABSBIT()", uinput.UI_SET_ABSBIT, C.UI_SET_ABSBIT},
		{"uinput.UI_SET_MSCBIT()", uinput.UI_SET_MSCBIT, C.UI_SET_MSCBIT},
		{"uinput.UI_SET_LEDBIT()", uinput.UI_SET_LEDBIT, C.UI_SET_LEDBIT},
		{"uinput.UI_SET_SNDBIT()", uinput.UI_SET_SNDBIT, C.UI_SET_SNDBIT},
		{"uinput.UI_SET_FFBIT()", uinput.UI_SET_FFBIT, C.UI_SET_FFBIT},
		{"uinput.UI_SET_PHYS()", uinput.UI_SET_PHYS, C.UI_SET_PHYS},
		{"uinput.UI_SET_SWBIT()", uinput.UI_SET_SWBIT, C.UI_SET_SWBIT},
		{"uinput.UI_SET_PROPBIT()", uinput.UI_SET_PROPBIT, C.UI_SET_PROPBIT},
		{"uinput.UI_BEGIN_FF_UPLOAD()", uinput.UI_BEGIN_FF_UPLOAD, C.UI_BEGIN_FF_UPLOAD},
		{"uinput.UI_END_FF_UPLOAD()", uinput.UI_END_FF_UPLOAD, C.UI_END_FF_UPLOAD},
		{"uinput.UI_BEGIN_FF_ERASE()", uinput.UI_BEGIN_FF_ERASE, C.UI_BEGIN_FF_ERASE},
		{"uinput.UI_END_FF_ERASE()", uinput.UI_END_FF_ERASE, C.UI_END_FF_ERASE},
		{"uinput.UI_GET_VERSION()", uinput.UI_GET_VERSION, C.UI_GET_VERSION},
		{"uinput.UI_GET_SYSNAME()", func() (uint32, error) {
			return uinput.UI_GET_SYSNAME(uinput.UINPUT_MAX_NAME_SIZE)
		}, uint32(C.wrapUI_GET_SYSNAME(C.UINPUT_MAX_NAME_SIZE))},
	}
}

func UinputLayouts() []StructLayouts {
	return []StructLayouts{
		{
			"uinput.FFUpload",
			goLayout(uinput.FFUpload{}),
			cLayout(func(n *C.size_t) *C.size_t {
				return C.layout_uinput_ff_upload(n)
			}),
		},
		{
			"uinput.FFErase",
			goLayout(uinput.FFErase{}),
			cLayout(func(n *C.size_t) *C.size_t {
				return C.layout_uinput_ff_erase(n)
			}),
		},
		{
			"uinput.Setup",
			goLayout(uinput.Setup{}),
			cLayout(func(n *C.size_t) *C.size_t {
				return C.layout_uinput_setup(n)
			}),
		},
		{
			"uinput.AbsSetup",
			goLayout(uinput.AbsSetup{}),
			cLayout(func(n *C.size_t) *C.size_t {
				return C.layout_uinput_abs_setup(n)
			}),
		},
		{
			"uinput.UserDev",
			//nolint:staticcheck // reason: required for ABI compatibility
			goLayout(uinput.UserDev{}),
			cLayout(func(n *C.size_t) *C.size_t {
				return C.layout_uinput_user_dev(n)
			}),
		},
	}
}
