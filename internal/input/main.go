package input

import (
	"os"
)

var (
	Release    = get("INPUT_RELEASE", "15.0-RELEASE")
	Arch       = get("INPUT_ARCH", "amd64")
	Filesystem = get("INPUT_FILESYSTEM", "zfs")
	Memory     = get("INPUT_MEMORY", "4G")
	Run        = get("INPUT_RUN", "uname -a")
)

func get(key, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	} else {
		return v
	}
}
