/* -----------------------------------------------------------------
* drive.go -
*
* DMTF Redfish Drive resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Drive - The Drive schema represents a single physical drive for a system, including links to associated volumes.
// Modeled after DMTF Redfish schema Drive
type Drive struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly associated with this drive.
	Assembly map[string]interface{} `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this drive.
	AssetTag string `json:"AssetTag,omitempty"`

	// The size, in bytes, of the smallest addressable unit, or block.
	BlockSizeBytes int `json:"BlockSizeBytes,omitempty"`

	// The speed, in gigabit per second (Gbit/s), at which this drive can communicate to a storage controller in ideal conditions.
	CapableSpeedGbs float64 `json:"CapableSpeedGbs,omitempty"`

	// The size, in bytes, of this drive.
	CapacityBytes int `json:"CapacityBytes,omitempty"`

	// The link to a collection of certificates for device identity and attestation.
	Certificates map[string]interface{} `json:"Certificates,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The encryption ability of this drive.
	EncryptionAbility EncryptionAbility `json:"EncryptionAbility,omitempty"`

	// The status of the encryption of this drive.
	EncryptionStatus EncryptionStatus `json:"EncryptionStatus,omitempty"`

	// The link to the environment metrics for this drive.
	EnvironmentMetrics map[string]interface{} `json:"EnvironmentMetrics,omitempty"`

	// An indication of whether this drive currently predicts a failure in the near future.
	FailurePredicted bool `json:"FailurePredicted,omitempty"`

	// The replacement mode for the hot spare drive.
	HotspareReplacementMode HotspareReplacementModeType `json:"HotspareReplacementMode,omitempty"`

	// The type of hot spare that this drive currently serves as.
	HotspareType HotspareType `json:"HotspareType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable names for the drive.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// The state of the indicator LED, that identifies the drive.
	IndicatorLED map[string]interface{} `json:"IndicatorLED,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the drive.
	Location []Location `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The manufacturer of this drive.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	Measurements []MeasurementBlock `json:"Measurements,omitempty"`

	// The type of media contained in this drive.
	MediaType MediaType `json:"MediaType,omitempty"`

	// The model number for the drive.
	Model string `json:"Model,omitempty"`

	// An indication of whether the drive is accessible from multiple paths.
	Multipath bool `json:"Multipath,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The speed, in gigabit per second (Gbit/s), at which this drive currently communicates to the storage controller.
	NegotiatedSpeedGbs float64 `json:"NegotiatedSpeedGbs,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The operations currently running on the Drive.
	Operations []Operations `json:"Operations,omitempty"`

	// The part number for this drive.
	PartNumber string `json:"PartNumber,omitempty"`

	// The location of the drive.
	PhysicalLocation map[string]interface{} `json:"PhysicalLocation,omitempty"`

	// The percentage of reads and writes that are predicted to be available for the media.
	PredictedMediaLifeLeftPercent float64 `json:"PredictedMediaLifeLeftPercent,omitempty"`

	// The protocol that this drive currently uses to communicate to the storage controller.
	Protocol map[string]interface{} `json:"Protocol,omitempty"`

	// An indication of whether the drive is prepared by the system for removal.
	ReadyToRemove bool `json:"ReadyToRemove,omitempty"`

	// The revision of this drive.  This is typically the firmware or hardware version of the drive.
	Revision string `json:"Revision,omitempty"`

	// The rotation speed of this drive, in revolutions per minute (RPM).
	RotationSpeedRPM float64 `json:"RotationSpeedRPM,omitempty"`

	// The SKU for this drive.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this drive.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The state of the status indicator, which communicates status information about this drive.
	StatusIndicator StatusIndicator `json:"StatusIndicator,omitempty"`

	// An indication of whether the drive write cache is enabled.
	WriteCacheEnabled bool `json:"WriteCacheEnabled,omitempty"`

}
