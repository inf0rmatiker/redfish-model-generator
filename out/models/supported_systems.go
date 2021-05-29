/* -----------------------------------------------------------------
* supported_systems.go -
*
* DMTF Redfish SupportedSystems resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A system that this attribute registry supports.
type SupportedSystems struct {
	// Firmware version.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The product name of the computer system to which this attribute registry applies.
	ProductName string `json:"ProductName,omitempty"`

	// The ID of the systems to which this attribute registry applies.
	SystemId string `json:"SystemId,omitempty"`

}
