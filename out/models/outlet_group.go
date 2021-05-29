/* -----------------------------------------------------------------
* outlet_group.go -
*
* DMTF Redfish OutletGroup resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The OutletGroup schema contains definitions for an electrical outlet group.
type OutletGroup struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The creator of this outlet group.
	CreatedBy string `json:"CreatedBy,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The energy reading for this outlet group.
	EnergykWh SensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The number of seconds to delay power on after a PowerControl action to cycle power.  Zero seconds indicates no delay.
	PowerCycleDelaySeconds float64 `json:"PowerCycleDelaySeconds,omitempty"`

	// Indicates if the outlet group can be powered.
	PowerEnabled bool `json:"PowerEnabled,omitempty"`

	// The number of seconds to delay power off after a PowerControl action.  Zero seconds indicates no delay to power off.
	PowerOffDelaySeconds float64 `json:"PowerOffDelaySeconds,omitempty"`

	// The number of seconds to delay power up after a power cycle or a PowerControl action.  Zero seconds indicates no delay to power up.
	PowerOnDelaySeconds float64 `json:"PowerOnDelaySeconds,omitempty"`

	// The number of seconds to delay power on after power has been restored.  Zero seconds indicates no delay.
	PowerRestoreDelaySeconds float64 `json:"PowerRestoreDelaySeconds,omitempty"`

	// The desired power state of the outlet group when power is restored after a power loss.
	PowerRestorePolicy PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`

	// The power state of the outlet group.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The power reading for this outlet group.
	PowerWatts SensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
