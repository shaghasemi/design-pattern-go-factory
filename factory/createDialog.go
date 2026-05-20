package factory

import "runtime"

func CreateDialog() Dialog {
	println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		return &WindowsDialog{}
	}
	return &LinuxDialog{}
}
