/* -----------------------------------------------------------------
* power_distribution.go -
*
* DMTF Redfish PowerDistribution resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerDistribution - This is the schema definition for a power distribution component or unit, such as a floor power distribution unit (PDU) or switchgear.
// Modeled after DMTF Redfish schema PowerDistribution
type PowerDistribution struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The user-assigned asset tag for this equipment.
	AssetTag string `json:"AssetTag,omitempty"`

	// A link to the branch circuits for this equipment.
	Branches map[string]interface{} `json:"Branches,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The type of equipment this resource represents.
	EquipmentType PowerEquipmentType `json:"EquipmentType"`

	// A link to the feeder circuits for this equipment.
	Feeders map[string]interface{} `json:"Feeders,omitempty"`

	// The firmware version of this equipment.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the equipment.
	Location map[string]interface{} `json:"Location,omitempty"`

	// A link to the power input circuits for this equipment.
	Mains map[string]interface{} `json:"Mains,omitempty"`

	// The manufacturer of this equipment.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// A link to the summary metrics for this equipment.
	Metrics map[string]interface{} `json:"Metrics,omitempty"`

	// The product model number of this equipment.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A link to the outlet groups for this equipment.
	OutletGroups map[string]interface{} `json:"OutletGroups,omitempty"`

	// A link to the outlets for this equipment.
	Outlets map[string]interface{} `json:"Outlets,omitempty"`

	// The part number for this equipment.
	PartNumber string `json:"PartNumber,omitempty"`

	// The production or manufacturing date of this equipment.
	ProductionDate string `json:"ProductionDate,omitempty"`

	// A link to the collection of sensors located in the equipment and sub-components.
	Sensors map[string]interface{} `json:"Sensors,omitempty"`

	// The serial number for this equipment.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// A link to the subfeed circuits for this equipment.
	Subfeeds map[string]interface{} `json:"Subfeeds,omitempty"`

	// The configuration settings for an automatic transfer switch.
	TransferConfiguration TransferConfiguration `json:"TransferConfiguration,omitempty"`

	// The criteria used to initiate a transfer for an automatic transfer switch.
	TransferCriteria TransferCriteria `json:"TransferCriteria,omitempty"`

	// The UUID for this equipment.
	UUID map[string]interface{} `json:"UUID,omitempty"`

	// The hardware version of this equipment.
	Version string `json:"Version,omitempty"`

}
