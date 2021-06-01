/* -----------------------------------------------------------------
* power_subsystem.go -
*
* DMTF Redfish PowerSubsystem resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerSubsystem - This PowerSubsystem schema contains the definition for the power subsystem of a chassis.
// Modeled after DMTF Redfish schema PowerSubsystem
type PowerSubsystem struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Power allocation for this subsystem.
	Allocation PowerAllocation `json:"Allocation,omitempty"`

	// The total amount of power that can be allocated to this subsystem.  This value can be either the power supply capacity or the power budget that an upstream chassis assigns to this subsystem.
	CapacityWatts float64 `json:"CapacityWatts,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to the collection of power supplies within this subsystem.
	PowerSupplies map[string]interface{} `json:"PowerSupplies,omitempty"`

	// The redundancy information for the set of power supplies in this subsystem.
	PowerSupplyRedundancy []RedundantGroup `json:"PowerSupplyRedundancy,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
