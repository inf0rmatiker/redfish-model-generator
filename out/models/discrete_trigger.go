/* -----------------------------------------------------------------
* discrete_trigger.go -
*
* DMTF Redfish DiscreteTrigger resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The characteristics of the discrete trigger.
type DiscreteTrigger struct {
	// The amount of time that a trigger event persists before the metric action is performed.
	DwellTime string `json:"DwellTime,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The severity of the event message.
	Severity Health `json:"Severity,omitempty"`

	// The discrete metric value that constitutes a trigger event.
	Value string `json:"Value,omitempty"`

}
