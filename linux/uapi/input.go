package uapi

// InputEvent represents a single input event delivered by the Linux kernel’s
// input subsystem.
type InputEvent struct {
	// Sec is the seconds portion of the event timestamp.
	Sec uint64

	// Usec is the microseconds portion of the event timestamp.
	Usec uint64

	// Type is the high-level category of the event, such as EV_KEY for key
	// or button events, EV_REL for relative motion, or EV_ABS for
	// absolute axes.
	Type uint16

	// Code is the specific identifier within Type, such as a keycode when
	// Type is EV_KEY or an axis code when Type is EV_ABS.
	Code uint16

	// Value holds the data associated with the event.
	// For key events, 0 means release, 1 means press, and 2 means
	// autorepeat. For motion events, it carries the delta or absolute
	// coordinate.
	Value int32
}

// InputID identifies an input device by its bus type, vendor id, product id,
// and version.
type InputID struct {
	// Bustype is the bus type for the device.
	Bustype uint16

	// Vendor is the vendor identifier assigned by the bus.
	Vendor uint16

	// Product is the product identifier assigned by the vendor.
	Product uint16

	// Version is the version or revision number of the device.
	Version uint16
}

// InputAbsInfo holds the parameters of an absolute input axis.
//
// From [input.h]:
//
// struct input_absinfo - used by [EVIOCGABS]/[EVIOCSABS] ioctls
// @value: latest reported value for the axis.
// @minimum: specifies minimum value for the axis.
// @maximum: specifies maximum value for the axis.
// @fuzz: specifies fuzz value that is used to filter noise from the event
// stream.
// @flat: values that are within this value will be discarded by joydev
// interface and reported as 0 instead.
// @resolution: specifies resolution for the values reported for the axis.
//
// Note that input core does not clamp reported values to the
// [minimum, maximum] limits, such task is left to userspace.
//
// The default resolution for main axes ([ABS_X], [ABS_Y], [ABS_Z],
// [ABS_MT_POSITION_X], [ABS_MT_POSITION_Y]) is reported in units
// per millimeter (units/mm), resolution for rotational axes
// ([ABS_RX], [ABS_RY], [ABS_RZ]) is reported in units per radian.
// The resolution for the size axes ([ABS_MT_TOUCH_MAJOR],
// [ABS_MT_TOUCH_MINOR], [ABS_MT_WIDTH_MAJOR], [ABS_MT_WIDTH_MINOR])
// is reported in units per millimeter (units/mm).
// When [INPUT_PROP_ACCELEROMETER] is set the resolution changes.
// The main axes ([ABS_X], [ABS_Y], [ABS_Z]) are then reported in
// units per g (units/g) and in units per degree per second
// (units/deg/s) for rotational axes ([ABS_RX], [ABS_RY], [ABS_RZ]).
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type InputAbsInfo struct {
	// Value is the current position of the axis.
	Value int32

	// Minimum is the lowest value the axis can report.
	Minimum int32

	// Maximum is the highest value the axis can report.
	Maximum int32

	// Fuzz is the noise filter threshold for the axis.
	Fuzz int32

	// Flat is the dead zone around the axis center that is reported as zero.
	Flat int32

	// Resolution is the axis resolution in units per millimeter.
	Resolution int32
}

