/* -----------------------------------------------------------------
* enum_stanza_type.go -
*
* DMTF Redfish StanzaType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type StanzaType string

const (
	// A stanza that describes the desired end state for computer system composition operation.  The resources consumed by the composed computer system are moved to the active pool.
	StanzaType_COMPOSE_SYSTEM StanzaType = "ComposeSystem"

	// A stanza that references a computer system to decompose and return resources to the free pool.
	StanzaType_DECOMPOSE_SYSTEM StanzaType = "DecomposeSystem"

	// A stanza that describes the desired end state for a composed resource block.  The resources consumed by the composed resource block are moved to the active pool.
	StanzaType_COMPOSE_RESOURCE StanzaType = "ComposeResource"

	// A stanza that references a composed resource block to decompose and return resources to the free pool.
	StanzaType_DECOMPOSE_RESOURCE StanzaType = "DecomposeResource"

	// A stanza that describes an OEM-specific request.
	StanzaType_OEM StanzaType = "OEM"

)
