/* -----------------------------------------------------------------
* data_security_lo_s_capabilities.go -
*
* SNIA Swordfish DataSecurityLoSCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataSecurityLoSCapabilities - Describe data security capabilities.
// Modeled after SNIA Swordfish schema DataSecurityLoSCapabilities
type DataSecurityLoSCapabilities struct {
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

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Supported AntiVirus providers.
	SupportedAntivirusEngineProviders []string `json:"SupportedAntivirusEngineProviders,omitempty"`

	// Supported policies that trigger an AntiVirus scan.
	SupportedAntivirusScanPolicies []AntiVirusScanTrigger `json:"SupportedAntivirusScanPolicies,omitempty"`

	// Supported key sizes for transport channel encryption.
	SupportedChannelEncryptionStrengths []KeySize `json:"SupportedChannelEncryptionStrengths,omitempty"`

	// Supported data sanitization policies.
	SupportedDataSanitizationPolicies []DataSanitizationPolicy `json:"SupportedDataSanitizationPolicies,omitempty"`

	// Supported authentication types for hosts (servers) or initiator endpoints.
	SupportedHostAuthenticationTypes []AuthenticationType `json:"SupportedHostAuthenticationTypes,omitempty"`

	// Collection of known and supported DataSecurityLinesOfService.
	SupportedLinesOfService []DataSecurityLineOfService `json:"SupportedLinesOfService,omitempty"`

	SupportedLinesOfServiceOdataCount int `json:"SupportedLinesOfService@odata.count,omitempty"`

	// Supported key sizes for media encryption.
	SupportedMediaEncryptionStrengths []KeySize `json:"SupportedMediaEncryptionStrengths,omitempty"`

	// Supported protocols that provide encrypted communication.
	SupportedSecureChannelProtocols []SecureChannelProtocol `json:"SupportedSecureChannelProtocols,omitempty"`

	// Supported authentication types for users (or programs).
	SupportedUserAuthenticationTypes []AuthenticationType `json:"SupportedUserAuthenticationTypes,omitempty"`

}
