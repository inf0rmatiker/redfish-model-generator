/* -----------------------------------------------------------------
* volume_info.go -
*
* DMTF Redfish VolumeInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VolumeInfo - The combination of permissions and volume information.
// Modeled after DMTF Redfish schema VolumeInfo
type VolumeInfo struct {
	// Supported IO access capabilities.
	AccessCapabilities []AccessCapability `json:"AccessCapabilities,omitempty"`

	// The access state for this connection.
	AccessState AccessState `json:"AccessState,omitempty"`

	// The specified volume.
	Volume map[string]interface{} `json:"Volume,omitempty"`

}
