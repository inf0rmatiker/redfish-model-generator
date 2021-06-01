/* -----------------------------------------------------------------
* trusted_modules.go -
*
* DMTF Redfish TrustedModules resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// TrustedModules - The Trusted Module installed in the system.
// Modeled after DMTF Redfish schema TrustedModules
type TrustedModules struct {
	// The firmware version of this Trusted Module.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The second firmware version of this Trusted Module, if applicable.
	FirmwareVersion2 string `json:"FirmwareVersion2,omitempty"`

	// The interface type of the Trusted Module.
	InterfaceType InterfaceType `json:"InterfaceType,omitempty"`

	// The interface type selection supported by this Trusted Module.
	InterfaceTypeSelection InterfaceTypeSelection `json:"InterfaceTypeSelection,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
