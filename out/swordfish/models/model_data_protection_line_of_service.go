/* -----------------------------------------------------------------
* data_protection_line_of_service.go -
*
* SNIA Swordfish DataProtectionLineOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataProtectionLineOfService - Describes a data protection service option.
// Modeled after SNIA Swordfish schema DataProtectionLineOfService
type DataProtectionLineOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The replica is in a separate fault domain.
	IsIsolated bool `json:"IsIsolated,omitempty"`

	// Minimum lifetime (seconds) that replica must be maintained.
	MinLifetime string `json:"MinLifetime,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Geographic distribution scopes.
	RecoveryGeographicObjective map[string]interface{} `json:"RecoveryGeographicObjective,omitempty"`

	// Time interval defining how much source data that can be lost on failure.
	RecoveryPointObjectiveTime string `json:"RecoveryPointObjectiveTime,omitempty"`

	// An enumeration value that indicates the expected time to access an alternate replica.
	RecoveryTimeObjective map[string]interface{} `json:"RecoveryTimeObjective,omitempty"`

	// Location that supplies data access to the replica.
	ReplicaAccessLocation map[string]interface{} `json:"ReplicaAccessLocation,omitempty"`

	// The replica's class of service.
	ReplicaClassOfService map[string]interface{} `json:"ReplicaClassOfService,omitempty"`

	// Type of replica.
	ReplicaType map[string]interface{} `json:"ReplicaType,omitempty"`

	// A schedule for making periodic point in time replicas.
	Schedule map[string]interface{} `json:"Schedule,omitempty"`

}
