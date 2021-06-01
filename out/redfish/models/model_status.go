/* -----------------------------------------------------------------
* status.go -
*
* DMTF Redfish Status resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Status - The status and health of a resource and its children.
// Modeled after DMTF Redfish schema Status
type Status struct {
	// Conditions in this resource that require attention.
	Conditions []Condition `json:"Conditions,omitempty"`

	// The health state of this resource in the absence of its dependent resources.
	Health Health `json:"Health,omitempty"`

	// The overall health state from the view of this resource.
	HealthRollup Health `json:"HealthRollup,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The known state of the resource, such as, enabled.
	State State `json:"State,omitempty"`

}
