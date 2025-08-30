package evdev

import (
	"errors"
	"fmt"

	"github.com/andrieee44/gopkg/linux/uapi/input"
)

// Snapshot captures a point-in-time view of a [Device]'s state and
// capabilities. A Snapshot includes identifiers, repeat settings, enabled
// and supported events, absolute axis metadata, multi-touch information,
// and descriptor fields such as Name, Filename, and Version.
type Snapshot struct {
	// ID is the evdev device’s identifier.
	ID input.ID

	// Repeat contains the current key‐repeat parameters of the evdev device
	// in milliseconds.
	Repeat map[input.RepeatCode]uint32

	// Absolute reports the absolute value for each absolute code.
	Absolute map[input.AbsoluteCode]input.AbsInfo

	// MultiTouch reports the absolute values for each code across
	// multi-touch slots
	MultiTouch map[input.AbsoluteCode][]int32

	// Key reports the current state of each key code.
	Key map[input.KeyCode]bool

	// Switch reports the current state of each switch code.
	Switch map[input.SwitchCode]bool

	// LED reports the current state of each LED code.
	LED map[input.LEDCode]bool

	// Sound reports the current state of each sound code.
	Sound map[input.SoundCode]bool

	// Sync lists the sync event codes reported by the device.
	Sync []input.SyncCode

	// Relative lists the relative axis codes reported by the device.
	Relative []input.RelativeCode

	// Misc lists miscellaneous event codes reported by the device.
	Misc []input.MiscCode

	// ForceFeedback lists the force‑feedback effect codes supported
	// by the device.
	ForceFeedback []input.FFCode

	// Power lists the power‑management key codes supported by the
	// device.
	Power []input.KeyCode

	// ForceFeedbackStatus lists the force‑feedback status codes
	// reported by the device.
	ForceFeedbackStatus []input.FFStatusCode

	// Name is the evdev device’s name.
	Name string

	// Filename is the name of the underlying file.
	Filename string

	// Version is the evdev device's driver version.
	Version int32

	dev *Device
}

// SnapshotPretty is a human‑readable representation of a [Snapshot].
// It uses string keys for codes and names to make the data easier to
// inspect in formats such as JSON or text output.
type SnapshotPretty struct {
	// ID is the evdev device’s identifier.
	ID input.ID

	// Repeat contains the current key‐repeat parameters of the evdev device
	// in milliseconds.
	Repeat map[string]uint32

	// Absolute reports the absolute value for each absolute code.
	Absolute map[string]input.AbsInfo

	// MultiTouch reports the absolute values for each code across
	// multi-touch slots
	MultiTouch map[string][]int32

	// Key reports the current state of each key code.
	Key map[string]bool

	// Switch reports the current state of each switch code.
	Switch map[string]bool

	// LED reports the current state of each LED code.
	LED map[string]bool

	// Sound reports the current state of each sound code.
	Sound map[string]bool

	// Sync lists the sync event codes reported by the device.
	Sync []string

	// Relative lists the relative axis codes reported by the device.
	Relative []string

	// Misc lists miscellaneous event codes reported by the device.
	Misc []string

	// ForceFeedback lists the force‑feedback effect codes supported
	// by the device.
	ForceFeedback []string

	// Power lists the power‑management key codes supported by the
	// device.
	Power []string

	// ForceFeedbackStatus lists the force‑feedback status codes
	// reported by the device.
	ForceFeedbackStatus []string

	// Name is the evdev device’s name.
	Name string

	// Filename is the name of the underlying file.
	Filename string

	// Version is the evdev device's driver version.
	Version int32
}

// Pretty returns a SnapshotPretty containing a human-readable form of snap.
// The returned value has the same data as snap, but with event codes and
// identifiers converted to descriptive strings.
func (snap *Snapshot) Pretty() *SnapshotPretty {
	return &SnapshotPretty{
		ID:                  snap.ID,
		Repeat:              mapPretty(snap.Repeat),
		Absolute:            mapPretty(snap.Absolute),
		MultiTouch:          mapPretty(snap.MultiTouch),
		Key:                 mapPretty(snap.Key),
		Switch:              mapPretty(snap.Switch),
		LED:                 mapPretty(snap.LED),
		Sound:               mapPretty(snap.Sound),
		Sync:                slicePretty(snap.Sync),
		Relative:            slicePretty(snap.Relative),
		Misc:                slicePretty(snap.Misc),
		ForceFeedback:       slicePretty(snap.ForceFeedback),
		Power:               slicePretty(snap.Power),
		ForceFeedbackStatus: slicePretty(snap.ForceFeedbackStatus),
		Name:                snap.Name,
		Filename:            snap.Filename,
		Version:             snap.Version,
	}
}

func (snap *Snapshot) repeat() error {
	var (
		settings [2]uint32
		err      error
	)

	settings, err = snap.dev.Repeat()
	if err != nil {
		return err
	}

	snap.Repeat = map[input.RepeatCode]uint32{
		input.REP_DELAY:  settings[0],
		input.REP_PERIOD: settings[1],
	}

	return nil
}

