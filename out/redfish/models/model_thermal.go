/* -----------------------------------------------------------------
* thermal.go -
*
* DMTF Redfish Thermal resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Thermal - This is the schema definition for the Thermal properties.  It represents the properties for Temperature and Cooling.
// Modeled after DMTF Redfish schema Thermal
type Thermal struct {
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

	Status map[string]interface{} `json:"Status,omitempty"`

	TemperaturesOdataCount int `json:"Temperatures@odata.count,omitempty"`

	Temperatures@odata.navigationLink string `json:"Temperatures@odata.navigationLink,omitempty"`

	// This is the definition for temperature sensors.
	Temperatures []Temperature `json:"Temperatures,omitempty"`

	FansOdataCount int `json:"Fans@odata.count,omitempty"`

	Fans@odata.navigationLink string `json:"Fans@odata.navigationLink,omitempty"`

	// This is the definition for fans.
	Fans []Fan `json:"Fans,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	Redundancy@odata.navigationLink string `json:"Redundancy@odata.navigationLink,omitempty"`

	// This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

}
