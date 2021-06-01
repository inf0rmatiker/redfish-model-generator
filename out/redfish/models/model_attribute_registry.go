/* -----------------------------------------------------------------
* attribute_registry.go -
*
* DMTF Redfish AttributeRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// AttributeRegistry - The AttributeRegistry schema contains a set of key-value pairs that represents the structure of a Registry. It includes mechanisms for building user interfaces (menus), allowing consistent navigation of the contents. The Attribute Registry is specific to a particular implementation or product. The attributes and property names are not standardized.
// Modeled after DMTF Redfish schema AttributeRegistry
type AttributeRegistry struct {
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

	// This is the RFC 5646 compliant language code for the registry.
	Language string `json:"Language"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the organization or company that publishes this registry.
	OwningEntity string `json:"OwningEntity"`

	// List of all attributes and their metadata for this component.
	RegistryEntries RegistryEntries `json:"RegistryEntries,omitempty"`

	// This is the attribute registry version which is used in the middle portion of a AttributeRegistry.
	RegistryVersion string `json:"RegistryVersion"`

	// Array of systems supported by this attribute registry.
	SupportedSystems []SupportedSystems `json:"SupportedSystems,omitempty"`

}
