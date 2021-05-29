/* -----------------------------------------------------------------
* assembly_data.go -
*
* DMTF Redfish AssemblyData resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AssemblyData struct {
	OdataId string `json:"@odata.id"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The URI at which to access an image of the assembly information.
	BinaryDataURI string `json:"BinaryDataURI,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The engineering change level of the assembly.
	EngineeringChangeLevel string `json:"EngineeringChangeLevel,omitempty"`

	// The location of the assembly.
	Location Location `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The model number of the assembly.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number of the assembly.
	PartNumber string `json:"PartNumber,omitempty"`

	// The area or device to which the assembly data applies.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// The producer or manufacturer of the assembly.
	Producer string `json:"Producer,omitempty"`

	// The production date of the assembly.
	ProductionDate string `json:"ProductionDate,omitempty"`

	// The SKU of the assembly.
	SKU string `json:"SKU,omitempty"`

	// The serial number of the assembly.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number of the assembly.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

	// The vendor of the assembly.
	Vendor string `json:"Vendor,omitempty"`

	// The hardware version of the assembly.
	Version string `json:"Version,omitempty"`

}
