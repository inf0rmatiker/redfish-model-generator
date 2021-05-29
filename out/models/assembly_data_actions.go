/* -----------------------------------------------------------------
* assembly_data_actions.go -
*
* DMTF Redfish AssemblyDataActions resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The available actions for this Resource.
type AssemblyDataActions struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
