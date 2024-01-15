# Golang Load DLLs

## **[ GOLANG ]** Carregando DLLs e exemplo de uso com ReadProcessMemory da kernel32.dll

Os códigos neste repositório são dos exemplos passados no vídeo de explicação de como carregar DLLs com Golang e um exemplo de como utiliza-la.

___

[YouTube Video](https://youtu.be/Dobo_jLt50c)
[![YouTube Thumb](thumb.png)](https://youtu.be/Dobo_jLt50c)
 ___

[Doc ReadProcessMemory](https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory) 
```cpp
BOOL ReadProcessMemory(
    [in]  HANDLE  hProcess,
    [in]  LPCVOID lpBaseAddress,
    [out] LPVOID  lpBuffer,
    [in]  SIZE_T  nSize,
    [out] SIZE_T  *lpNumberOfBytesRead
);
```

___

[Doc WriteProcessMemory](https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory)
```cpp
BOOL WriteProcessMemory(
	[in]  HANDLE  hProcess,
	[in]  LPVOID  lpBaseAddress,
	[in]  LPCVOID lpBuffer,
	[in]  SIZE_T  nSize,
	[out] SIZE_T  *lpNumberOfBytesWritten
);
```
