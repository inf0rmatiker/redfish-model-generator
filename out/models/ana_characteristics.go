/* -----------------------------------------------------------------
* ana_characteristics.go -
*
* DMTF Redfish ANACharacteristics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ANA characteristics and volume information for a storage controller.
type ANACharacteristics struct {
	// Reported ANA access state.
	AccessState ANAAccessState `json:"AccessState,omitempty"`

	// The specified volume.
	Volume Volume `json:"Volume,omitempty"`

}
