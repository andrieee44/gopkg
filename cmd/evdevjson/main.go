// Package evdevjson reads Linux input events from the evdev interface
// and outputs them as JSON.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/andrieee44/gopkg/linux/evdev"
	"github.com/andrieee44/gopkg/linux/uapi/input"
)

type globalFlags struct {
	pretty bool
}

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)

	os.Exit(1)
}

func exitIf(err error) {
	if err != nil {
		exit(err)
	}
}

func usageFn(fs *flag.FlagSet) func() {
	return func() {
		var err error

		_, err = fmt.Fprintf(fs.Output(), `Usage:
  %s snapshot [<evdev-file>...] [--pretty]
  %s monitor  <evdev-file>  [--pretty]
`, os.Args[0], os.Args[0])
		exitIf(err)

		fs.PrintDefaults()
	}
}

func setupGlobalFlags(fs *flag.FlagSet) *globalFlags {
	var gFlags *globalFlags

	gFlags = new(globalFlags)
	fs.Usage = usageFn(fs)
	fs.BoolVar(&gFlags.pretty, "pretty", false, "Human‑friendly JSON formatting")

	return gFlags
}

func getSnapshots(paths []string) map[string]*evdev.Snapshot {
	var (
		snapshots map[string]*evdev.Snapshot
		snapshot  *evdev.Snapshot
		device    *evdev.Device
		path      string
		err       error
	)

	snapshots = make(map[string]*evdev.Snapshot, len(paths))

	for _, path = range paths {
		device, err = evdev.NewDevice(path)
		exitIf(err)

		snapshot, err = device.Snapshot()
		exitIf(err)

		exitIf(device.Close())

		snapshots[snapshot.Filename] = snapshot
	}

	return snapshots
}

func prettifySnapshots(snapshots map[string]*evdev.Snapshot) map[string]*evdev.SnapshotPretty {
	var (
		snapshotsPretty map[string]*evdev.SnapshotPretty
		key             string
		snapshot        *evdev.Snapshot
	)

	snapshotsPretty = make(map[string]*evdev.SnapshotPretty, len(snapshots))

	for key, snapshot = range snapshots {
		snapshotsPretty[key] = snapshot.Pretty()
	}

	return snapshotsPretty
}

func snapshotCmd(args []string) {
	var (
		fs        *flag.FlagSet
		gFlags    *globalFlags
		snapshots map[string]*evdev.Snapshot
		results   any
		jsonData  []byte
		err       error
	)

	fs = flag.NewFlagSet("snapshot", flag.ExitOnError)
	gFlags = setupGlobalFlags(fs)

	exitIf(fs.Parse(args))
	args = fs.Args()

	if len(args) == 0 {
		args, err = filepath.Glob("/dev/input/event*")
		exitIf(err)
	}

	snapshots = getSnapshots(args)

	results = snapshots
	if gFlags.pretty {
		results = prettifySnapshots(snapshots)
	}

	jsonData, err = json.Marshal(results)
	exitIf(err)

	fmt.Println(string(jsonData))
}

func prettifyEvent(event input.Event) input.EventPretty {
	var (
		eventPretty input.EventPretty
		err         error
	)

	eventPretty, err = input.PrettifyEvent(event)
	exitIf(err)

	return eventPretty
}

func monitorCmd(args []string) {
	var (
		fs        *flag.FlagSet
		gFlags    *globalFlags
		device    *evdev.Device
		eventChan <-chan input.Event
		errChan   <-chan error
		event     input.Event
		results   any
		jsonMsg   []byte
		err       error
	)

	fs = flag.NewFlagSet("snapshot", flag.ExitOnError)
	gFlags = setupGlobalFlags(fs)

	exitIf(fs.Parse(args))
	args = fs.Args()

	if len(args) == 0 {
		exit(errors.New("missing argument"))
	}

	device, err = evdev.NewDevice(args[0])
	exitIf(err)

	defer func() {
		exitIf(device.Close())
	}()

	eventChan, errChan = device.ReadEvents()

	for {
		select {
		case event = <-eventChan:
			results = event
			if gFlags.pretty {
				results = prettifyEvent(event)
			}

			jsonMsg, err = json.Marshal(results)
			exitIf(err)

			fmt.Println(string(jsonMsg))
		case err = <-errChan:
			exit(err)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		usageFn(flag.CommandLine)()

		os.Exit(1)
	}

	switch os.Args[1] {
	case "snapshot":
		snapshotCmd(os.Args[2:])
	case "monitor":
		monitorCmd(os.Args[2:])
	default:
		usageFn(flag.CommandLine)()

		os.Exit(1)
	}
}
