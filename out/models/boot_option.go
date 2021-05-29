/* -----------------------------------------------------------------
* boot_option.go -
*
* DMTF Redfish BootOption resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The BootOption schema reports information about a single boot option in a system.  It represents the properties of a bootable device available in the system.
type BootOption struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The alias of this boot source.
	Alias BootSource `json:"Alias,omitempty"`

	// An indication of whether the boot option is enabled.  If `true`, it is enabled.  If `false`, the boot option that the boot order array on the computer system contains is skipped.  In the UEFI context, this property shall influence the load option active flag for the boot option.
	BootOptionEnabled bool `json:"BootOptionEnabled,omitempty"`

	// The unique boot option.
	BootOptionReference string `json:"BootOptionReference"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The user-readable display name of the boot option that appears in the boot order list in the user interface.
	DisplayName string `json:"DisplayName,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to resources or objects associated with this boot option.
	RelatedItem []map[string]string `json:"RelatedItem,omitempty"`

	RelatedItem@odata.count count `json:"RelatedItem@odata.count,omitempty"`

	// The UEFI device path to access this UEFI boot option.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`

}
