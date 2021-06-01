/* -----------------------------------------------------------------
* secure_boot.go -
*
* DMTF Redfish SecureBoot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SecureBoot - This resource contains UEFI Secure Boot information. It represents properties for managing the UEFI Secure Boot functionality of a system.
// Modeled after DMTF Redfish schema SecureBoot
type SecureBoot struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// Enable or disable UEFI Secure Boot (takes effect on next boot).
	SecureBootEnable bool `json:"SecureBootEnable,omitempty"`

	// Secure Boot state during the current boot cycle.
	SecureBootCurrentBoot SecureBootCurrentBootType `json:"SecureBootCurrentBoot,omitempty"`

	// Current Secure Boot Mode.
	SecureBootMode SecureBootModeType `json:"SecureBootMode,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

}
