/* -----------------------------------------------------------------
* ana_characteristics.go -
*
* DMTF Redfish ANACharacteristics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ANACharacteristics - The ANA characteristics and volume information for a storage controller.
// Modeled after DMTF Redfish schema ANACharacteristics
type ANACharacteristics struct {
	// Reported ANA access state.
	AccessState ANAAccessState `json:"AccessState,omitempty"`

	// The specified volume.
	Volume map[string]interface{} `json:"Volume,omitempty"`

}