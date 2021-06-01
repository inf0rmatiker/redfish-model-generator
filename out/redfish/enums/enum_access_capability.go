/* -----------------------------------------------------------------
* enum_access_capability.go -
*
* DMTF Redfish AccessCapability enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AccessCapability string

const (
	// Endpoints are allowed to perform reads from the specified resource.
	AccessCapability_READ AccessCapability = "Read"

	// Endpoints are allowed to perform writes to the specified resource.
	AccessCapability_WRITE AccessCapability = "Write"

)
