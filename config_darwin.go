package configdir

import "os"
import "strings"

var hasVendorName = true

var systemSettingFolders []string
var globalSettingFolder string
var cacheFolder string

func init() {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		globalSettingFolder = os.Getenv("XDG_CONFIG_HOME")
	} else {
		globalSettingFolder = os.Getenv("HOME") + "/Library/Application Support"
	}
	if os.Getenv("XDG_CONFIG_DIRS") != "" {
		systemSettingFolders = strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")
	} else {
		systemSettingFolders = []string{"/Library/Application Support"}
	}
	if os.Getenv("XDG_CACHE_HOME") != "" {
		cacheFolder = os.Getenv("XDG_CACHE_HOME")
	} else {
		cacheFolder = os.Getenv("HOME") + "/Library/Caches"
	}
}
