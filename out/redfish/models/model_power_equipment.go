/* -----------------------------------------------------------------
* power_equipment.go -
*
* DMTF Redfish PowerEquipment resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerEquipment - This is the schema definition for the set of power equipment.
// Modeled after DMTF Redfish schema PowerEquipment
type PowerEquipment struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// A link to a collection of floor power distribution units.
	FloorPDUs map[string]interface{} `json:"FloorPDUs,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A link to a collection of rack-level power distribution units.
	RackPDUs map[string]interface{} `json:"RackPDUs,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// A link to a collection of switchgear.
	Switchgear map[string]interface{} `json:"Switchgear,omitempty"`

	// A link to a collection of transfer switches.
	TransferSwitches map[string]interface{} `json:"TransferSwitches,omitempty"`

}
