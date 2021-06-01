/* -----------------------------------------------------------------
* boot_option.go -
*
* DMTF Redfish BootOption resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// BootOption - This is the schema definition for the BootOption resource. It represents the properties of a bootable device available in the System.
// Modeled after DMTF Redfish schema BootOption
type BootOption struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The alias of this Boot Source when described in the BootSourceOverrideTarget property in the Computersystem resource.
	Alias map[string]interface{} `json:"Alias,omitempty"`

	// A flag that shows if the Boot Option is enabled.
	BootOptionEnabled bool `json:"BootOptionEnabled,omitempty"`

	// The unique boot option string that is referenced in the BootOrder.
	BootOptionReference string `json:"BootOptionReference"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The user-readable display string of the Boot Option.
	DisplayName string `json:"DisplayName,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The ID(s) of the resources associated with this Boot Option.
	RelatedItem []idRef `json:"RelatedItem,omitempty"`

	RelatedItemOdataCount int `json:"RelatedItem@odata.count,omitempty"`

	// The UEFI device path used to access this UEFI Boot Option.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`

}
