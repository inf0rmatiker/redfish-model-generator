/* -----------------------------------------------------------------
* enum_pc_ie_types.go -
*
* DMTF Redfish PCIeTypes enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PCIeTypes string

const (
	// A PCIe v1.0 slot.
	PCIeTypes_GEN1 PCIeTypes = "Gen1"

	// A PCIe v2.0 slot.
	PCIeTypes_GEN2 PCIeTypes = "Gen2"

	// A PCIe v3.0 slot.
	PCIeTypes_GEN3 PCIeTypes = "Gen3"

	// A PCIe v4.0 slot.
	PCIeTypes_GEN4 PCIeTypes = "Gen4"

	// A PCIe v5.0 slot.
	PCIeTypes_GEN5 PCIeTypes = "Gen5"

)