// InputKeymapEntry maps a hardware scan code to a logical key code.
//
// From [input.h]:
//
// struct input_keymap_entry - used by [EVIOCGKEYCODE]/[EVIOCSKEYCODE] ioctls
// @scancode: scancode represented in machine-endian form.
// @len: length of the scancode that resides in @scancode buffer.
// @index: index in the keymap, may be used instead of scancode
// @flags: allows to specify how kernel should handle the request.
// For example, setting [INPUT_KEYMAP_BY_INDEX] flag indicates that kernel
// should perform lookup in keymap by @index instead of @scancode
// @keycode: key code assigned to this scancode
//
// The structure is used to retrieve and modify keymap data. Users have
// option of performing lookup either by @scancode itself or by @index
// in keymap entry. [EVIOCGKEYCODE] will also return scancode or index
// (depending on which element was used to perform lookup).
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type InputKeymapEntry struct {
	// Flags controls how the kernel handles this request.
	// For example, setting INPUT_KEYMAP_BY_INDEX causes the kernel to
	// look up the mapping by Index instead of by Scancode.
	Flags uint8

	// Len is the length in bytes of the scancode stored in Scancode.
	Len uint8

	// Index is the keymap index used when Flags includes
	// INPUT_KEYMAP_BY_INDEX.
	Index uint16

	// Keycode is the logical key code assigned to this scancode.
	Keycode uint32

	// Scancode holds the hardware scan code in machine-endian form.
	// Only the first Len bytes are significant.
	Scancode [32]uint8
}

// InputMask represents a bitmask of event codes for a given event type.
// It is used with the [EVIOCGBIT] and [EVIOCSBIT] ioctls.
type InputMask struct {
	// Type specifies the event type (for example, EV_KEY or EV_ABS).
	Type uint32

	// CodesSize specifies the length in bytes of the buffer pointed to
	// by CodesPtr.
	CodesSize uint32

	// CodesPtr specifies the user‐space address of the codes bitmask buffer.
	CodesPtr uint32
}

// FFReplay defines the scheduling parameters for a force-feedback effect.
//
// From [input.h]:
//
// struct ff_replay - defines scheduling of the force-feedback effect
// @length: duration of the effect
// @delay: delay before effect should start playing
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFReplay struct {
	// Length is the duration of the effect, in milliseconds.
	Length uint16

	// Delay is the pause before the effect starts playing, in milliseconds.
	Delay uint16
}

// FFTrigger defines what triggers a force-feedback effect.
//
// From [input.h]:
//
// struct ff_trigger - defines what triggers the force-feedback effect
// @button: number of the button triggering the effect
// @interval: controls how soon the effect can be re-triggered
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFTrigger struct {
	// Button is the button number that fires the effect.
	Button uint16

	// Interval is the minimum delay, in milliseconds, before the
	// effect can be triggered again.
	Interval uint16
}

// FFEnvelope describes a generic force-feedback effect envelope.
//
// From [input.h]:
//
// struct ff_envelope - generic force-feedback effect envelope
// @attack_length: duration of the attack (ms)
// @attack_level: level at the beginning of the attack
// @fade_length: duration of fade (ms)
// @fade_level: level at the end of fade
//
// The @attack_level and @fade_level are absolute values; when applying
// envelope force-feedback core will convert to positive/negative
// value based on polarity of the default level of the effect.
// Valid range for the attack and fade levels is 0x0000 - 0x7fff
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFEnvelope struct {
	// AttackLength is the duration of the attack phase, in milliseconds.
	AttackLength uint16

	// AttackLevel is the intensity at the start of the attack phase.
	// Valid range is 0x0000 to 0x7fff.
	AttackLevel uint16

	// FadeLength is the duration of the fade phase, in milliseconds.
	FadeLength uint16

	// FadeLevel is the intensity at the end of the fade phase.
	// Valid range is 0x0000 to 0x7fff.
	FadeLevel uint16
}

// FFConstantEffect defines parameters of a constant force-feedback effect.
//
// From [input.h]:
//
// struct ff_constant_effect - defines parameters of a constant
// force-feedback effect
// @level: strength of the effect; may be negative
// @envelope: envelope data
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFConstantEffect struct {
	// Level is the strength of the effect; may be negative.
	Level int16

	// Envelope holds the force-feedback envelope data.
	Envelope FFEnvelope
}

