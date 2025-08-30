% evdevjson(1) | User Commands

# NAME
**evdevjson** - snapshot or monitor Linux evdev device state

# SYNOPSIS
`evdevjson` monitor [--pretty] <path>
`evdevjson` snapshot [--pretty] [<path>...]

# DESCRIPTION
**evdevjson** is a command-line utility for inspecting the state of Linux
input devices via the **evdev** interface. It can operate in two modes:

- **monitor** - Continuously read and display events from the specified
  evdev device.
- **snapshot** - Output the current state of one or more evdev devices.
  If no `<path>` arguments are given, the state of all available evdev
  devices will be output.

Unlike `evtest`, which prints human-readable event logs, **evdevjson**
produces JSON output, making it easier to parse, validate, and integrate
with scripts or other programs.

# OPTIONS
**-h**, **--help**
:   Show usage information and exit.

**--pretty**
:   Replace numeric event codes in the JSON output with their symbolic
    names (e.g., "30" to "30 (KEY_A)"). This makes the data easier to
    interpret by humans but does not change the structure or formatting -
    the output remains compact, machine-readable JSON. For visually
    formatted output, pipe to a JSON pretty-printer such as `jq`.

# EXAMPLES
Monitor events from `/dev/input/event3` with numeric codes:
```bash
evdevjson monitor /dev/input/event3
```

Monitor events with symbolic codes, piped to `jq` for readability:
```bash
evdevjson monitor --pretty /dev/input/event3 | jq
```

Snapshot the current state of `/dev/input/event5` with symbolic codes:
```bash
evdevjson snapshot --pretty /dev/input/event5 | jq
```

Snapshot multiple devices with symbolic codes:
```bash
evdevjson snapshot --pretty /dev/input/event3 /dev/input/event5 | jq
```

Snapshot all available devices (no path arguments):
```bash
evdevjson snapshot --pretty | jq
```

# EXIT STATUS
**0** - Successful execution\
**1** - An error occurred

# AUTHOR
Kris Andrie Ortega (andrieee44@gmail.com)

# COPYRIGHT
Copyright (C) 2025 Kris Andrie Ortega.\
License: GNU General Public License v3.0.
