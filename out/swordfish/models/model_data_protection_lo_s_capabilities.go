/* -----------------------------------------------------------------
* data_protection_lo_s_capabilities.go -
*
* SNIA Swordfish DataProtectionLoSCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataProtectionLoSCapabilities - The capabilities to protect data from loss by the use of a replica.
// Modeled after SNIA Swordfish schema DataProtectionLoSCapabilities
type DataProtectionLoSCapabilities struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// Contains links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Collection of known and supported DataProtectionLinesOfService.
	SupportedLinesOfService []DataProtectionLineOfService `json:"SupportedLinesOfService,omitempty"`

	SupportedLinesOfServiceOdataCount int `json:"SupportedLinesOfService@odata.count,omitempty"`

	// Supported minimum lifetime that replica must be maintained.
	SupportedMinLifetimes []string `json:"SupportedMinLifetimes,omitempty"`

	// Supported types of failure domains.
	SupportedRecoveryGeographicObjectives []FailureDomainScope `json:"SupportedRecoveryGeographicObjectives,omitempty"`

	// Supported time intervals defining how much source information can be lost on failure.
	SupportedRecoveryPointObjectiveTimes []string `json:"SupportedRecoveryPointObjectiveTimes,omitempty"`

	// Supported expectations for time to access an alternate replica.
	SupportedRecoveryTimeObjectives []RecoveryAccessScope `json:"SupportedRecoveryTimeObjectives,omitempty"`

	// Supported replica types.
	SupportedReplicaTypes []ReplicaType `json:"SupportedReplicaTypes,omitempty"`

	// Allocating a replica in a separate fault domain is supported.
	SupportsIsolated bool `json:"SupportsIsolated,omitempty"`

}