func (snap *Snapshot) absolute() error {
	var (
		codes   []input.AbsoluteCode
		code    input.AbsoluteCode
		absInfo input.AbsInfo
		values  []int32
		err     error
	)

	codes, err = snap.dev.Absolutes()
	if err != nil {
		return err
	}

	snap.Absolute = make(map[input.AbsoluteCode]input.AbsInfo, len(codes))
	snap.MultiTouch = make(map[input.AbsoluteCode][]int32, mtLen(codes))

	for _, code = range codes {
		absInfo, err = snap.dev.AbsInfo(code)
		if err != nil {
			return err
		}

		snap.Absolute[code] = absInfo

		values, err = snap.dev.MTSlotValues(code)
		if err == nil {
			snap.MultiTouch[code] = values

			continue
		}

		if !errors.Is(err, ErrNotMultiTouch) {
			return err
		}
	}

	return nil
}

func (snap *Snapshot) dispatcher(events []input.EventCode) error {
	var (
		event input.EventCode
		err   error
	)

	for _, event = range events {
		switch event {
		case input.EV_REP:
			err = snap.repeat()
		case input.EV_ABS:
			err = snap.absolute()
		case input.EV_KEY:
			err = enabled(&snap.Key, snap.dev.Keys, snap.dev.EnabledKeycodes)
		case input.EV_SW:
			err = enabled(&snap.Switch, snap.dev.Switches, snap.dev.EnabledSwitches)
		case input.EV_LED:
			err = enabled(&snap.LED, snap.dev.LEDs, snap.dev.EnabledLEDs)
		case input.EV_SND:
			err = enabled(&snap.Sound, snap.dev.Sounds, snap.dev.EnabledSounds)
		case input.EV_SYN:
			err = supported(&snap.Sync, snap.dev.Syncs)
		case input.EV_REL:
			err = supported(&snap.Relative, snap.dev.Relatives)
		case input.EV_MSC:
			err = supported(&snap.Misc, snap.dev.Miscs)
		case input.EV_FF:
			err = supported(&snap.ForceFeedback, snap.dev.ForceFeedbacks)
		case input.EV_PWR:
			err = supported(&snap.Power, snap.dev.Powers)
		case input.EV_FF_STATUS:
			err = supported(&snap.ForceFeedbackStatus, snap.dev.FFStatuses)
		default:
			return fmt.Errorf(
				"%s: event %s: %w",
				snap.dev.Filename(),
				event.Pretty(),
				input.ErrUnsupportedEvent,
			)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func enabled[T input.Code](
	codesMap *map[T]bool,
	codesFn, enabledCodesFn func() ([]T, error),
) error {
	var (
		codes, enabledCodes []T
		code                T
		err                 error
	)

	codes, err = codesFn()
	if err != nil {
		return err
	}

	enabledCodes, err = enabledCodesFn()
	if err != nil {
		return err
	}

	*codesMap = make(map[T]bool, len(codes))

	for _, code = range codes {
		(*codesMap)[code] = false
	}

	for _, code = range enabledCodes {
		(*codesMap)[code] = true
	}

	return nil
}

func supported[T input.Code](codes *[]T, codesFn func() ([]T, error)) error {
	var err error

	*codes, err = codesFn()

	return err
}

func mapPretty[K input.Code, V any](codes map[K]V) map[string]V {
	var (
		pretty map[string]V
		keys   []K
		key    K
		idx    int
		code   input.Coder
	)

	pretty = make(map[string]V, len(codes))
	keys = make([]K, 0, len(codes))

	for key = range codes {
		keys = append(keys, key)
	}

	for idx, code = range input.AsCoders(keys) {
		pretty[code.Pretty()] = codes[keys[idx]]
	}

	return pretty
}

func slicePretty[T input.Code](codes []T) []string {
	var (
		pretty []string
		idx    int
		code   input.Coder
	)

	pretty = make([]string, len(codes))

	for idx, code = range input.AsCoders(codes) {
		pretty[idx] = code.Pretty()
	}

	return pretty
}

func mtLen(codes []input.AbsoluteCode) int {
	var (
		code   input.AbsoluteCode
		length int
	)

	for _, code = range codes {
		if input.IsMultiTouch(code) {
			length++
		}
	}

	return length
}

func newSnapshot(dev *Device) (*Snapshot, error) {
	var (
		snap   *Snapshot
		events []input.EventCode
		err    error
	)

	snap = &Snapshot{
		dev:      dev,
		Filename: dev.Filename(),
	}

	snap.Version, err = dev.Version()
	if err != nil {
		return nil, err
	}

	snap.ID, err = dev.ID()
	if err != nil {
		return nil, err
	}

	snap.Name, err = dev.Name(256)
	if err != nil {
		return nil, err
	}

	events, err = dev.Events()
	if err != nil {
		return nil, err
	}

	err = snap.dispatcher(events)
	if err != nil {
		return nil, err
	}

	return snap, nil
}
