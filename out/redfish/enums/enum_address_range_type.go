/* -----------------------------------------------------------------
* enum_address_range_type.go -
*
* DMTF Redfish AddressRangeType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AddressRangeType string

const (
	// Volatile memory.
	AddressRangeType_VOLATILE AddressRangeType = "Volatile"

	// Byte accessible persistent memory.
	AddressRangeType_PMEM AddressRangeType = "PMEM"

	// Block accesible memory.
	AddressRangeType_BLOCK AddressRangeType = "Block"

)
