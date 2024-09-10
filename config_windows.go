package configdir

import "os"

var hasVendorName = true
var systemSettingFolders = []string{os.Getenv("PROGRAMDATA")}
var globalSettingFolder, _ = os.UserConfigDir()
var cacheFolder, _ = os.UserCacheDir()
