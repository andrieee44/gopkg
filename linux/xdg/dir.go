package xdg

import (
	"fmt"
	"os"
	"path/filepath"
)

func home() string {
	var home string

	home = os.Getenv("HOME")
	if home == "" {
		return "/"
	}

	return home
}

func xdg(env string, subPaths ...string) string {
	env = os.Getenv(env)
	if env == "" || !filepath.IsAbs(env) {
		env = filepath.Join(subPaths...)
	}

	return env
}

func xdgFile(xdgPath, relPath string) (*os.File, error) {
	const userOnly os.FileMode = 0o700

	var (
		file *os.File
		path string
		err  error
	)

	path = filepath.Join(xdgPath, relPath)

	err = os.MkdirAll(filepath.Dir(path), userOnly)
	if err != nil {
		return nil, fmt.Errorf("xdg.xdgFile: %w", err)
	}

	file, err = os.OpenFile(filepath.Clean(path), os.O_RDWR|os.O_CREATE, userOnly)
	if err != nil {
		return nil, fmt.Errorf("xdg.xdgFile: %w", err)
	}

	return file, nil
}

// DataFile opens the file with read/write access using a relative path
// (e.g., "appname/app.data") that includes the filename and optional
// directories. Missing directories are auto-created relative to the
// base data directory, and any errors return details about the attempted
// paths. Don't forget to call *[os.File.Close]() after use.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_DATA_HOME defines the base directory relative to which user-specific
// data files should be stored. If $XDG_DATA_HOME is either not set or empty,
// a default equal to $HOME/.local/share should be used.
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func DataFile(relPath string) (*os.File, error) {
	return xdgFile(xdg("XDG_DATA_HOME", home(), ".local/share"), relPath)
}

// ConfigFile opens the file with read/write access using a relative path
// (e.g., "appname/app.config") that includes the filename and optional
// directories. Missing directories are auto-created relative to the base
// config directory, and any errors return details about the attempted paths.
// Don't forget to call *[os.File.Close]() after use.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_CONFIG_HOME defines the base directory relative to which
// user-specific configuration files should be stored. If $XDG_CONFIG_HOME
// is either not set or empty, a default equal to $HOME/.config should be
// used.
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func ConfigFile(relPath string) (*os.File, error) {
	return xdgFile(xdg("XDG_CONFIG_HOME", home(), ".config"), relPath)
}

// StateFile opens the file with read/write access using a relative path
// (e.g., "appname/app.state") that includes the filename and optional
// directories. Missing directories are auto-created relative to the base
// state directory, and any errors return details about the attempted paths.
// Don't forget to call *[os.File.Close]() after use.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_STATE_HOME defines the base directory relative to which
// user-specific state files should be stored. If $XDG_STATE_HOME is either
// not set or empty, a default equal to $HOME/.local/state should be used.
//
// The $XDG_STATE_HOME contains state data that should persist between
// (application) restarts, but that is not important or portable enough to
// the user that it should be stored in $XDG_DATA_HOME.
//
// It may contain:
//
// - actions history (logs, history, recently used files, ...)
//
// - current state of the application that can be reused on a restart
// (view, layout, open files, undo history, ...)
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func StateFile(relPath string) (*os.File, error) {
	return xdgFile(xdg("XDG_STATE_HOME", home(), ".local/state"), relPath)
}

// DataDirs retrieves the value of $XDG_DATA_DIRS if it is defined,
// non-empty, and points to an absolute relPath; otherwise, it returns
// /usr/local/share/:/usr/share/ which is the default value.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_DATA_DIRS defines the preference-ordered set of base directories
// to search for data files in addition to the $XDG_DATA_HOME base directory.
// The directories in $XDG_DATA_DIRS should be separated with a colon ':'.
//
// If $XDG_DATA_DIRS is either not set or empty, a value equal to
// /usr/local/share/:/usr/share/ should be used.
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func DataDirs() string {
	return xdg("XDG_DATA_DIRS", "/usr/local/share/:/usr/share/")
}

