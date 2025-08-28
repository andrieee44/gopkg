package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andrieee44/gopkg/linux/evdev"
)

func main() {
	var (
		devices   []*evdev.Device
		device    *evdev.Device
		snapshots map[string]*evdev.Snapshot
		snapshot  *evdev.Snapshot
		jsonMsg   []byte
		errs      []error
		err       error
	)

	devices, errs = evdev.Devices()
	if len(errs) != 0 {
		for _, err = range errs {
			fmt.Fprintf(os.Stderr, "evdevjson: %s\n", err)

			os.Exit(1)
		}
	}

	snapshots = make(map[string]*evdev.Snapshot, len(devices))

	for _, device = range devices {
		snapshot, err = device.Snapshot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "evdevjson: %s\n", err)

			os.Exit(1)
		}

		snapshots[snapshot.Filename] = snapshot
	}

	jsonMsg, err = json.Marshal(snapshots)
	if err != nil {
		fmt.Fprintf(os.Stderr, "evdevjson: %s\n", err)

		os.Exit(1)
	}

	fmt.Println(string(jsonMsg))
}
