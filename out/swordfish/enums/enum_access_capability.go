/* -----------------------------------------------------------------
* enum_access_capability.go -
*
* SNIA Swordfish AccessCapability enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AccessCapability string

const (
	// Endpoints are allowed to perform reads from the specified resource.
	AccessCapability_READ AccessCapability = "Read"

	// Endpoints are allowed to perform reads from and writes to the specified resource.
	AccessCapability_READ_WRITE AccessCapability = "ReadWrite"

)
