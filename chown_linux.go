package lumberjack

import (
	"os"
	"syscall"
)

// os_Chown is a var so we can mock it out during tests.
var os_Chown = os.Chown

//根据名字和文件属性修改文件own
func chown(name string, info os.FileInfo) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	f.Close()
	//获取当前的文件stat属性
	stat := info.Sys().(*syscall.Stat_t)

	//修改文件的own
	return os_Chown(name, int(stat.Uid), int(stat.Gid))
}
