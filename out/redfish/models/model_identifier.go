/* -----------------------------------------------------------------
* identifier.go -
*
* DMTF Redfish Identifier resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Identifier - This type describes any additional identifiers for a resource.
// Modeled after DMTF Redfish schema Identifier
type Identifier struct {
	// This indicates the world wide, persistent name of the resource.
	DurableName string `json:"DurableName,omitempty"`

	// This represents the format of the DurableName property.
	DurableNameFormat DurableNameFormat `json:"DurableNameFormat,omitempty"`

}
