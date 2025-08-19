// goevtest is a native Go package and CLI tool for inspecting Linux input
// event devices, inspired by the classic evtest C utility.
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/andrieee44/gopkg/linux/evdev"
	"github.com/andrieee44/gopkg/linux/uapi/input"
	"github.com/andrieee44/gopkg/linux/uapi/ioctl"
)

func exit(err error) {
	fmt.Fprintln(os.Stderr, "evtest:", err)
	os.Exit(1)
}

func exitIf(err error) {
	if err != nil {
		exit(err)
	}
}

func printState(
	event uapi.InputEventCode,
	supportedFn, enabledFn func() ([]uapi.InputCoder, error),
) {
	var (
		supportedCodes, enabledCodes []uapi.InputCoder
		supportedCode, enabledCode   uapi.InputCoder
		state                        int
		err                          error
	)

	supportedCodes, err = supportedFn()
	exitIf(err)

	enabledCodes, err = enabledFn()
	exitIf(err)

	fmt.Printf("  Event type %d (%s)\n", event, event)

	for _, supportedCode = range supportedCodes {
		state = 0

		for _, enabledCode = range enabledCodes {
			if supportedCode == enabledCode {
				state = 1
			}
		}

		fmt.Printf("    Event code %d (%s) state %d\n", supportedCode.Value(), supportedCode, state)
	}
}

func printRepeat(dev *evdev.Device) {
	var (
		repeat [2]uint32
		err    error
	)

	repeat, err = dev.Repeat()
	exitIf(err)

	fmt.Printf(`Key repeat handling:
  Repeat type 20 (EV_REP)
    Repeat code 0 (REP_DELAY)
      Value    %d
    Repeat code 1 (REP_PERIOD)
      Value    %d
`, repeat[0], repeat[1])
}

func printDefault(dev *evdev.Device, code uapi.InputCoder) {
	var (
		event uapi.InputEventCode
		codes []uapi.InputCoder
		ok    bool
		err   error
	)

	event, ok = code.(uapi.InputEventCode)
	if !ok {
		exit(fmt.Errorf("evtest: invalid event %d (%s)", code.Value(), code))
	}

	codes, err = dev.Codes(event)
	exitIf(err)

	fmt.Printf("  Event type %d (%s)\n", event, event)

	for _, code = range codes {
		fmt.Printf("    Event code %d (%s)\n", code.Value(), code)
	}
}

func printAbsInfo(absInfo uapi.InputAbsInfo, values []int32) {
	var (
		value int32
		idx   int
	)

	for idx, value = range values {
		fmt.Printf("      Value[%d]    %8d\n", idx, value)
	}

	fmt.Printf(`      Min        %8d
      Max        %8d
      Fuzz       %8d
      Flat       %8d
      Resolution %8d
`, absInfo.Minimum, absInfo.Maximum, absInfo.Fuzz, absInfo.Flat,
		absInfo.Resolution)
}

func printAbs(dev *evdev.Device) {
	var (
		codes   []uapi.InputCoder
		code    uapi.InputCoder
		absInfo uapi.InputAbsInfo
		values  []int32
		err     error
	)

	fmt.Printf("  Event type %d (%s)\n", uapi.EV_ABS, uapi.EV_ABS)

	codes, err = dev.AbsoluteCodes()
	exitIf(err)

	for _, code = range codes {
		absInfo, err = dev.AbsInfo(code)
		exitIf(err)

		fmt.Printf("    Event code %d (%s)\n", code.Value(), code)

		values, err = dev.MTSlots(code)
		if err != nil {
			if !errors.Is(err, evdev.ErrNotMultitouch) {
				exit(err)
			}

			values = []int32{absInfo.Value}
		}

		printAbsInfo(absInfo, values)
	}
}

func printEvents(dev *evdev.Device) {
	var (
		codes []uapi.InputCoder
		code  uapi.InputCoder
		err   error
	)

	codes, err = dev.EventCodes()
	exitIf(err)

	for _, code = range codes {
		switch code.Value() {
		case uapi.EV_REP.Value():
			printRepeat(dev)
		case uapi.EV_KEY.Value():
			printState(uapi.EV_KEY, dev.Keycodes, dev.EnabledKeycodes)
		case uapi.EV_LED.Value():
			printState(uapi.EV_LED, dev.LEDCodes, dev.EnabledLEDs)
		case uapi.EV_SND.Value():
			printState(uapi.EV_LED, dev.SoundCodes, dev.EnabledSounds)
		case uapi.EV_SW.Value():
			printState(uapi.EV_SW, dev.SwitchCodes, dev.EnabledSwitches)
		case uapi.EV_ABS.Value():
			printAbs(dev)
		default:
			printDefault(dev, code)
		}
	}
}

func printDev(dev *evdev.Device) {
	var (
		version    int32
		id         uapi.InputID
		name       string
		properties []uapi.InputCoder
		property   uapi.InputCoder
		err        error
	)

	version, err = dev.Version()
	exitIf(err)

	id, err = dev.ID()
	exitIf(err)

	name, err = dev.Name()
	exitIf(err)

	properties, err = dev.Properties()
	exitIf(err)

	fmt.Printf("Input driver version is %x\n", version)

	fmt.Printf(
		"Input device ID: bus 0x%x vendor 0x%x product 0x%x version 0x%x\n",
		id.Bustype, id.Vendor, id.Product, id.Version,
	)

	fmt.Printf("Input device name: %q\n", name)
	fmt.Println("Supported events:")

	printEvents(dev)

	fmt.Println("Properties:")

	for _, property = range properties {
		fmt.Printf("  Property type %d (%s)\n", property.Value(), property)
	}
}

func main() {
	var (
		dev        *evdev.Device
		paths      []string
		path, name string
		devN       int
		err        error
	)

	if len(os.Args) != 1 {
		dev, err = evdev.NewDevice(os.Args[1])
		exitIf(err)

		printDev(dev)

		err = dev.Close()
		exitIf(err)

		return
	}

	paths, err = filepath.Glob("/dev/input/event*")
	exitIf(err)

	fmt.Println("No device specified, trying to scan all of /dev/input/event*")

	if syscall.Geteuid() != 0 {
		fmt.Println("Not running as root, no devices may be available.")
	}

	fmt.Println("Available devices:")

	for _, path = range paths {
		dev, err = evdev.NewDevice(path)
		exitIf(err)

		name, err = dev.Name()
		exitIf(err)

		err = dev.Close()
		exitIf(err)

		fmt.Printf("%s:	%s\n", path, name)
	}

	fmt.Printf("Select the device event number [0-%d]: ", len(paths)-1)

	_, err = fmt.Scan(&devN)
	exitIf(err)

	dev, err = evdev.NewDevice(fmt.Sprintf("/dev/input/event%d", devN))
	exitIf(err)

	printDev(dev)

	err = dev.Close()
	exitIf(err)
}
