/* -----------------------------------------------------------------
* mapped_volume.go -
*
* SNIA Swordfish MappedVolume resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MappedVolume - Relate a SCSI Logical Unit Number to a Volume.
// Modeled after SNIA Swordfish schema MappedVolume
type MappedVolume struct {
	// Supported IO access capability.
	AccessCapability AccessCapability `json:"AccessCapability,omitempty"`

	// A SCSI Logical Unit Number for a Volume.
	LogicalUnitNumber string `json:"LogicalUnitNumber,omitempty"`

	// A mapped Volume.
	Volume map[string]interface{} `json:"Volume,omitempty"`

}
