/* -----------------------------------------------------------------
* enum_character_code_set.go -
*
* SNIA Swordfish CharacterCodeSet enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type CharacterCodeSet string

const (
	// ASCII character encoding.
	CharacterCodeSet_ASCII CharacterCodeSet = "ASCII"

	// Unicode character encoding.
	CharacterCodeSet_UNICODE CharacterCodeSet = "Unicode"

	// ISO-2022 character encoding.
	CharacterCodeSet_ISO2022 CharacterCodeSet = "ISO2022"

	// ISO-8859-1 character encoding.
	CharacterCodeSet_ISO8859_1 CharacterCodeSet = "ISO8859_1"

	// Extended Unix Code encoding.
	CharacterCodeSet_EXTENDED_UNIX_CODE CharacterCodeSet = "ExtendedUNIXCode"

	// UTF-8 character encoding.
	CharacterCodeSet_UTF_8 CharacterCodeSet = "UTF_8"

	// UTF-16 character encoding.
	CharacterCodeSet_UTF_16 CharacterCodeSet = "UTF_16"

	// UCS-2 character encoding.
	CharacterCodeSet_UCS_2 CharacterCodeSet = "UCS_2"

)
