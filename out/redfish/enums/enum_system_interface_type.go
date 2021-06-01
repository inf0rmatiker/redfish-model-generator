/* -----------------------------------------------------------------
* enum_system_interface_type.go -
*
* DMTF Redfish SystemInterfaceType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SystemInterfaceType string

const (
	// The Intel QuickPath Interconnect.
	SystemInterfaceType_QPI SystemInterfaceType = "QPI"

	// The Intel UltraPath Interconnect.
	SystemInterfaceType_UPI SystemInterfaceType = "UPI"

	// A PCI Express interface.
	SystemInterfaceType_PC_IE SystemInterfaceType = "PCIe"

	// An Ethernet interface.
	SystemInterfaceType_ETHERNET SystemInterfaceType = "Ethernet"

	// An OEM-defined interface.
	SystemInterfaceType_OEM SystemInterfaceType = "OEM"

)
