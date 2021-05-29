/* -----------------------------------------------------------------
* thermal_subsystem.go -
*
* DMTF Redfish ThermalSubsystem resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This ThermalSubsystem schema contains the definition for the thermal subsystem of a chassis.
type ThermalSubsystem struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The redundancy information for the groups of fans in this subsystem.
	FanRedundancy array `json:"FanRedundancy,omitempty"`

	// The link to the collection of fans within this subsystem.
	Fans FanCollection `json:"Fans,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The link to the summary of thermal metrics for this subsystem.
	ThermalMetrics ThermalMetrics `json:"ThermalMetrics,omitempty"`

}
