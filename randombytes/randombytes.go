package randombytes

import "unsafe"

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"

// The RandomBytesBuf function fills size bytes starting at buf with an unpredictable sequence of bytes.
func RandomBytesBuf(buf []byte) {
	buflen := len(buf)
	C.randombytes_buf(unsafe.Pointer(&buf[0]),
		C.size_t(buflen))
}
