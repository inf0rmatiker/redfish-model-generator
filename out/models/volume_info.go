/* -----------------------------------------------------------------
* volume_info.go -
*
* DMTF Redfish VolumeInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The combination of permissions and volume information.
type VolumeInfo struct {
	// Supported IO access capabilities.
	AccessCapabilities array `json:"AccessCapabilities,omitempty"`

	// The access state for this connection.
	AccessState AccessState `json:"AccessState,omitempty"`

	// The specified volume.
	Volume Volume `json:"Volume,omitempty"`

}
