// +build !linux

package lumberjack

import (
	"os"
)

//修改own
func chown(_ string, _ os.FileInfo) error {
	return nil
}
