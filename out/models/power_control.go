/* -----------------------------------------------------------------
* power_control.go -
*
* DMTF Redfish PowerControl resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PowerControl struct {
	OdataId string `json:"@odata.id"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The area, device, or set of devices to which this power control applies.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// The total amount of power that has been allocated or budgeted to chassis.
	PowerAllocatedWatts float64 `json:"PowerAllocatedWatts,omitempty"`

	// The amount of reserve power capacity, in watts, that remains.  This value is the PowerCapacityWatts value minus the PowerAllocatedWatts value.
	PowerAvailableWatts float64 `json:"PowerAvailableWatts,omitempty"`

	// The total amount of power that can be allocated to the chassis.  This value can be either the power supply capacity or the power budget that an upstream chassis assigns to this chassis.
	PowerCapacityWatts float64 `json:"PowerCapacityWatts,omitempty"`

	// The actual power that the chassis consumes, in watts.
	PowerConsumedWatts float64 `json:"PowerConsumedWatts,omitempty"`

	// The power limit status and configuration information for this chassis.
	PowerLimit PowerLimit `json:"PowerLimit,omitempty"`

	// The power readings for this chassis.
	PowerMetrics PowerMetric `json:"PowerMetrics,omitempty"`

	// The potential power, in watts, that the chassis requests, which might be higher than the current level being consumed because the requested power includes a budget that the chassis wants for future use.
	PowerRequestedWatts float64 `json:"PowerRequestedWatts,omitempty"`

	// An array of links to resources or objects associated with this power limit.
	RelatedItem []map[string]string `json:"RelatedItem,omitempty"`

	RelatedItem@odata.count count `json:"RelatedItem@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
