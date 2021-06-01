/* -----------------------------------------------------------------
* supported_systems.go -
*
* DMTF Redfish SupportedSystems resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SupportedSystems - A system supported by this attribute registry.
// Modeled after DMTF Redfish schema SupportedSystems
type SupportedSystems struct {
	// Firmware version.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The product name of the system.
	ProductName string `json:"ProductName,omitempty"`

	// The system ID of the system.
	SystemId string `json:"SystemId,omitempty"`

}
