/* -----------------------------------------------------------------
* enum_entity_role.go -
*
* DMTF Redfish EntityRole enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EntityRole string

const (
	// The entity is acting as an initiator.
	EntityRole_INITIATOR EntityRole = "Initiator"

	// The entity is acting as a target.
	EntityRole_TARGET EntityRole = "Target"

	// The entity is acting as both an initiator and a target.
	EntityRole_BOTH EntityRole = "Both"

)
