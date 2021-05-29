/* -----------------------------------------------------------------
* fan_actions.go -
*
* DMTF Redfish FanActions resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The available actions for this resource.
type FanActions struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
