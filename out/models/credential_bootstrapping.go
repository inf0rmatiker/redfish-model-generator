/* -----------------------------------------------------------------
* credential_bootstrapping.go -
*
* DMTF Redfish CredentialBootstrapping resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The credential bootstrapping settings for this interface.
type CredentialBootstrapping struct {
	// An indication of whether credential bootstrapping is enabled after a reset for this interface.
	EnableAfterReset bool `json:"EnableAfterReset,omitempty"`

	// An indication of whether credential bootstrapping is enabled for this interface.
	Enabled bool `json:"Enabled,omitempty"`

	// The role used for the bootstrap account created for this interface.
	RoleId string `json:"RoleId,omitempty"`

}
