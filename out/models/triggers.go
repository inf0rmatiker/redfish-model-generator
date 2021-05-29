/* -----------------------------------------------------------------
* triggers.go -
*
* DMTF Redfish Triggers resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Triggers schema describes a trigger that applies to metrics.
type Triggers struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The conditions when a discrete metric triggers.
	DiscreteTriggerCondition DiscreteTriggerConditionEnum `json:"DiscreteTriggerCondition,omitempty"`

	// The list of discrete triggers.
	DiscreteTriggers array `json:"DiscreteTriggers,omitempty"`

	// The array of MessageIds that specify when a trigger condition is met based on an event.
	EventTriggers []string `json:"EventTriggers,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An array of URIs with wildcards and property identifiers for this trigger.  Each wildcard shall be replaced with its corresponding entry in the Wildcard array property.
	MetricProperties []string `json:"MetricProperties,omitempty"`

	// The metric type of the trigger.
	MetricType MetricTypeEnum `json:"MetricType,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The thresholds when a numeric metric triggers.
	NumericThresholds Thresholds `json:"NumericThresholds,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The actions that the trigger initiates.
	TriggerActions array `json:"TriggerActions,omitempty"`

	// The wildcards and their substitution values for the entries in the MetricProperties array property.
	Wildcards array `json:"Wildcards,omitempty"`

}
