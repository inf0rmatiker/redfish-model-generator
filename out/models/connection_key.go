/* -----------------------------------------------------------------
* connection_key.go -
*
* DMTF Redfish ConnectionKey resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The permission key information required to access the target resources for a connection.
type ConnectionKey struct {
	// The Gen-Z-specific permission key information for this connection.
	GenZ GenZConnectionKey `json:"GenZ,omitempty"`

}
