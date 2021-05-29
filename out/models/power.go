/* -----------------------------------------------------------------
* power.go -
*
* DMTF Redfish Power resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Power schema describes power metrics and represents the properties for power consumption and power limiting.
type Power struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The set of power control functions, including power reading and limiting.
	PowerControl array `json:"PowerControl,omitempty"`

	PowerControl@odata.count count `json:"PowerControl@odata.count,omitempty"`

	// The set of power supplies associated with this system or device.
	PowerSupplies array `json:"PowerSupplies,omitempty"`

	PowerSupplies@odata.count count `json:"PowerSupplies@odata.count,omitempty"`

	// The redundancy information for the set of power supplies in this chassis.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The set of voltage sensors for this chassis.
	Voltages array `json:"Voltages,omitempty"`

	Voltages@odata.count count `json:"Voltages@odata.count,omitempty"`

}
