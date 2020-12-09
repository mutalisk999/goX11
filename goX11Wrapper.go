package goX11

// #cgo CFLAGS: -Ix11
// #cgo LDFLAGS: x11/libx11.a
// #include "x11.h"
import "C"
import (
	"encoding/hex"
	"unsafe"
)

func CalcX11Hash(bytesInput []byte) []byte {
	// avoid of the situation of empty bytes array
	bytesInput = append(bytesInput, byte(0x0))
	input := (*C.char)(unsafe.Pointer(&bytesInput[0]))
	inputLen := C.uint(len(bytesInput) - 1)

	bytesOutput := make([]byte, 32)
	output := (*C.char)(unsafe.Pointer(&bytesOutput[0]))

	C.x11_hash(input, output, inputLen)

	return bytesOutput
}

func CalcX11HashHex(bytesInput []byte) string {
	bytesRet := CalcX11Hash(bytesInput)
	return hex.EncodeToString(bytesRet)
}
