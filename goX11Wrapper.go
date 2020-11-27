package goX11

// #cgo CFLAGS: -Ix11/x11
// #cgo LDFLAGS: x11/libx11.a
// #include "x11/x11.h"
import "C"
import (
	"encoding/hex"
	"unsafe"
)

func CalcX11Hash(bytesInput []byte) []byte {
	input := C.CString(string(bytesInput))
	defer C.free(unsafe.Pointer(input))

	inputLen := C.uint(len(bytesInput))

	bytesOutput := make([]byte, 32)
	output := (*C.char)(unsafe.Pointer(&bytesOutput[0]))

	C.x11_hash(input, output, inputLen)

	return bytesOutput
}

func CalcX11HashHex(bytesInput []byte) string {
	bytesRet := CalcX11Hash(bytesInput)
	return hex.EncodeToString(bytesRet)
}
