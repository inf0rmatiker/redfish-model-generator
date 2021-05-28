/* -----------------------------------------------------------------
 * enum_network_device_technology.go -
 *
 * DMTF Redfish NetworkDeviceTechnology enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type NetworkDeviceTechnology string

const (
	// Neither enumerated nor visible to the operating system.
	NetworkDeviceTechnology_DISABLED NetworkDeviceTechnology = "Disabled"

	// Appears to the operating system as an Ethernet device.
	NetworkDeviceTechnology_ETHERNET NetworkDeviceTechnology = "Ethernet"

	// Appears to the operating system as a Fibre Channel device.
	NetworkDeviceTechnology_FIBRE_CHANNEL NetworkDeviceTechnology = "FibreChannel"

	// Appears to the operating system as an iSCSI device.
	NetworkDeviceTechnology_I_SCSI NetworkDeviceTechnology = "iSCSI"

	// Appears to the operating system as an FCoE device.
	NetworkDeviceTechnology_FIBRE_CHANNEL_OVER_ETHERNET NetworkDeviceTechnology = "FibreChannelOverEthernet"

	// Appears to the operating system as an InfiniBand device.
	NetworkDeviceTechnology_INFINI_BAND NetworkDeviceTechnology = "InfiniBand"

)
