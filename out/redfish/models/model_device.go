/* -----------------------------------------------------------------
* device.go -
*
* DMTF Redfish Device resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Device - A storage device such as a disk drive or optical media device.
// Modeled after DMTF Redfish schema Device
type Device struct {
	// The size of the storage device.
	CapacityBytes int `json:"CapacityBytes,omitempty"`

	// The name of the manufacturer of this device.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The product model number of this device.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

}
