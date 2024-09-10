//go:build !windows && !darwin
// +build !windows,!darwin

package configdir

import (
	"os"
	"strings"
)

// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

var hasVendorName = true
var systemSettingFolders []string
var globalSettingFolder string
var cacheFolder string

func init() {

	globalSettingFolder, _ = os.UserConfigDir()

	if os.Getenv("XDG_CONFIG_DIRS") != "" {
		systemSettingFolders = strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")
	} else {
		systemSettingFolders = []string{"/etc/xdg"}
	}
	cacheFolder, _ = os.UserCacheDir()
}
