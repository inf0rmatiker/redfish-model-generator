/* -----------------------------------------------------------------
* redundancy.go -
*
* DMTF Redfish Redundancy resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The common redundancy definition and structure used in other Redfish schemas.
type Redundancy struct {
	OdataId string `json:"@odata.id"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The maximum number of members allowable for this particular redundancy group.
	MaxNumSupported int `json:"MaxNumSupported,omitempty"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The minimum number of members needed for this group to be redundant.
	MinNumNeeded int `json:"MinNumNeeded"`

	// The redundancy mode of the group.
	Mode RedundancyMode `json:"Mode"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether redundancy is enabled.
	RedundancyEnabled bool `json:"RedundancyEnabled,omitempty"`

	// The links to components of this redundancy set.
	RedundancySet []map[string]string `json:"RedundancySet"`

	RedundancySet@odata.count count `json:"RedundancySet@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status"`

}