// FFRampEffect defines parameters of a ramp force-feedback effect.
//
// From [input.h]:
//
// struct ff_ramp_effect - defines parameters of a ramp force-feedback effect
// @start_level: beginning strength of the effect; may be negative
// @end_level: final strength of the effect; may be negative
// @envelope: envelope data
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFRampEffect struct {
	// StartLevel is the beginning strength of the effect; may be negative.
	StartLevel int16

	// EndLevel is the final strength of the effect; may be negative.
	EndLevel int16

	// Envelope holds the envelope parameters defining how the effect’s
	// magnitude changes over its duration.
	Envelope FFEnvelope
}

// FFConditionEffect defines parameters of a spring or friction
// force-feedback effect.
//
// From [input.h]:
//
// struct ff_condition_effect - defines a spring or friction force-feedback
// effect
// @right_saturation: maximum level when joystick moved all way to the right
// @left_saturation: same for the left side
// @right_coeff: controls how fast the force grows when the joystick moves
// to the right
// @left_coeff: same for the left side
// @deadband: size of the dead zone, where no force is produced
// @center: position of the dead zone
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFConditionEffect struct {
	// RightSaturation is the maximum force level when the joystick is moved
	// fully to the right.
	RightSaturation uint16

	// LeftSaturation is the maximum force level when the joystick is moved
	// fully to the left.
	LeftSaturation uint16

	// RightCoeff controls how quickly the force grows as the joystick moves
	// to the right.
	RightCoeff int16

	// LeftCoeff controls how quickly the force grows as the joystick moves
	// to the left.
	LeftCoeff int16

	// Deadband is the size of the zone, in device units, around Center
	// where no force is produced.
	Deadband uint16

	// Center is the position of the dead zone, in device units.
	Center uint16
}

// FFPeriodicEffect defines parameters of a periodic force-feedback effect.
//
// From [input.h]:
//
// struct ff_periodic_effect - defines parameters of a periodic
// force-feedback effect
// @waveform: kind of the effect (wave)
// @period: period of the wave (ms)
// @magnitude: peak value
// @offset: mean value of the wave (roughly)
// @phase: 'horizontal' shift
// @envelope: envelope data
// @custom_len: number of samples ([FF_CUSTOM] only)
// @custom_data: buffer of samples ([FF_CUSTOM] only)
//
// Known waveforms - [FF_SQUARE], [FF_TRIANGLE], [FF_SINE], [FF_SAW_UP],
// [FF_SAW_DOWN], [FF_CUSTOM]. The exact syntax [FF_CUSTOM] is undefined
// for the time being as no driver supports it yet.
//
// Note: the data pointed by custom_data is copied by the driver.
// You can therefore dispose of the memory after the upload/update.
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFPeriodicEffect struct {
	// Waveform is the type of the effect (wave shape).
	// Known values: FF_SQUARE, FF_TRIANGLE, FF_SINE, FF_SAW_UP,
	// FF_SAW_DOWN, FF_CUSTOM.
	Waveform uint16

	// Period is the duration of one cycle of the wave, in milliseconds.
	Period uint16

	// Magnitude is the peak force value; may be negative.
	Magnitude int16

	// Offset is the average force value; may be negative.
	Offset int16

	// Phase is the horizontal shift of the waveform, in [0..Period).
	Phase uint16

	// Envelope holds attack/fade parameters to shape the waveform over time.
	Envelope FFEnvelope

	// CustomLen is the number of samples in CustomData when Waveform is
	// [FF_CUSTOM].
	CustomLen uint16

	// CustomData points to a buffer of raw samples for a custom waveform.
	// The driver copies this data, so it can be released after uploading.
	CustomData *int16
}

// FFRumbleEffect defines the parameters of a dual-motor force-feedback
// rumble effect.
//
// From [input.h]:
//
// struct ff_rumble_effect - defines parameters of a periodic force-feedback
// effect
// @strong_magnitude: magnitude of the heavy motor
// @weak_magnitude: magnitude of the light one
//
// Some rumble pads have two motors of different weight. Strong_magnitude
// represents the magnitude of the vibration generated by the heavy one.
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFRumbleEffect struct {
	// StrongMagnitude is the magnitude of the heavy motor’s vibration.
	StrongMagnitude uint16

	// WeakMagnitude is the magnitude of the light motor’s vibration.
	WeakMagnitude uint16
}

