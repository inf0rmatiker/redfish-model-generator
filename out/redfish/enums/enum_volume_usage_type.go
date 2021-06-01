/* -----------------------------------------------------------------
* enum_volume_usage_type.go -
*
* DMTF Redfish VolumeUsageType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type VolumeUsageType string

const (
	// The volume is allocated for use as a consumable data volume.
	VolumeUsageType_DATA VolumeUsageType = "Data"

	// The volume is allocated for use as a consumable data volume reserved for system use.
	VolumeUsageType_SYSTEM_DATA VolumeUsageType = "SystemData"

	// The volume is allocated for use as a non-consumable cache only volume.
	VolumeUsageType_CACHE_ONLY VolumeUsageType = "CacheOnly"

	// The volume is allocated for use as a non-consumable system reserved volume.
	VolumeUsageType_SYSTEM_RESERVE VolumeUsageType = "SystemReserve"

	// The volume is allocated for use as a non-consumable reserved volume for replication use.
	VolumeUsageType_REPLICATION_RESERVE VolumeUsageType = "ReplicationReserve"

)
