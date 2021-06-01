/* -----------------------------------------------------------------
* assembly_data.go -
*
* DMTF Redfish AssemblyData resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	OdataId map[string]interface{} `json:"@odata.id"`

	// The available actions for this resource.
	Actions AssemblyDataActions `json:"Actions,omitempty"`

	// URI that provides the ability to access an image of the assembly information.
	BinaryDataURI string `json:"BinaryDataURI,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Engineering change level of the Assembly.
	EngineeringChangeLevel string `json:"EngineeringChangeLevel,omitempty"`

	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// Model number of the Assembly.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Part number of the Assembly.
	PartNumber string `json:"PartNumber,omitempty"`

	// Producer or manufacturer of the Assembly.
	Producer string `json:"Producer,omitempty"`

	// Production date of the Assembly.
	ProductionDate string `json:"ProductionDate,omitempty"`

	// SKU of the Assembly.
	SKU string `json:"SKU,omitempty"`

	// Spare part number of the Assembly.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// Vendor of the Assembly.
	Vendor string `json:"Vendor,omitempty"`

	// Version of the Assembly.
	Version string `json:"Version,omitempty"`

}
