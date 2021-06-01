/* -----------------------------------------------------------------
* gen_z_connection_key.go -
*
* DMTF Redfish GenZConnectionKey resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// GenZConnectionKey - The Gen-Z-specific permission key information for a connection.
// Modeled after DMTF Redfish schema GenZConnectionKey
type GenZConnectionKey struct {
	// The Access Key for this connection.
	AccessKey string `json:"AccessKey,omitempty"`

	// Indicates whether Region Key domain checking is enabled for this connection.
	RKeyDomainCheckingEnabled bool `json:"RKeyDomainCheckingEnabled,omitempty"`

	// The read-only Region Key for this connection.
	RKeyReadOnlyKey string `json:"RKeyReadOnlyKey,omitempty"`

	// The read-write Region Key for this connection.
	RKeyReadWriteKey string `json:"RKeyReadWriteKey,omitempty"`

}
