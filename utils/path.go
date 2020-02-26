package utils

import (
	"path"
	"runtime"
)

func GetAbsPath(fp string) (absPath string) {
	_, filename, _, _ := runtime.Caller(1)
	absPath = path.Join(path.Dir(filename), fp)
	return
}
