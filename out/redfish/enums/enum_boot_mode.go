/* -----------------------------------------------------------------
* enum_boot_mode.go -
*
* DMTF Redfish BootMode enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type BootMode string

const (
	// Do not indicate to UEFI/BIOS that this device is bootable.
	BootMode_DISABLED BootMode = "Disabled"

	// Boot this device using the embedded PXE support.  Only applicable if the NetworkDeviceFunctionType is set to Ethernet.
	BootMode_PXE BootMode = "PXE"

	// Boot this device using the embedded iSCSI boot support and configuration.  Only applicable if the NetworkDeviceFunctionType is set to iSCSI.
	BootMode_I_SCSI BootMode = "iSCSI"

	// Boot this device using the embedded Fibre Channel support and configuration.  Only applicable if the NetworkDeviceFunctionType is set to FibreChannel.
	BootMode_FIBRE_CHANNEL BootMode = "FibreChannel"

	// Boot this device using the embedded Fibre Channel over Ethernet (FCoE) boot support and configuration.  Only applicable if the NetworkDeviceFunctionType is set to FibreChannelOverEthernet.
	BootMode_FIBRE_CHANNEL_OVER_ETHERNET BootMode = "FibreChannelOverEthernet"

)
