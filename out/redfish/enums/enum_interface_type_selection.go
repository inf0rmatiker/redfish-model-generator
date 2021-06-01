/* -----------------------------------------------------------------
* enum_interface_type_selection.go -
*
* DMTF Redfish InterfaceTypeSelection enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type InterfaceTypeSelection string

const (
	// The TrustedModule does not support switching the InterfaceType.
	InterfaceTypeSelection_NONE InterfaceTypeSelection = "None"

	// The TrustedModule supports switching InterfaceType through a firmware update.
	InterfaceTypeSelection_FIRMWARE_UPDATE InterfaceTypeSelection = "FirmwareUpdate"

	// The TrustedModule supports switching InterfaceType through platform software, such as a BIOS configuration attribute.
	InterfaceTypeSelection_BIOS_SETTING InterfaceTypeSelection = "BiosSetting"

	// The TrustedModule supports switching InterfaceType through an OEM proprietary mechanism.
	InterfaceTypeSelection_OEM_METHOD InterfaceTypeSelection = "OemMethod"

)
