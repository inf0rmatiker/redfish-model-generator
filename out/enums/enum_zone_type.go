/* -----------------------------------------------------------------
 * enum_zone_type.go -
 *
 * DMTF Redfish ZoneType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ZoneType string

const (
	// The zone in which all endpoints are added by default when instantiated.
	ZoneType_DEFAULT ZoneType = "Default"

	// A zone that contains endpoints.
	ZoneType_ZONE_OF_ENDPOINTS ZoneType = "ZoneOfEndpoints"

	// A zone that contains zones.
	ZoneType_ZONE_OF_ZONES ZoneType = "ZoneOfZones"

	// A zone that contains resource blocks.
	ZoneType_ZONE_OF_RESOURCE_BLOCKS ZoneType = "ZoneOfResourceBlocks"

)