// FFEffect defines parameters of a force-feedback effect for
//
// From [input.h]:
//
// struct ff_effect - defines force feedback effect
// @type: type of the effect ([FF_CONSTANT], [FF_PERIODIC], [FF_RAMP],
// [FF_SPRING], [FF_FRICTION], [FF_DAMPER], [FF_RUMBLE], [FF_INERTIA], or
// [FF_CUSTOM])
// @id: an unique id assigned to an effect
// @direction: direction of the effect
// @trigger: trigger conditions (struct ff_trigger)
// @replay: scheduling of the effect (struct ff_replay)
// @u: effect-specific structure (one of ff_constant_effect, ff_ramp_effect,
// ff_periodic_effect, ff_condition_effect, ff_rumble_effect) further
// defining effect parameters
//
// This structure is sent through ioctl from the application to the driver.
// To create a new effect application should set its @id to -1; the kernel
// will return assigned @id which can later be used to update or delete
// this effect.
//
// Direction of the effect is encoded as follows:
//
//	0 deg -> 0x0000 (down)
//	90 deg -> 0x4000 (left)
//	180 deg -> 0x8000 (up)
//	270 deg -> 0xC000 (right)
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
type FFEffect struct {
	// Type is the effect type.
	Type uint16

	// Id is the effect identifier. Set to -1 when creating a new effect.
	Id int16

	// Direction is the force direction encoded in [0x0000..0xFFFF].
	Direction uint16

	// Trigger defines the trigger conditions for the effect.
	Trigger FFTrigger

	// Replay defines the scheduling parameters for the effect.
	Replay FFReplay

	// Effect holds effect-specific parameters as a raw union payload.
	Effect [32]byte
}

