/* -----------------------------------------------------------------
* attribute_registry.go -
*
* DMTF Redfish AttributeRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The AttributeRegistry schema contains a set of key-value pairs that represent the structure of an attribute registry.  It includes mechanisms for building user interfaces, or menus, allowing consistent navigation of the contents.  The attribute registry is specific to an implementation or product.  The attributes and property names are not standardized.
type AttributeRegistry struct {
	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The RFC5646-conformant language code for the attribute registry.
	Language string `json:"Language"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The organization or company that publishes this attribute registry.
	OwningEntity string `json:"OwningEntity"`

	// The list of all attributes and their metadata for this component.
	RegistryEntries RegistryEntries `json:"RegistryEntries,omitempty"`

	// The attribute registry version.
	RegistryVersion string `json:"RegistryVersion"`

	// An array of systems that this attribute registry supports.
	SupportedSystems array `json:"SupportedSystems,omitempty"`

}
