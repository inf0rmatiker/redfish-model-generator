/* -----------------------------------------------------------------
* data_security_line_of_service.go -
*
* SNIA Swordfish DataSecurityLineOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataSecurityLineOfService - Describes data security service level requirements.
// Modeled after SNIA Swordfish schema DataSecurityLineOfService
type DataSecurityLineOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// AntiVirus provider.
	AntivirusEngineProvider string `json:"AntivirusEngineProvider,omitempty"`

	// Policy for triggering an AntiVirus scan.
	AntivirusScanPolicies []AntiVirusScanTrigger `json:"AntivirusScanPolicies,omitempty"`

	// Key size for transport channel encryption.
	ChannelEncryptionStrength map[string]interface{} `json:"ChannelEncryptionStrength,omitempty"`

	// Data sanitization policy.
	DataSanitizationPolicy map[string]interface{} `json:"DataSanitizationPolicy,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Authentication type for hosts (servers) or initiator endpoints.
	HostAuthenticationType map[string]interface{} `json:"HostAuthenticationType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Key size for media encryption.
	MediaEncryptionStrength map[string]interface{} `json:"MediaEncryptionStrength,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Protocol that provide encrypted communication.
	SecureChannelProtocol map[string]interface{} `json:"SecureChannelProtocol,omitempty"`

	// Authentication type for users (or programs).
	UserAuthenticationType map[string]interface{} `json:"UserAuthenticationType,omitempty"`

}
