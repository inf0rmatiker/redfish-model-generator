/* -----------------------------------------------------------------
* discrete_trigger.go -
*
* DMTF Redfish DiscreteTrigger resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DiscreteTrigger - The characteristics of the discrete trigger.
// Modeled after DMTF Redfish schema DiscreteTrigger
type DiscreteTrigger struct {
	// This time the trigger occurance persists before a trigger event has occurred.
	DwellTime string `json:"DwellTime,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// This property contains the value of the Severity property in the Event message.
	Severity map[string]interface{} `json:"Severity,omitempty"`

	// The value of the discrete metric that constitutes a trigger occurance.
	Value string `json:"Value,omitempty"`

}
