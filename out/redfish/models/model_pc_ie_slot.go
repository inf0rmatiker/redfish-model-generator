/* -----------------------------------------------------------------
* pc_ie_slot.go -
*
* DMTF Redfish PCIeSlot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PCIeSlot - This type defines information for a PCIe slot.
// Modeled after DMTF Redfish schema PCIeSlot
type PCIeSlot struct {
	// An indication of whether this PCIe slot supports hotplug.
	HotPluggable bool `json:"HotPluggable,omitempty"`

	// The number of PCIe lanes supported by this slot.
	Lanes int `json:"Lanes,omitempty"`

	// The links to other Resources that are related to this Resource.
	Links PCIeLinks `json:"Links,omitempty"`

	// The location of the PCIe slot.
	Location map[string]interface{} `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The PCIe specification supported by this slot.
	PCIeType map[string]interface{} `json:"PCIeType,omitempty"`

	// The PCIe slot type for this slot.
	SlotType SlotTypes `json:"SlotType,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
