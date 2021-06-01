/* -----------------------------------------------------------------
* enum_memory_classification.go -
*
* DMTF Redfish MemoryClassification enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type MemoryClassification string

const (
	// Volatile memory.
	MemoryClassification_VOLATILE MemoryClassification = "Volatile"

	// Byte-accessible persistent memory.
	MemoryClassification_BYTE_ACCESSIBLE_PERSISTENT MemoryClassification = "ByteAccessiblePersistent"

	// Block-accessible memory.
	MemoryClassification_BLOCK MemoryClassification = "Block"

)
