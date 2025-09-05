package uinput

import (
	"github.com/andrieee44/gopkg/linux/internal/ioctlwrap"
	"github.com/andrieee44/gopkg/linux/uapi/input"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

// FFUpload represents a force feedback upload request exchanged with the
// kernel via the uinput interface. It is used when adding or updating a
// force feedback effect on a virtual input device. The kernel fills this
// structure when processing an upload request, allowing the caller to
// inspect the new effect state and the previous state it replaced.
type FFUpload struct {
	// RequestID is the unique identifier for this upload transaction.
	// It is assigned by the kernel and used to match the request with
	// its corresponding completion.
	RequestID uint32

	// Retval is the return value set by the driver after processing the
	// upload. A value of 0 indicates success; a negative value indicates
	// an error.
	Retval int32

	// Effect contains the new input.FFEffect to be stored in the
	// device's iowrap buffer.
	Effect input.FFEffect

	// Old contains the previous input.FFEffect that is being replaced
	// by this upload.
	Old input.FFEffect
}

// FFErase represents a force feedback erase request exchanged with the
// kernel via the uinput interface. It is used when removing a force
// feedback effect from a virtual input device's iowrap buffer. The
// kernel fills this structure when processing an erase request, allowing
// the caller to approve or reject the removal.
type FFErase struct {
	// RequestID is the unique identifier for this erase transaction.
	// It is assigned by the kernel and used to match the request with
	// its corresponding completion.
	RequestID uint32

	// Retval is the return value set by the driver after processing the
	// erase request. A value of 0 indicates success; a negative value
	// indicates an error.
	Retval int32

	// EffectID is the identifier of the force feedback effect to be
	// removed from the device's buffer.
	EffectID uint32
}

// Setup defines the basic properties of a virtual input device. It holds
// the device identity, human‑readable name, and optional force feedback
// capacity, which are applied when registering the device.
//
// From [uinput.h]:
//
// To write a force-feedback-capable driver, the upload_effect
// and erase_effect callbacks in [Setup] must be implemented.
// The uinput driver will generate a fake input event when one of
// these callbacks are invoked. The userspace code then uses
// ioctls to retrieve additional parameters and send the return code.
// The callback blocks until this return code is sent.
//
// The described callback mechanism is only used if ff_effects_max
// is set.
//
// To implement upload_effect():
//  1. Wait for an event with Type == [EV_UINPUT] and Code == [UI_FF_UPLOAD].
//     A request ID will be given in Value.
//  2. Allocate a [FFUpload] struct, fill in RequestID with
//     the Value from the [EV_UINPUT] event.
//  3. Issue a [UI_BEGIN_FF_UPLOAD] ioctl, giving it the
//     [FFUpload] struct. It will be filled in with the
//     [input.FFEffect] passed to upload_effect().
//  4. Perform the effect upload, and place a return code back into
//     the [FFUpload] struct.
//  5. Issue a [UI_END_FF_UPLOAD] ioctl, also giving it the
//     [FFUpload] struct. This will complete execution
//     of our upload_effect() handler.
//
// To implement erase_effect():
//  1. Wait for an event with Type == [EV_UINPUT] and Code == [UI_FF_ERASE].
//     A request ID will be given in Value.
//  2. Allocate a [FFErase] struct, fill in RequestID with
//     the Value from the [EV_UINPUT] event.
//  3. Issue a [UI_BEGIN_FF_ERASE] ioctl, giving it the
//     [FFErase] struct. It will be filled in with the
//     effect ID passed to erase_effect().
//  4. Perform the effect erasure, and place a return code back
//     into the [FFErase] struct.
//  5. Issue a [UI_END_FF_ERASE] ioctl, also giving it the
//     [FFErase] struct. This will complete execution
//     of our erase_effect() handler.
//
// [EV_UINPUT] is the new event type, used only by uinput.
// Code is [UI_FF_UPLOAD] or [UI_FF_ERASE], and Value
// is the unique request ID. This number was picked
// arbitrarily, above [EV_MAX] (since the input system
// never sees it) but in the range of a 16-bit int.
//
// [uinput.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
type Setup struct {
	// ID describes the bus type, vendor, product, and version of the
	// device, as defined by input.ID.
	ID input.ID

	// Name is a null‑terminated UTF‑8 string containing the device name
	// presented to userspace. It must not exceed UINPUT_MAX_NAME_SIZE
	// bytes including the terminating null byte.
	Name [UINPUT_MAX_NAME_SIZE]byte

	// FFEffectsMax specifies the maximum number of force feedback effects
	// the device can store in its iowrap buffer. Set to zero if the
	// device does not support force feedback.
	FFEffectsMax uint32
}

// AbsSetup defines the configuration for a single absolute axis on a
// virtual input device. It specifies the axis code and its associated
// parameters.
type AbsSetup struct {
	// Code is the event code of the absolute axis (e.g., ABS_X, ABS_Y).
	Code input.AbsoluteCode

	// AbsInfo describes the properties of the axis, such as its minimum
	// and maximum values, fuzz, flat, and resolution, as defined by
	// input.AbsInfo.
	AbsInfo input.AbsInfo
}

// UserDev defines the properties and capabilities of a virtual input
// device. It includes the device name, identity, optional force feedback
// capacity, and absolute axis parameters.
//
// Deprecated: This structure is considered deprecated in the Linux kernel
// documentation in favor of newer setup ioctls, but may still be used for
// compatibility.
type UserDev struct {
	// Name is a null‑terminated UTF‑8 string containing the device name
	// presented to userspace. It must not exceed UINPUT_MAX_NAME_SIZE
	// bytes including the terminating null byte.
	Name [UINPUT_MAX_NAME_SIZE]byte

	// ID describes the bus type, vendor, product, and version of the
	// device, as defined by input.ID.
	ID input.ID

	// FFEffectsMax specifies the maximum number of force feedback effects
	// the device can store in its iowrap buffer. Set to zero if the
	// device does not support force feedback.
	FFEffectsMax uint32

	// AbsMax holds the maximum values for each absolute axis, indexed by
	// axis code.
	AbsMax [input.ABS_CNT]int32

	// AbsMin holds the minimum values for each absolute axis, indexed by
	// axis code.
	AbsMin [input.ABS_CNT]int32

	// AbsFuzz specifies the tolerance for each absolute axis, used to
	// filter out small, insignificant changes.
	AbsFuzz [input.ABS_CNT]int32

	// AbsFlat specifies the dead zone for each absolute axis, within
	// which movement is ignored.
	AbsFlat [input.ABS_CNT]int32
}

const (
	// UINPUT_VERSION is the version number of the uinput kernel interface.
	UINPUT_VERSION = 5

	// UINPUT_MAX_NAME_SIZE is the maximum length of a device name in bytes.
	UINPUT_MAX_NAME_SIZE = 80

	// UINPUT_IOCTL_BASE is the ioctl base character used for uinput commands.
	UINPUT_IOCTL_BASE = 'U'

	// EV_UINPUT is the event type code for uinput-specific events.
	EV_UINPUT = 0x0101

	// UI_FF_UPLOAD is the ioctl code for uploading a force-feedback effect.
	UI_FF_UPLOAD = 1

	// UI_FF_ERASE is the ioctl code for erasing a force-feedback effect.
	UI_FF_ERASE = 2
)

// UI_DEV_CREATE is the ioctl request code to create a virtual input device.
// It takes no arguments.
func UI_DEV_CREATE() (uint32, error) {
	return ioctlwrap.IO(UINPUT_IOCTL_BASE, 1, "UI_DEV_CREATE request failed")
}

// UI_DEV_DESTROY is the ioctl request code to destroy a virtual input device.
// It takes no arguments.
func UI_DEV_DESTROY() (uint32, error) {
	return ioctlwrap.IO(UINPUT_IOCTL_BASE, 2, "UI_DEV_DESTROY request failed")
}

// UI_DEV_SETUP is the ioctl request code to set up parameters for a virtual
// input device. It writes a Setup struct.
//
// From [uinput.h]:
//
// [UI_DEV_SETUP] - Set device parameters for setup
//
// This ioctl sets parameters for the input device to be created.  It
// supersedes the old "struct [UserDev]" method, which wrote this data
// via write(). To actually set the absolute axes [UI_ABS_SETUP] should be
// used.
//
// The ioctl takes a "struct [Setup]" object as argument. The fields of
// this object are as follows:
//
//	          ID: See the description of "struct [input.ID]". This field is
//	              copied unchanged into the new device.
//	        Name: This is used unchanged as name for the new device.
//	FFEffectsMax: This limits the maximum numbers of force-feedback effects.
//	              See below for a description of FF with uinput.
//
// This ioctl can be called multiple times and will overwrite previous values.
// If this ioctl fails with -EINVAL, it is recommended to use the old
// "[UserDev]" method via write() as a fallback, in case you run on an
// old kernel that does not support this ioctl.
//
// This ioctl may fail with -EINVAL if it is not supported or if you passed
// incorrect values, -ENOMEM if the kernel runs out of memory or -EFAULT if
// the passed [Setup] object cannot be read/written.
// If this call fails, partial data may have already been applied to the
// iowrap device.
//
// [uinput.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
func UI_DEV_SETUP() (uint32, error) {
	return ioctlwrap.IOW[Setup](UINPUT_IOCTL_BASE, 3, "UI_DEV_SETUP request failed")
}

// UI_ABS_SETUP is the ioctl request code to configure an absolute axis on a
// virtual input device. It writes an AbsSetup struct.
//
// From [uinput.h]:
//
// UI_ABS_SETUP - Set absolute axis information for the device to setup
//
// This ioctl sets one absolute axis information for the input device to be
// created. It supersedes the old "struct [UserDev]" method, which wrote
// part of this data and the content of [UI_DEV_SETUP] via write().
//
// The ioctl takes a "struct [AbsSetup]" object as argument. The fields
// of this object are as follows:
//
//	   Code: The corresponding input code associated with this axis
//	         ([input.ABS_X], [input.ABS_Y], etc...)
//	AbsInfo: See "struct [input.AbsInfo]" for a description of this field.
//	         This field is copied unchanged into the kernel for the
//	         specified axis. If the axis is not enabled via
//	         [UI_SET_ABSBIT], this ioctl will enable it.
//
// This ioctl can be called multiple times and will overwrite previous values.
// If this ioctl fails with -EINVAL, it is recommended to use the old
// "[UserDev]" method via write() as a fallback, in case you run on an
// old kernel that does not support this ioctl.
//
// This ioctl may fail with -EINVAL if it is not supported or if you passed
// incorrect values, -ENOMEM if the kernel runs out of memory or -EFAULT if
// the passed [Setup] object cannot be read/written.
// If this call fails, partial data may have already been applied to the
// iowrap device.
//
// [uinput.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
func UI_ABS_SETUP() (uint32, error) {
	return ioctlwrap.IOW[AbsSetup](UINPUT_IOCTL_BASE, 4, "UI_ABS_SETUP request failed")
}

// UI_SET_EVBIT is the ioctl request code to enable an event type. It writes
// an int.
func UI_SET_EVBIT() (uint32, error) {
	return ioctlwrap.IOW[int](UINPUT_IOCTL_BASE, 100, "UI_SET_EVBIT request failed")
}

// UI_SET_KEYBIT is the ioctl request code to enable a key code. It writes an
// int.
func UI_SET_KEYBIT() (uint32, error) {
	return ioctlwrap.IOW[int](UINPUT_IOCTL_BASE, 101, "UI_SET_KEYBIT request failed")
}

// UI_SET_RELBIT is the ioctl request code to enable a relative axis. It
// writes an int32.
func UI_SET_RELBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 102, "UI_SET_RELBIT request failed")
}

