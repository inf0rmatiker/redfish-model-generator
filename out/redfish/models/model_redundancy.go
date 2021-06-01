/* -----------------------------------------------------------------
* redundancy.go -
*
* DMTF Redfish Redundancy resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Redundancy - This is the redundancy definition to be used in other resource schemas.
// Modeled after DMTF Redfish schema Redundancy
type Redundancy struct {
	OdataId map[string]interface{} `json:"@odata.id"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// This is the maximum number of members allowable for this particular redundancy group.
	MaxNumSupported int `json:"MaxNumSupported"`

	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// This is the minumum number of members needed for this group to be redundant.
	MinNumNeeded int `json:"MinNumNeeded"`

	// This is the redundancy mode of the group.
	Mode RedundancyMode `json:"Mode"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This indicates whether redundancy is enabled.
	RedundancyEnabled bool `json:"RedundancyEnabled,omitempty"`

	// Contains any ids that represent components of this redundancy set.
	RedundancySet []idRef `json:"RedundancySet"`

	RedundancySetOdataCount int `json:"RedundancySet@odata.count,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status"`

}
