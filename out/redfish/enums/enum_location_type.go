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
	// Defines a slot as the type of location.
	LocationType_SLOT LocationType = "Slot"

	// Defines a bay as the type of location.
	LocationType_BAY LocationType = "Bay"

	// Defines a connector as the type of location.
	LocationType_CONNECTOR LocationType = "Connector"

	// Defines a socket as the type of location.
	LocationType_SOCKET LocationType = "Socket"

)
