/* -----------------------------------------------------------------
* secure_boot.go -
*
* DMTF Redfish SecureBoot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The SecureBoot schema contains UEFI Secure Boot information and represents properties for managing the UEFI Secure Boot functionality of a system.
type SecureBoot struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The UEFI Secure Boot state during the current boot cycle.
	SecureBootCurrentBoot SecureBootCurrentBootType `json:"SecureBootCurrentBoot,omitempty"`

	// A link to the collection of UEFI Secure Boot databases.
	SecureBootDatabases SecureBootDatabaseCollection `json:"SecureBootDatabases,omitempty"`

	// An indication of whether UEFI Secure Boot is enabled.
	SecureBootEnable bool `json:"SecureBootEnable,omitempty"`

	// The current UEFI Secure Boot Mode.
	SecureBootMode SecureBootModeType `json:"SecureBootMode,omitempty"`

}
