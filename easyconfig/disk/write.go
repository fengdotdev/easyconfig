package disk

import (
	"easyconfig/pkg/errorplus"
)

func (d *Disk) Write(b []byte, overwrite ...bool) error {

	allowOverwrite := false

	if len(overwrite) == 1 && overwrite[0] {
		allowOverwrite = true
	}

	if allowOverwrite {
		return nil
	}
	 errmsg := "file already exists"
	return errorplus.ES(errmsg, d.Write)
}