// UI_SET_ABSBIT is the ioctl request code to enable an absolute axis. It
// writes an int32.
func UI_SET_ABSBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 103, "UI_SET_ABSBIT request failed")
}

// UI_SET_MSCBIT is the ioctl request code to enable a miscellaneous event.
// It writes an int32.
func UI_SET_MSCBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 104, "UI_SET_MSCBIT request failed")
}

// UI_SET_LEDBIT is the ioctl request code to enable an LED event. It writes
// an int32.
func UI_SET_LEDBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 105, "UI_SET_LEDBIT request failed")
}

// UI_SET_SNDBIT is the ioctl request code to enable a sound event. It writes
// an int32.
func UI_SET_SNDBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 106, "UI_SET_SNDBIT request failed")
}

// UI_SET_FFBIT is the ioctl request code to enable a force-feedback effect.
// It writes an int32.
func UI_SET_FFBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 107, "UI_SET_FFBIT request failed")
}

// UI_SET_PHYS is the ioctl request code to set the physical path of the
// device. It writes a string.
func UI_SET_PHYS() (uint32, error) {
	return ioctlwrap.IOW[*byte](UINPUT_IOCTL_BASE, 108, "UI_SET_PHYS request failed")
}

// UI_SET_SWBIT is the ioctl request code to enable a switch event. It writes
// an int32.
func UI_SET_SWBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 109, "UI_SET_SWBIT request failed")
}