// ConfigDirs retrieves the value of $XDG_CONFIG_DIRS if it is defined,
// non-empty, and points to an absolute relPath; otherwise, it returns
// /etc/xdg which is the default value.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_CONFIG_DIRS defines the preference-ordered set of base directories
// to search for configuration files in addition to the $XDG_CONFIG_HOME
// base directory. The directories in $XDG_CONFIG_DIRS should be separated
// with a colon ':'.
//
// If $XDG_CONFIG_DIRS is either not set or empty, a value equal to
// /etc/xdg should be used.
//
// The order of base directories denotes their importance; the first
// directory listed is the most important. When the same information is
// defined in multiple places the information defined relative to the
// more important base directory takes precedent. The base directory
// defined by $XDG_DATA_HOME is considered more important than any of the
// base directories defined by $XDG_DATA_DIRS. The base directory defined
// by $XDG_CONFIG_HOME is considered more important than any of the base
// directories defined by $XDG_CONFIG_DIRS.
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func ConfigDirs() string {
	return xdg("XDG_CONFIG_DIRS", "/etc/xdg")
}

// CacheFile opens the file with read/write access using a relative path
// (e.g., "appname/app.cache") that includes the filename and optional
// directories. Missing directories are auto-created relative to the
// base cache directory, and any errors return details about the attempted
// paths. Don't forget to call *[os.File.Close]() after use.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_CACHE_HOME defines the base directory relative to which
// user-specific non-essential data files should be stored. If
// $XDG_CACHE_HOME is either not set or empty, a default equal to
// $HOME/.cache should be used.
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func CacheFile(relPath string) (*os.File, error) {
	return xdgFile(xdg("XDG_CACHE_HOME", home(), "$HOME/.cache"), relPath)
}

// RuntimeFile opens the file with read/write access using a relative
// path (e.g., "appname/app.runtime") that includes the filename and
// optional directories. Missing directories are auto-created relative
// to the base runtime directory, and any errors return details about the
// attempted paths. Don't forget to call *[os.File.Close]() after use.
//
// From the [XDG Base Directory Specification]:
//
// $XDG_RUNTIME_DIR defines the base directory relative to which
// user-specific non-essential runtime files and other file objects
// (such as sockets, named pipes, ...) should be stored. The directory
// MUST be owned by the user, and they MUST be the only one having read
// and write access to it. Its Unix access mode MUST be 0700.
//
// The lifetime of the directory MUST be bound to the user being logged in.
// It MUST be created when the user first logs in and if the user fully
// logs out the directory MUST be removed. If the user logs in more than
// once they should get pointed to the same directory, and it is mandatory
// that the directory continues to exist from their first login to their
// last logout on the system, and not removed in between. Files in the
// directory MUST not survive reboot or a full logout/login cycle.
//
// The directory MUST be on a local file system and not shared with any
// other system. The directory MUST by fully-featured by the standards of
// the operating system. More specifically, on Unix-like operating systems
// AF_UNIX sockets, symbolic links, hard links, proper permissions,
// file locking, sparse files, memory mapping, file change notifications,
// a reliable hard link count must be supported, and no restrictions on the
// file name character set should be imposed. Files in this directory MAY
// be subjected to periodic clean-up. To ensure that your files are not
// removed, they should have their access time timestamp modified at
// least once every 6 hours of monotonic time or the 'sticky' bit should
// be set on the file.
//
// If $XDG_RUNTIME_DIR is not set applications should fall back to a
// replacement directory with similar capabilities and print a warning
// message. Applications should use this directory for communication and
// synchronization purposes and should not place larger files in it, since
// it might reside in runtime memory and cannot necessarily be swapped out
// to disk.
//
// [XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest
func RuntimeFile(relPath string) (*os.File, error) {
	return xdgFile(xdg("XDG_RUNTIME_DIR", "/tmp"), relPath)
}