const (
	// EV_VERSION is the version identifier for the Linux input-event
	// interface. It corresponds to the EVIOCGVERSION ioctl request.
	EV_VERSION = 0x010001

	// INPUT_KEYMAP_BY_INDEX is a flag for the EVIOCGKEYCODE_V2 and
	// EVIOCSKEYCODE_V2 ioctls. It tells the kernel to identify the keymap
	// entry by its Index field. When set, the ioctl uses KeymapEntry.Index
	// to select which key mapping to get or set.
	INPUT_KEYMAP_BY_INDEX = 1 << 0

	// ID_BUS is the index for the bus field in device identification.
	ID_BUS = 0

	// ID_VENDOR is the index for the vendor field in device identification.
	ID_VENDOR = 1

	// ID_PRODUCT is the index for the product field in device identification.
	ID_PRODUCT = 2

	// ID_VERSION is the index for the version field in device identification.
	ID_VERSION = 3

	// BUS_PCI represents devices on the PCI bus.
	BUS_PCI InputBusCode = 0x01

	// BUS_ISAPNP represents devices on the ISA Plug-and-Play bus.
	BUS_ISAPNP InputBusCode = 0x02

	// BUS_USB represents devices on the USB bus.
	BUS_USB InputBusCode = 0x03

	// BUS_HIL represents devices on the Hewlett-Packard HIL bus.
	BUS_HIL InputBusCode = 0x04

	// BUS_BLUETOOTH represents devices on the Bluetooth bus.
	BUS_BLUETOOTH InputBusCode = 0x05

	// BUS_VIRTUAL represents a virtual (software) bus.
	BUS_VIRTUAL InputBusCode = 0x06

	// BUS_ISA represents devices on the ISA bus.
	BUS_ISA InputBusCode = 0x10

	// BUS_I8042 represents devices on the i8042 PS/2 controller bus.
	BUS_I8042 InputBusCode = 0x11

	// BUS_XTKBD represents devices on the XTKBD (XTerminal keyboard) bus.
	BUS_XTKBD InputBusCode = 0x12

	// BUS_RS232 represents devices on the RS-232 serial bus.
	BUS_RS232 InputBusCode = 0x13

	// BUS_GAMEPORT represents devices on the legacy gameport bus.
	BUS_GAMEPORT InputBusCode = 0x14

	// BUS_PARPORT represents devices on the parallel port bus.
	BUS_PARPORT InputBusCode = 0x15

	// BUS_AMIGA represents devices on the Amiga proprietary bus.
	BUS_AMIGA InputBusCode = 0x16

	// BUS_ADB represents devices on the Apple Desktop Bus.
	BUS_ADB InputBusCode = 0x17

	// BUS_I2C represents devices on the I2C bus.
	BUS_I2C InputBusCode = 0x18

	// BUS_HOST represents devices local to the host.
	BUS_HOST InputBusCode = 0x19

	// BUS_GSC represents devices on the GSC proprietary bus.
	BUS_GSC InputBusCode = 0x1A

	// BUS_ATARI represents devices on the Atari proprietary bus.
	BUS_ATARI InputBusCode = 0x1B

	// BUS_SPI represents devices on the SPI bus.
	BUS_SPI InputBusCode = 0x1C

	// BUS_RMI represents devices on the RMI (formerly Cirque) bus.
	BUS_RMI InputBusCode = 0x1D

	// BUS_CEC represents devices on the CEC (Consumer Electronics Control)
	// bus.
	BUS_CEC InputBusCode = 0x1E

	// BUS_INTEL_ISHTP represents devices on the Intel ISHTP bus.
	BUS_INTEL_ISHTP InputBusCode = 0x1F

	// BUS_AMD_SFH represents devices on the AMD SFH bus.
	BUS_AMD_SFH InputBusCode = 0x20

	// MT_TOOL_FINGER identifies a finger in multitouch protocols.
	MT_TOOL_FINGER InputMultiTouchCode = 0x00

	// MT_TOOL_PEN identifies a stylus (pen) in multitouch protocols.
	MT_TOOL_PEN InputMultiTouchCode = 0x01

	// MT_TOOL_PALM identifies a palm in multitouch protocols.
	MT_TOOL_PALM InputMultiTouchCode = 0x02

	// MT_TOOL_DIAL identifies a dial or rotary controller in multitouch
	// protocols.
	MT_TOOL_DIAL InputMultiTouchCode = 0x0A

	// MT_TOOL_MAX is the maximum valid multitouch tool value.
	MT_TOOL_MAX InputMultiTouchCode = 0x0F

	// FF_STATUS_STOPPED indicates the force-feedback effect is stopped.
	FF_STATUS_STOPPED InputFFStatusCode = 0x00

	// FF_STATUS_PLAYING indicates the force-feedback effect is playing.
	FF_STATUS_PLAYING InputFFStatusCode = 0x01

	// FF_STATUS_MAX is the highest valid status value.
	FF_STATUS_MAX InputFFStatusCode = 0x01

	// FF_RUMBLE identifies a rumble effect type.
	FF_RUMBLE InputFFCode = 0x50

	// FF_PERIODIC identifies a periodic (waveform) effect type.
	FF_PERIODIC InputFFCode = 0x51

	// FF_CONSTANT identifies a constant force effect type.
	FF_CONSTANT InputFFCode = 0x52

	// FF_SPRING identifies a spring (condition) effect type.
	FF_SPRING InputFFCode = 0x53

	// FF_FRICTION identifies a friction (condition) effect type.
	FF_FRICTION InputFFCode = 0x54

	// FF_DAMPER identifies a damper (condition) effect type.
	FF_DAMPER InputFFCode = 0x55

	// FF_INERTIA identifies an inertia (condition) effect type.
	FF_INERTIA InputFFCode = 0x56

	// FF_RAMP identifies a ramp effect type.
	FF_RAMP InputFFCode = 0x57

	// FF_EFFECT_MIN is the lowest defined effect type value.
	FF_EFFECT_MIN InputFFCode = FF_RUMBLE

	// FF_EFFECT_MAX is the highest defined effect type value.
	FF_EFFECT_MAX InputFFCode = FF_RAMP

	// FF_SQUARE identifies a square waveform for periodic effects.
	FF_SQUARE InputFFCode = 0x58

	// FF_TRIANGLE identifies a triangle waveform for periodic effects.
	FF_TRIANGLE InputFFCode = 0x59

	// FF_SINE identifies a sine waveform for periodic effects.
	FF_SINE InputFFCode = 0x5A

	// FF_SAW_UP identifies a saw-up waveform for periodic effects.
	FF_SAW_UP InputFFCode = 0x5B

	// FF_SAW_DOWN identifies a saw-down waveform for periodic effects.
	FF_SAW_DOWN InputFFCode = 0x5C

	// FF_CUSTOM identifies a custom waveform for periodic effects.
	FF_CUSTOM InputFFCode = 0x5D

	// FF_WAVEFORM_MIN is the lowest defined waveform value.
	FF_WAVEFORM_MIN InputFFCode = FF_SQUARE

	// FF_WAVEFORM_MAX is the highest defined waveform value.
	FF_WAVEFORM_MAX InputFFCode = FF_CUSTOM

	// FF_GAIN controls the global gain (strength) of all effects.
	FF_GAIN InputFFCode = 0x60

	// FF_AUTOCENTER controls the auto-centering feature of condition effects.
	FF_AUTOCENTER InputFFCode = 0x61

	// FF_MAX_EFFECTS is the highest effect-property identifier.
	FF_MAX_EFFECTS InputFFCode = FF_GAIN

	// FF_MAX is the highest valid force-feedback constant.
	FF_MAX InputFFCode = 0x7F

	// FF_CNT is the total number of defined force-feedback constants.
	FF_CNT InputFFCode = FF_MAX + 1
)

