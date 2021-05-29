/* -----------------------------------------------------------------
* switch.go -
*
* DMTF Redfish Switch resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Switch schema contains properties that describe a fabric switch.
type Switch struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The user-assigned asset tag for this switch.
	AssetTag string `json:"AssetTag,omitempty"`

	// The link to a collection of certificates for device identity and attestation.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The current internal bandwidth of this switch.
	CurrentBandwidthGbps float64 `json:"CurrentBandwidthGbps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The domain ID for this switch.
	DomainID int `json:"DomainID,omitempty"`

	// An indication of whether this switch is enabled.
	Enabled bool `json:"Enabled,omitempty"`

	// The link to the environment metrics for this switch.
	EnvironmentMetrics EnvironmentMetrics `json:"EnvironmentMetrics,omitempty"`

	// The firmware version of this switch.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the indicator LED, which identifies the switch.
	IndicatorLED IndicatorLED `json:"IndicatorLED,omitempty"`

	// An indication of whether the switch is in a managed or unmanaged state.
	IsManaged bool `json:"IsManaged,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the switch.
	Location Location `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The link to the collection of log services associated with this switch.
	LogServices LogServiceCollection `json:"LogServices,omitempty"`

	// The manufacturer of this switch.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The maximum internal bandwidth of this switch as currently configured.
	MaxBandwidthGbps float64 `json:"MaxBandwidthGbps,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	Measurements array `json:"Measurements,omitempty"`

	// The product model number of this switch.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this switch.
	PartNumber string `json:"PartNumber,omitempty"`

	// The link to the collection ports for this switch.
	Ports PortCollection `json:"Ports,omitempty"`

	// The current power state of the switch.
	PowerState PowerState `json:"PowerState,omitempty"`

	// Redundancy information for the switches.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The SKU for this switch.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this switch.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The protocols this switch supports.
	SupportedProtocols array `json:"SupportedProtocols,omitempty"`

	// The protocol being sent over this switch.
	SwitchType Protocol `json:"SwitchType,omitempty"`

	// The total number of lanes, phys, or other physical transport links that this switch contains.
	TotalSwitchWidth int `json:"TotalSwitchWidth,omitempty"`

	// The UUID for this switch.
	UUID UUID `json:"UUID,omitempty"`

}
