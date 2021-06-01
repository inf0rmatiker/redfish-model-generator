/* -----------------------------------------------------------------
* message_registry.go -
*
* DMTF Redfish MessageRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MessageRegistry - The MessageRegistry schema describes all Message Registries.  It represents the properties for the Message Registries themselves.
// Modeled after DMTF Redfish schema MessageRegistry
type MessageRegistry struct {
	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The RFC5646-conformant language code for the Message Registry.
	Language string `json:"Language"`

	// The message keys contained in the Message Registry.
	Messages map[string]interface{} `json:"Messages"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The organization or company that publishes this Message Registry.
	OwningEntity string `json:"OwningEntity"`

	// The single-word prefix that is used in forming and decoding MessageIds.
	RegistryPrefix string `json:"RegistryPrefix"`

	// The Message Registry version in the middle portion of a MessageId.
	RegistryVersion string `json:"RegistryVersion"`

}