// UI_SET_PROPBIT is the ioctl request code to enable a device property. It
// writes an int32.
func UI_SET_PROPBIT() (uint32, error) {
	return ioctlwrap.IOW[int32](UINPUT_IOCTL_BASE, 110, "UI_SET_PROPBIT request failed")
}

// UI_BEGIN_FF_UPLOAD is the ioctl request code to begin uploading a
// force-feedback effect. It reads/writes an FFUpload struct.
func UI_BEGIN_FF_UPLOAD() (uint32, error) {
	return ioctlwrap.IOWR[FFUpload](UINPUT_IOCTL_BASE, 200,
		"UI_BEGIN_FF_UPLOAD request failed")
}

// UI_END_FF_UPLOAD is the ioctl request code to end uploading a
// force-feedback effect. It writes an FFUpload struct.
func UI_END_FF_UPLOAD() (uint32, error) {
	return ioctlwrap.IOW[FFUpload](UINPUT_IOCTL_BASE, 201,
		"UI_END_FF_UPLOAD request failed")
}

// UI_BEGIN_FF_ERASE is the ioctl request code to begin erasing a
// force-feedback effect. It reads/writes an FFErase struct.
func UI_BEGIN_FF_ERASE() (uint32, error) {
	return ioctlwrap.IOWR[FFErase](UINPUT_IOCTL_BASE, 202,
		"UI_BEGIN_FF_ERASE request failed")
}

