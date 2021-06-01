/* -----------------------------------------------------------------
* enum_hotspare_type.go -
*
* DMTF Redfish HotspareType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type HotspareType string

const (
	// The drive is not currently a hot spare.
	HotspareType_NONE HotspareType = "None"

	// The drive is currently serving as a hot spare for all other drives in the storage system.
	HotspareType_GLOBAL HotspareType = "Global"

	// The drive is currently serving as a hot spare for all other drives in the chassis.
	HotspareType_CHASSIS HotspareType = "Chassis"

	// The drive is currently serving as a hot spare for a user-defined set of drives.
	HotspareType_DEDICATED HotspareType = "Dedicated"

)
