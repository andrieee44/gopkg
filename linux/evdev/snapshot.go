package evdev

import (
	"fmt"

	"github.com/andrieee44/gopkg/linux/uapi/input"
)

type Snapshot struct {
	ID             input.ID
	Repeat         map[string]uint32
	Enabled        map[string]map[string]bool
	Supported      map[string][]string
	Absolute       map[string]input.AbsInfo
	MultiTouch     map[string][]input.AbsInfo
	Name, Filename string
	Version        int32
	dev            *Device
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

	snap.Repeat = make(map[string]uint32, 2)
	snap.Repeat[codeStr(input.REP_DELAY)] = settings[0]
	snap.Repeat[codeStr(input.REP_PERIOD)] = settings[1]

	return nil
}

func (snap *Snapshot) allocAbsolute(codes []input.AbsoluteCode) {
	var (
		mTcount, absCount int
		code              input.AbsoluteCode
	)

	for _, code = range codes {
		if input.IsMultiTouch(code) {
			mTcount++
		} else {
			absCount++
		}
	}

	snap.Absolute = make(map[string]input.AbsInfo, absCount)
	snap.MultiTouch = make(map[string][]input.AbsInfo, mTcount)
}

func (snap *Snapshot) multiTouch(
	code input.AbsoluteCode,
	absInfo input.AbsInfo,
) error {
	var (
		codeInfo string
		values   []int32
		idx      int
		err      error
	)

	codeInfo = codeStr(code)

	values, err = snap.dev.MTSlotValues(code)
	if err != nil {
		return err
	}

	snap.MultiTouch[codeInfo] = make([]input.AbsInfo, len(values))

	for idx = range snap.MultiTouch[codeInfo] {
		snap.MultiTouch[codeInfo][idx] = absInfo
		snap.MultiTouch[codeInfo][idx].Value = values[idx]
	}

	return nil
}

func (snap *Snapshot) absolute() error {
	var (
		codes   []input.AbsoluteCode
		code    input.AbsoluteCode
		absInfo input.AbsInfo
		err     error
	)

	codes, err = snap.dev.AbsoluteCodes()
	if err != nil {
		return err
	}

	snap.allocAbsolute(codes)

	for _, code = range codes {
		absInfo, err = snap.dev.AbsInfo(code)
		if err != nil {
			return err
		}

		if !input.IsMultiTouch(code) {
			snap.Absolute[codeStr(code)] = absInfo

			continue
		}

		err = snap.multiTouch(code, absInfo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (snap *Snapshot) enabled(event input.EventCode) error {
	var (
		codes, enabledCodes []input.Coder
		codeInfo            string
		code                input.Coder
		err                 error
	)

	codes, err = snap.dev.Codes(event)
	if err != nil {
		return err
	}

	enabledCodes, err = snap.dev.EnabledCodes(event)
	if err != nil {
		return err
	}

	codeInfo = codeStr(event)
	snap.Enabled[codeInfo] = make(map[string]bool, len(codes))

	for _, code = range codes {
		snap.Enabled[codeInfo][codeStr(code)] = false
	}

	for _, code = range enabledCodes {
		snap.Enabled[codeInfo][codeStr(code)] = true
	}

	return nil
}

func (snap *Snapshot) supported(event input.EventCode) error {
	var (
		codes    []input.Coder
		codeInfo string
		idx      int
		err      error
	)

	codes, err = snap.dev.Codes(event)
	if err != nil {
		return err
	}

	codeInfo = codeStr(event)
	snap.Supported[codeInfo] = make([]string, len(codes))

	for idx = range codes {
		snap.Supported[codeInfo][idx] = codeStr(codes[idx])
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
		case input.EV_KEY, input.EV_SW, input.EV_LED, input.EV_SND:
			err = snap.enabled(event)
		case input.EV_SYN,
			input.EV_REL,
			input.EV_MSC,
			input.EV_FF,
			input.EV_PWR,
			input.EV_FF_STATUS:
			err = snap.supported(event)
		default:
			return fmt.Errorf(
				"%s: event %d (%s): %w",
				snap.dev.Filename(),
				event,
				event,
				ErrUnsupportedEvent,
			)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (snap *Snapshot) allocEnabledAndSupported(events []input.EventCode) {
	var (
		event                        input.EventCode
		enabledCount, supportedCount int
	)

	for _, event = range events {
		switch event {
		case input.EV_KEY, input.EV_SW, input.EV_LED, input.EV_SND:
			enabledCount++
		case input.EV_SYN,
			input.EV_REL,
			input.EV_MSC,
			input.EV_FF,
			input.EV_PWR,
			input.EV_FF_STATUS:
			supportedCount++
		}
	}

	snap.Enabled = make(map[string]map[string]bool, enabledCount)
	snap.Supported = make(map[string][]string, supportedCount)
}

func codeStr(coder input.Coder) string {
	return fmt.Sprintf("%s (%d)", coder, coder)
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

	events, err = dev.EventCodes()
	if err != nil {
		return nil, err
	}

	snap.allocEnabledAndSupported(events)

	err = snap.dispatcher(events)
	if err != nil {
		return nil, err
	}

	return snap, nil
}
