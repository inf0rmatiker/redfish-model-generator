/* -----------------------------------------------------------------
* enum_connection_type.go -
*
* DMTF Redfish ConnectionType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ConnectionType string

const (
	// A connection to storage related resources, such as volumes.
	ConnectionType_STORAGE ConnectionType = "Storage"

	// A connection to memory related resources.
	ConnectionType_MEMORY ConnectionType = "Memory"

)
