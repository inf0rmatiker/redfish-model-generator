/* -----------------------------------------------------------------
* message_registry.go -
*
* DMTF Redfish MessageRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The MessageRegistry schema describes all message registries.  It represents the properties for the message registries themselves.
type MessageRegistry struct {
	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The RFC5646-conformant language code for the message registry.
	Language string `json:"Language"`

	// The message keys contained in the message registry.
	Messages MessageProperty `json:"Messages"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The organization or company that publishes this message registry.
	OwningEntity string `json:"OwningEntity"`

	// The single-word prefix that is used in forming and decoding MessageIds.
	RegistryPrefix string `json:"RegistryPrefix"`

	// The message registry version in the middle portion of a MessageId.
	RegistryVersion string `json:"RegistryVersion"`

}
