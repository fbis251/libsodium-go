package cryptopwhash

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"

func OpslimitInteractive() int {
	return int(C.crypto_pwhash_opslimit_interactive())
}

func MemlimitInteractive() int {
	return int(C.crypto_pwhash_memlimit_interactive())
}
