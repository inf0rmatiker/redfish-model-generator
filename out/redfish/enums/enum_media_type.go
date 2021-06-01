/* -----------------------------------------------------------------
* enum_media_type.go -
*
* DMTF Redfish MediaType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type MediaType string

const (
	// The drive media type is traditional magnetic platters.
	MediaType_HDD MediaType = "HDD"

	// The drive media type is solid state or flash memory.
	MediaType_SSD MediaType = "SSD"

	// The drive media type is shingled magnetic recording.
	MediaType_SMR MediaType = "SMR"

)
