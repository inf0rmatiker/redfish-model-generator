/* -----------------------------------------------------------------
* operating_config.go -
*
* DMTF Redfish OperatingConfig resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// OperatingConfig - The OperatingConfig schema specifies a configuration that can be used when the processor is operational.
// Modeled after DMTF Redfish schema OperatingConfig
type OperatingConfig struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The base (nominal) clock speed of the processor in MHz.
	BaseSpeedMHz int `json:"BaseSpeedMHz,omitempty"`

	// The clock speed for sets of cores when the configuration is operational.
	BaseSpeedPrioritySettings []BaseSpeedPrioritySettings `json:"BaseSpeedPrioritySettings,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The maximum temperature of the junction in degrees Celsius.
	MaxJunctionTemperatureCelsius int `json:"MaxJunctionTemperatureCelsius,omitempty"`

	// The maximum clock speed to which the processor can be configured in MHz.
	MaxSpeedMHz int `json:"MaxSpeedMHz,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The thermal design point of the processor in watts.
	TDPWatts int `json:"TDPWatts,omitempty"`

	// The number of cores in the processor that can be configured.
	TotalAvailableCoreCount int `json:"TotalAvailableCoreCount,omitempty"`

	// The turbo profiles for the processor.  A turbo profile is the maximum turbo clock speed as a function of the number of active cores.
	TurboProfile []TurboProfileDatapoint `json:"TurboProfile,omitempty"`

}