// EVIOCGVERSION is the ioctl request code to get the evdev
// driver version. It reads an int into the provided variable.
func EVIOCGVERSION() uint {
	return IOR('E', 0x01, int32(0))
}

// EVIOCGID is the ioctl request code to retrieve the device identifier.
// It reads into an ID struct.
func EVIOCGID() uint {
	return IOR('E', 0x02, InputID{})
}

// EVIOCGREP is the ioctl request code to get keyboard auto‐repeat
// settings. It reads a [2]uint32: [0] = delay in ms, [1] = period in ms.
func EVIOCGREP() uint {
	return IOR('E', 0x03, [2]uint32{})
}

// EVIOCSREP is the ioctl request code to set keyboard auto‐repeat
// settings. It writes a [2]uint32: [0] = delay in ms, [1] = period in ms.
func EVIOCSREP() uint {
	return IOW('E', 0x03, [2]uint32{})
}

// EVIOCGKEYCODE is the ioctl request code to get a simple keycode
// mapping. It reads a [2]uint32: [0] = scancode, [1] = keycode.
func EVIOCGKEYCODE() uint {
	return IOR('E', 0x04, [2]uint32{})
}

// EVIOCGKEYCODE_V2 is the ioctl request code to get an extended
// keymap entry. It reads into a KeymapEntry struct.
func EVIOCGKEYCODE_V2() uint {
	return IOR('E', 0x04, InputKeymapEntry{})
}

// EVIOCSKEYCODE is the ioctl request code to set a simple keycode
// mapping. It writes a [2]uint32: [0] = scancode, [1] = keycode.
func EVIOCSKEYCODE() uint {
	return IOW('E', 0x04, [2]uint32{})
}

