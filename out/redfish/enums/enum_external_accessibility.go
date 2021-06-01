/* -----------------------------------------------------------------
* enum_external_accessibility.go -
*
* DMTF Redfish ExternalAccessibility enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ExternalAccessibility string

const (
	// Any external entity with the correct access details, which may include authorization information, can access the endpoints that this zone lists.
	ExternalAccessibility_GLOBALLY_ACCESSIBLE ExternalAccessibility = "GloballyAccessible"

	// Any external entity that another zone does not explicitly list can access the endpoints that this zone lists.
	ExternalAccessibility_NON_ZONED_ACCESSIBLE ExternalAccessibility = "NonZonedAccessible"

	// Only accessible by endpoints that this zone explicitly lists.
	ExternalAccessibility_ZONE_ONLY ExternalAccessibility = "ZoneOnly"

)
