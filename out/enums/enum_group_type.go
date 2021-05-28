/* -----------------------------------------------------------------
 * enum_group_type.go -
 *
 * DMTF Redfish GroupType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type GroupType string

const (
	// The group contains the client (initiator) endpoints.
	GroupType_CLIENT GroupType = "Client"

	// The group contains the server (target) endpoints.
	GroupType_SERVER GroupType = "Server"

	// The group contains the initiator endpoints.
	GroupType_INITIATOR GroupType = "Initiator"

	// The group contains the target endpoints.
	GroupType_TARGET GroupType = "Target"

)
