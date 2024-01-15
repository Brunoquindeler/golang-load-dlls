package main

import (
	"syscall"
)

var (
	kernel32DLL               = syscall.MustLoadDLL("kernel32.dll")
	procReadProcessFromMemory = kernel32DLL.MustFindProc("ReadProcessMemory")
	// kernel32DLL               = syscall.NewLazyDLL("kernel32.dll")
	// procReadProcessFromMemory = kernel32DLL.NewProc("ReadProcessMemory")
)

func main() {

}
