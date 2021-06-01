/* -----------------------------------------------------------------
* thermal_subsystem.go -
*
* DMTF Redfish ThermalSubsystem resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ThermalSubsystem - This ThermalSubsystem schema contains the definition for the thermal subsystem of a chassis.
// Modeled after DMTF Redfish schema ThermalSubsystem
type ThermalSubsystem struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The redundancy information for the groups of fans in this subsystem.
	FanRedundancy []RedundantGroup `json:"FanRedundancy,omitempty"`

	// The link to the collection of fans within this subsystem.
	Fans map[string]interface{} `json:"Fans,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The link to the summary of thermal metrics for this subsystem.
	ThermalMetrics map[string]interface{} `json:"ThermalMetrics,omitempty"`

}
