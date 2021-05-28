/* -----------------------------------------------------------------
 * enum_snmp_community_access_mode.go -
 *
 * DMTF Redfish SNMPCommunityAccessMode enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SNMPCommunityAccessMode string

const (
	// READ-WRITE access mode.
	SNMPCommunityAccessMode_FULL SNMPCommunityAccessMode = "Full"

	// READ-ONLY access mode.
	SNMPCommunityAccessMode_LIMITED SNMPCommunityAccessMode = "Limited"

)
