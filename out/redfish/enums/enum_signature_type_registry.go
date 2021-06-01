/* -----------------------------------------------------------------
* enum_signature_type_registry.go -
*
* DMTF Redfish SignatureTypeRegistry enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SignatureTypeRegistry string

const (
	// A signature defined in the UEFI Specification.
	SignatureTypeRegistry_UEFI SignatureTypeRegistry = "UEFI"

)
