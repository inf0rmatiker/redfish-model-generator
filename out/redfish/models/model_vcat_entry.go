/* -----------------------------------------------------------------
* vcat_entry.go -
*
* DMTF Redfish VCATEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VCATEntry - The VCATEntry schema defines an entry in a Virtual Channel Action Table.  A Virtual Channel is a mechanism used to create multiple, logical communication streams across a physical link.
// Modeled after DMTF Redfish schema VCATEntry
type VCATEntry struct {
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

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The hexadecimal value of the Virtual Channel Action Table entries.
	RawEntryHex string `json:"RawEntryHex,omitempty"`

	// An array of entries of the Virtual Channel Action Table.
	VCEntries []VCATableEntry `json:"VCEntries,omitempty"`

}
