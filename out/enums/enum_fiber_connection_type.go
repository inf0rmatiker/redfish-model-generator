/* -----------------------------------------------------------------
 * enum_fiber_connection_type.go -
 *
 * DMTF Redfish FiberConnectionType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type FiberConnectionType string

const (
	// The connection is using single mode operation.
	FiberConnectionType_SINGLE_MODE FiberConnectionType = "SingleMode"

	// The connection is using multi mode operation.
	FiberConnectionType_MULTI_MODE FiberConnectionType = "MultiMode"

)