// UI_END_FF_ERASE is the ioctl request code to end erasing a force-feedback
// effect. It writes an FFErase struct.
func UI_END_FF_ERASE() (uint32, error) {
	return ioctlwrap.IOW[FFErase](UINPUT_IOCTL_BASE, 203,
		"UI_END_FF_ERASE request failed")
}

// UI_GET_SYSNAME is the ioctl request code to read the sysfs name of a
// created virtual input device. It reads a string of up to length bytes.
//
// From [uinput.h]:
//
// UI_GET_SYSNAME - get the sysfs name of the created uinput device
//
// @return the sysfs name of the created virtual input device.
// The complete sysfs path is then /sys/devices/virtual/input/--NAME--
// Usually, it is in the form "inputN"
//
// [uinput.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
func UI_GET_SYSNAME(length uint32) (uint32, error) {
	return ioctlwrap.IOC(
		ioctl.IOC_READ,
		UINPUT_IOCTL_BASE,
		44,
		length,
		"UI_GET_SYSNAME request failed",
	)
}

// UI_GET_VERSION is the ioctl request code to read the uinput driver
// version. It reads a uint32.
//
// From [uinput.h]:
//
// UI_GET_VERSION - Return version of uinput protocol
//
// This writes uinput protocol version implemented by the kernel into
// the integer pointed to by the ioctl argument. The protocol version
// is hard-coded in the kernel and is independent of the uinput device.
//
// [uinput.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h
func UI_GET_VERSION() (uint32, error) {
	return ioctlwrap.IOR[uint32](UINPUT_IOCTL_BASE, 45, "UI_GET_VERSION request failed")
}
