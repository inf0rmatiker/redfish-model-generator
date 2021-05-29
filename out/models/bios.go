/* -----------------------------------------------------------------
* bios.go -
*
* DMTF Redfish Bios resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Bios schema contains properties related to the BIOS attribute registry.  The attribute registry describes the system-specific BIOS attributes and actions for changing to BIOS settings.  Changes to the BIOS typically require a system reset before they take effect.  It is likely that a client finds the `@Redfish.Settings` term in this resource, and if it is found, the client makes requests to change BIOS settings by modifying the resource identified by the `@Redfish.Settings` term.
type Bios struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The resource ID of the attribute registry that has the system-specific information about a BIOS resource.
	AttributeRegistry string `json:"AttributeRegistry,omitempty"`

	// The list of BIOS attributes specific to the manufacturer or provider.
	Attributes Attributes `json:"Attributes,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether there is a pending request to reset the BIOS attributes to default values.
	ResetBiosToDefaultsPending bool `json:"ResetBiosToDefaultsPending,omitempty"`

}
