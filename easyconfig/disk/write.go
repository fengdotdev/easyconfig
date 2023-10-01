package disk

import (
	"easyconfig/easyconfig/errorplus"
)

func (d *Disk) Write(b []byte, overwrite ...bool) error {

	allowOverwrite := false

	if len(overwrite) == 1 && overwrite[0] {
		allowOverwrite = true
	}

	if allowOverwrite {
		return nil
	}

	return errorplus.EString("file already exists", d.Write)
}