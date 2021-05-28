/* -----------------------------------------------------------------
 * enum_location_type.go -
 *
 * DMTF Redfish LocationType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type LocationType string

const (
	// A slot.
	LocationType_SLOT LocationType = "Slot"

	// A bay.
	LocationType_BAY LocationType = "Bay"

	// A connector or port.
	LocationType_CONNECTOR LocationType = "Connector"

	// A socket.
	LocationType_SOCKET LocationType = "Socket"

)
