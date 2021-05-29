/* -----------------------------------------------------------------
* redundant_group.go -
*
* DMTF Redfish RedundantGroup resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The redundancy information for the devices in a redundancy group.
type RedundantGroup struct {
	// The maximum number of devices supported in this redundancy group.
	MaxSupportedInGroup int `json:"MaxSupportedInGroup,omitempty"`

	// The minimum number of devices needed for this group to be redundant.
	MinNeededInGroup int `json:"MinNeededInGroup"`

	// The links to the devices included in this redundancy group.
	RedundancyGroup array `json:"RedundancyGroup"`

	RedundancyGroup@odata.count count `json:"RedundancyGroup@odata.count,omitempty"`

	// The redundancy mode of the group.
	RedundancyType RedundancyType `json:"RedundancyType"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status"`

}