// EVIOCSKEYCODE_V2 is the ioctl request code to set an extended
// keymap entry. It writes in a KeymapEntry struct.
func EVIOCSKEYCODE_V2() uint {
	return IOW('E', 0x04, InputKeymapEntry{})
}

// EVIOCGNAME returns the ioctl request code to retrieve the device name.
// The length parameter specifies the size of the buffer (in bytes) that
// will hold the returned name string.
func EVIOCGNAME(length uint) uint {
	return IOC(IOC_READ, 'E', 0x06, length)
}

// EVIOCGPHYS returns the ioctl request code to retrieve the device
// physical location path. The length parameter specifies the size of the
// buffer (in bytes) that will hold the returned physical path string.
func EVIOCGPHYS(length uint) uint {
	return IOC(IOC_READ, 'E', 0x07, length)
}

// EVIOCGUNIQ returns the ioctl request code to retrieve the device’s
// unique identifier. The length parameter specifies the size of the
// buffer (in bytes) that will hold the returned unique ID string.
func EVIOCGUNIQ(length uint) uint {
	return IOC(IOC_READ, 'E', 0x08, length)
}

// EVIOCGPROP returns the ioctl request code to retrieve the device’s
// property bitmask. The length parameter specifies the size of the
// buffer (in bytes) that will hold the returned bitmask.
func EVIOCGPROP(length uint) uint {
	return IOC(IOC_READ, 'E', 0x09, length)
}

// EVIOCGMTSLOTS returns the Linux ioctl command number for reading an
// arbitrary length byte buffer of multi-touch slot data from an input
// device.
//
// From [input.h]:
//
// EVIOCGMTSLOTS(len) - get MT slot values
// @len: size of the data buffer in bytes
//
// The ioctl buffer argument should be binary equivalent to
//
//	struct input_mt_request_layout {
//		__u32 code;
//		__s32 values[num_slots];
//	};
//
// where num_slots is the (arbitrary) number of MT slots to extract.
//
// The ioctl size argument (len) is the size of the buffer, which
// should satisfy len = (num_slots + 1) * sizeof(__s32). If len is
// too small to fit all available slots, the first num_slots are
// returned.
//
// Before the call, code is set to the wanted ABS_MT event type. On
// return, values[] is filled with the slot values for the specified
// ABS_MT code.
//
// If the request code is not an ABS_MT value, -EINVAL is returned.
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
func EVIOCGMTSLOTS(length uint) uint {
	return IOC(IOC_READ, 'E', 0x0a, length)
}

// EVIOCGKEY returns the ioctl request code for retrieving the key bitmask.
func EVIOCGKEY(length uint) uint {
	return IOC(IOC_READ, 'E', 0x18, length)
}

// EVIOCGLED returns the ioctl request code for retrieving the LED bitmask.
func EVIOCGLED(length uint) uint {
	return IOC(IOC_READ, 'E', 0x19, length)
}

// EVIOCGSND returns the ioctl request code for retrieving the sound bitmask.
func EVIOCGSND(length uint) uint {
	return IOC(IOC_READ, 'E', 0x1a, length)
}

// EVIOCGSW returns the ioctl request code for retrieving the switch bitmask.
func EVIOCGSW(length uint) uint {
	return IOC(IOC_READ, 'E', 0x1b, length)
}

// EVIOCGBIT returns the ioctl request code for retrieving the bitmask of
// event type ev. The ev parameter selects an event type offset (for example
// [EV_KEY], [EV_REL]). Passing ev == 0 returns a combined bitmask of all
// supported event types. The length parameter specifies, in bytes, the
// size of the buffer that will receive the bitmask.
func EVIOCGBIT(ev, length uint) uint {
	return IOC(IOC_READ, 'E', 0x20+ev, length)
}

// EVIOCGABS returns the ioctl request code for reading absolute-axis info
// into [InputAbsInfo].
func EVIOCGABS(abs uint) uint {
	return IOR('E', 0x40+abs, InputAbsInfo{})
}

