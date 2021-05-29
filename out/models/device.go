/* -----------------------------------------------------------------
* device.go -
*
* DMTF Redfish Device resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A storage device, such as a disk drive or optical media device.
type Device struct {
	// The size, in bytes, of the storage device.
	CapacityBytes int `json:"CapacityBytes,omitempty"`

	// The name of the manufacturer of this device.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The product model number of this device.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

}
