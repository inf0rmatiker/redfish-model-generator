/* -----------------------------------------------------------------
* message_registry_file.go -
*
* DMTF Redfish MessageRegistryFile resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MessageRegistryFile - This is the schema definition for the Schema File locator resource.
// Modeled after DMTF Redfish schema MessageRegistryFile
type MessageRegistryFile struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Language codes for the schemas available.
	Languages []s `json:"Languages"`

	// Location information for this registry file.
	Location []Location `json:"Location"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Registry Name, Major, and Minor version.  This Registry can reference any type of Registry, such as a Message Registry, Privilege Registry, or Attribute Registry.
	Registry string `json:"Registry"`

}
