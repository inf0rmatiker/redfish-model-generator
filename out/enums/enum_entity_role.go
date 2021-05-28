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
	// The entity sends commands, messages, or other types of requests to other entities on the fabric, but cannot receive commands from other entities.
	EntityRole_INITIATOR EntityRole = "Initiator"

	// The entity receives commands, messages, or other types of requests from other entities on the fabric, but cannot send commands to other entities.
	EntityRole_TARGET EntityRole = "Target"

	// The entity can both send and receive commands, messages, and other requests to or from other entities on the fabric.
	EntityRole_BOTH EntityRole = "Both"

)
