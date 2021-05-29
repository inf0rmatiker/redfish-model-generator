/* -----------------------------------------------------------------
* message_registry_file.go -
*
* DMTF Redfish MessageRegistryFile resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The MessageRegistryFile schema describes the Message Registry file locator Resource.
type MessageRegistryFile struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The RFC5646-conformant language codes for the available Message Registries.
	Languages []string `json:"Languages"`

	// The location information for this Message Registry file.
	Location array `json:"Location"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The registry name and its major and minor versions.  This registry can be any type of registry, such as a Message Registry, Privilege Registry, or Attribute Registry.
	Registry string `json:"Registry"`

}
