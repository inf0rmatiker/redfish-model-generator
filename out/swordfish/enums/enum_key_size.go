/* -----------------------------------------------------------------
* enum_key_size.go -
*
* SNIA Swordfish KeySize enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type KeySize string

const (
	// No key.
	KeySize_BITS_0 KeySize = "Bits_0"

	// 3DES 112 bit key.
	KeySize_BITS_112 KeySize = "Bits_112"

	// AES 128 bit key.
	KeySize_BITS_128 KeySize = "Bits_128"

	// AES 192 bit key.
	KeySize_BITS_192 KeySize = "Bits_192"

	// AES 256 bit key.
	KeySize_BITS_256 KeySize = "Bits_256"

)
