/* -----------------------------------------------------------------
* power.go -
*
* DMTF Redfish Power resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Power - This is the schema definition for the Power Metrics.  It represents the properties for Power Consumption and Power Limiting.
// Modeled after DMTF Redfish schema Power
type Power struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	PowerControlOdataCount int `json:"PowerControl@odata.count,omitempty"`

	PowerControl@odata.navigationLink string `json:"PowerControl@odata.navigationLink,omitempty"`

	// This is the definition for power control function (power reading/limiting).
	PowerControl []map[string]interface{} `json:"PowerControl,omitempty"`

	VoltagesOdataCount int `json:"Voltages@odata.count,omitempty"`

	Voltages@odata.navigationLink string `json:"Voltages@odata.navigationLink,omitempty"`

	// This is the definition for voltage sensors.
	Voltages []Voltage `json:"Voltages,omitempty"`

	PowerSuppliesOdataCount int `json:"PowerSupplies@odata.count,omitempty"`

	PowerSupplies@odata.navigationLink string `json:"PowerSupplies@odata.navigationLink,omitempty"`

	// Details of the power supplies associated with this system or device.
	PowerSupplies []PowerSupply `json:"PowerSupplies,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	Redundancy@odata.navigationLink string `json:"Redundancy@odata.navigationLink,omitempty"`

	// Redundancy information for the power subsystem of this system or device.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

}
