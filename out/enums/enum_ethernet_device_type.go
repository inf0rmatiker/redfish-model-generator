/* -----------------------------------------------------------------
 * enum_ethernet_device_type.go -
 *
 * DMTF Redfish EthernetDeviceType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type EthernetDeviceType string

const (
	// A physical Ethernet interface.
	EthernetDeviceType_PHYSICAL EthernetDeviceType = "Physical"

	// A virtual Ethernet interface.
	EthernetDeviceType_VIRTUAL EthernetDeviceType = "Virtual"

)
