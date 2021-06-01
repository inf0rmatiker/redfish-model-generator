/* -----------------------------------------------------------------
* gen_z.go -
*
* DMTF Redfish GenZ resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// GenZ - This type defines Gen-Z specific port properties.
// Modeled after DMTF Redfish schema GenZ
type GenZ struct {
	// The Linear Packet Relay Table for the port.
	LPRT map[string]interface{} `json:"LPRT,omitempty"`

	// the Multi-subnet Packet Relay Table for the port.
	MPRT map[string]interface{} `json:"MPRT,omitempty"`

	// the Virtual Channel Action Table for the port.
	VCAT map[string]interface{} `json:"VCAT,omitempty"`

}
