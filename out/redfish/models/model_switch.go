/* -----------------------------------------------------------------
* switch.go -
*
* DMTF Redfish Switch resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Switch - The Switch schema contains properties that describe a fabric switch.
// Modeled after DMTF Redfish schema Switch
type Switch struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The user-assigned asset tag for this switch.
	AssetTag string `json:"AssetTag,omitempty"`

	// The current internal bandwidth of this switch.
	CurrentBandwidthGbps float64 `json:"CurrentBandwidthGbps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The domain ID for this switch.
	DomainID int `json:"DomainID,omitempty"`

	// The firmware version of this switch.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the indicator LED, which identifies the switch.
	IndicatorLED map[string]interface{} `json:"IndicatorLED,omitempty"`

	// An indication of whether the switch is in a managed or unmanaged state.
	IsManaged bool `json:"IsManaged,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the switch.
	Location map[string]interface{} `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The link to the collection of log services associated with this switch.
	LogServices map[string]interface{} `json:"LogServices,omitempty"`

	// The manufacturer of this switch.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The maximum internal bandwidth of this switch as currently configured.
	MaxBandwidthGbps float64 `json:"MaxBandwidthGbps,omitempty"`

	// The product model number of this switch.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this switch.
	PartNumber string `json:"PartNumber,omitempty"`

	// The link to the collection ports for this switch.
	Ports map[string]interface{} `json:"Ports,omitempty"`

	// The current power state of the switch.
	PowerState map[string]interface{} `json:"PowerState,omitempty"`

	// Redundancy information for the switches.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	// The SKU for this switch.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this switch.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The protocols this switch supports.
	SupportedProtocols []Protocol `json:"SupportedProtocols,omitempty"`

	// The protocol being sent over this switch.
	SwitchType map[string]interface{} `json:"SwitchType,omitempty"`

	// The total number of lanes, phys, or other physical transport links that this switch contains.
	TotalSwitchWidth int `json:"TotalSwitchWidth,omitempty"`

	// The UUID for this switch.
	UUID map[string]interface{} `json:"UUID,omitempty"`

}
