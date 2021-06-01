/* -----------------------------------------------------------------
* enum_initialize_type.go -
*
* DMTF Redfish InitializeType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type InitializeType string

const (
	// The volume is prepared for use quickly, typically by erasing just the beginning and end of the space so that partitioning can be performed.
	InitializeType_FAST InitializeType = "Fast"

	// The volume is prepared for use slowly, typically by completely erasing the volume.
	InitializeType_SLOW InitializeType = "Slow"

)
