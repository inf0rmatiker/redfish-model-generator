/* -----------------------------------------------------------------
* bios.go -
*
* DMTF Redfish Bios resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Bios - The Bios schema contains properties related to the BIOS Attribute Registry.  The Attribute Registry describes the system-specific BIOS attributes and Actions for changing to BIOS settings.  Changes to the BIOS typically require a system reset before they take effect.  It is likely that a client will find the @Redfish.Settings term in this resource, and if it is found, the client makes requests to change BIOS settings by modifying the resource identified by the @Redfish.Settings term.
// Modeled after DMTF Redfish schema Bios
type Bios struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The Resource ID of the Attribute Registry that has the system-specific information about a BIOS resource.
	AttributeRegistry string `json:"AttributeRegistry,omitempty"`

	// The list of BIOS attributes specific to the manufacturer or provider.
	Attributes map[string]interface{} `json:"Attributes,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