// EVIOCSABS returns the ioctl request code for writing absolute-axis info
// from AbsInfo.
func EVIOCSABS(abs uint) uint {
	return IOW('E', 0xc0+abs, InputAbsInfo{})
}

// EVIOCSFF returns the ioctl request code for uploading (or updating)
// a force-feedback effect.
func EVIOCSFF() uint {
	return IOW('E', 0x80, FFEffect{})
}

// EVIOCRMFF returns the ioctl request code for erasing a previously
// uploaded force-feedback effect.
func EVIOCRMFF() uint {
	return IOW('E', 0x81, int32(0))
}

// EVIOCGEFFECTS returns the ioctl request code for querying how many
// force-feedback effects the device supports.
func EVIOCGEFFECTS() uint {
	return IOR('E', 0x84, int32(0))
}

// EVIOCGRAB returns the ioctl request code for grabbing or releasing an
// input device. Passing a non-zero argument locks event delivery to the
// calling process; zero releases it.
func EVIOCGRAB() uint {
	return IOW('E', 0x90, int32(0))
}

// EVIOCREVOKE returns the ioctl request code for revoking a grab on an
// input device.
func EVIOCREVOKE() uint {
	return IOW('E', 0x91, int32(0))
}

// EVIOCGMASK returns the ioctl request code to retrieve the per-client
// event mask for a specified event type.
//
// From [input.h]:
//
// This ioctl allows user to retrieve the current event mask for specific
// event type. The argument must be of type "struct input_mask" and
// specifies the event type to query, the address of the receive buffer
// and the size of the receive buffer.
//
// The event mask is a per-client mask that specifies which events are
// forwarded to the client. Each event code is represented by a single
// bit in the event mask. If the bit is set, the event is passed to the
// client normally. Otherwise, the event is filtered and will never be
// queued on the client's receive buffer.
//
// Event masks do not affect global state of the input device. They only
// affect the file descriptor they are applied to.
//
// The default event mask for a client has all bits set, i.e. all events
// are forwarded to the client. If the kernel is queried for an unknown
// event type or if the receive buffer is larger than the number of
// event codes known to the kernel, the kernel returns all zeroes for
// those codes.
//
// At maximum, codes_size bytes are copied.
//
// This ioctl may fail with ENODEV in case the file is revoked, EFAULT
// if the receive-buffer points to invalid memory, or EINVAL if the
// kernel does not implement the ioctl.
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
func EVIOCGMASK() uint {
	return IOR('E', 0x92, InputMask{})
}

// EVIOCSMASK returns the ioctl request code to set the per-client
// event mask for a specified event type.
//
// From [input.h]:
//
// This ioctl is the counterpart to [EVIOCGMASK]. Instead of receiving
// the current event mask, this changes the client's event mask for a
// specific type. See [EVIOCGMASK] for a description of event-masks and
// the argument-type.
//
// This ioctl provides full forward compatibility. If the passed event
// type is unknown to the kernel, or if the number of event codes
// specified in the mask is bigger than what is known to the kernel, the
// ioctl is still accepted and applied. However, any unknown codes are
// left untouched and stay cleared. That means, the kernel always filters
// unknown codes regardless of what the client requests. If the new mask
// doesn't cover all known event-codes, all remaining codes are
// automatically cleared and thus filtered.
//
// This ioctl may fail with ENODEV in case the file is revoked. EFAULT
// is returned if the receive-buffer points to invalid memory. EINVAL
// is returned if the kernel does not implement the ioctl.
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
func EVIOCSMASK() uint {
	return IOW('E', 0x93, InputMask{})
}

// EVIOCSCLOCKID returns the ioctl request code which sets the clock
// source used to timestamp input events on a Linux event device.
func EVIOCSCLOCKID() uint {
	return IOW('E', 0xa0, int32(0))
}
