/* -----------------------------------------------------------------
 * enum_facility_type.go -
 *
 * DMTF Redfish FacilityType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type FacilityType string

const (
	// A room inside of a building or floor.
	FacilityType_ROOM FacilityType = "Room"

	// A floor inside of a building.
	FacilityType_FLOOR FacilityType = "Floor"

	// A structure with a roof and walls.
	FacilityType_BUILDING FacilityType = "Building"

	// A small area consisting of several buildings.
	FacilityType_SITE FacilityType = "Site"

)
