/* -----------------------------------------------------------------
* secure_boot_database.go -
*
* DMTF Redfish SecureBootDatabase resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SecureBootDatabase - The SecureBootDatabase schema describes a UEFI Secure Boot database used to store certificates or hashes.
// Modeled after DMTF Redfish schema SecureBootDatabase
type SecureBootDatabase struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A link to the collection of certificates contained in this UEFI Secure Boot database.
	Certificates map[string]interface{} `json:"Certificates,omitempty"`

	// This property contains the name of the UEFI Secure Boot database.
	DatabaseId string `json:"DatabaseId,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A link to the collection of signatures contained in this UEFI Secure Boot database.
	Signatures map[string]interface{} `json:"Signatures,omitempty"`

}
