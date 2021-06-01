/* -----------------------------------------------------------------
* triggers.go -
*
* DMTF Redfish Triggers resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Triggers - Triggers which apply to a list of metrics.
// Modeled after DMTF Redfish schema Triggers
type Triggers struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Specifies the conditions when a discrete metric triggers.
	DiscreteTriggerCondition DiscreteTriggerConditionEnum `json:"DiscreteTriggerCondition,omitempty"`

	// List of discrete triggers.
	DiscreteTriggers []DiscreteTrigger `json:"DiscreteTriggers,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// A list of URIs with wildcards and property identifiers for which this trigger is defined. If a URI has wildcards, the wildcards are substituted as specified in the Wildcards array property.
	MetricProperties []string `json:"MetricProperties,omitempty"`

	// The type of trigger.
	MetricType MetricTypeEnum `json:"MetricType,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// Specifies the thresholds when a numeric metric triggers.
	NumericThresholds Thresholds `json:"NumericThresholds,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// This property specifies the actions to perform when the trigger occurs.
	TriggerActions []TriggerActionEnum `json:"TriggerActions,omitempty"`

	// A list of wildcards and their substitution values to be applied to the entries in the MetricProperties array property.
	Wildcards []Wildcard `json:"Wildcards,omitempty"`

}
