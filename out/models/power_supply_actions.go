/* -----------------------------------------------------------------
* power_supply_actions.go -
*
* DMTF Redfish PowerSupplyActions resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The available actions for this resource.
type PowerSupplyActions struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
