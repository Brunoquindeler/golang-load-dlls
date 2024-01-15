package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

const (
	PROCESS_ALL_ACCESS = 0x1F0FFF
)

var (
	kernel32DLL           = syscall.MustLoadDLL("kernel32.dll")
	procOpenProcess       = kernel32DLL.MustFindProc("OpenProcess")
	procReadProcessMemory = kernel32DLL.MustFindProc("ReadProcessMemory")
	procCloseHandle       = kernel32DLL.MustFindProc("CloseHandle")
)

var processID uint32

func init() {
	processID = 0 // Adicionar pid da aplicação a ser lida
}

func main() {
	processHandle, _, _ := procOpenProcess.Call(uintptr(PROCESS_ALL_ACCESS), 0, uintptr(processID))
	if processHandle == 0 {
		log.Fatal("falha ao abrir processo")
	}
	defer procCloseHandle.Call(processHandle)

	addressToRead := uintptr(0x0) // Adicionar endereço do ponteiro a ser lido

	for {
		var (
			buffer    int
			bufferPtr = (uintptr)(unsafe.Pointer(&buffer)) // Conversão de tipo seguro para não seguro para ter interoperabilidade
		)

		procReadProcessMemory.Call(
			processHandle,
			addressToRead,
			uintptr(bufferPtr),
			uintptr(unsafe.Sizeof(buffer)),
			0,
		)

		fmt.Printf("Valor lido: %d\n", buffer)

		time.Sleep(time.Second)
	}
}

// https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory
/*
C++
	BOOL ReadProcessMemory(
		[in]  HANDLE  hProcess,
		[in]  LPCVOID lpBaseAddress,
		[out] LPVOID  lpBuffer,
		[in]  SIZE_T  nSize,
		[out] SIZE_T  *lpNumberOfBytesRead
	);
*/
