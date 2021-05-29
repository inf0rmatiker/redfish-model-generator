/* -----------------------------------------------------------------
* outlet.go -
*
* DMTF Redfish Outlet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Outlet schema contains definition for an electrical outlet.
type Outlet struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The current reading for this single phase outlet.
	CurrentAmps SensorCurrentExcerpt `json:"CurrentAmps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The combination of current-carrying conductors.
	ElectricalContext ElectricalContext `json:"ElectricalContext,omitempty"`

	// The energy reading for this outlet.
	EnergykWh SensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	// The frequency reading for this outlet.
	FrequencyHz SensorExcerpt `json:"FrequencyHz,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the indicator LED, which identifies the outlet.
	IndicatorLED IndicatorLED `json:"IndicatorLED,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The nominal voltage for this outlet.
	NominalVoltage NominalVoltageType `json:"NominalVoltage,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The type of receptacle according to NEMA, IEC, or regional standards.
	OutletType ReceptacleType `json:"OutletType,omitempty"`

	// The number of ungrounded current-carrying conductors (phases) and the total number of conductors (wires).
	PhaseWiringType PhaseWiringType `json:"PhaseWiringType,omitempty"`

	// The current readings for this outlet.
	PolyPhaseCurrentAmps CurrentSensors `json:"PolyPhaseCurrentAmps,omitempty"`

	// The voltage readings for this outlet.
	PolyPhaseVoltage VoltageSensors `json:"PolyPhaseVoltage,omitempty"`

	// The number of seconds to delay power on after a PowerControl action to cycle power.  Zero seconds indicates no delay.
	PowerCycleDelaySeconds float64 `json:"PowerCycleDelaySeconds,omitempty"`

	// Indicates if the outlet can be powered.
	PowerEnabled bool `json:"PowerEnabled,omitempty"`

	// The number of seconds to delay power off after a PowerControl action.  Zero seconds indicates no delay to power off.
	PowerOffDelaySeconds float64 `json:"PowerOffDelaySeconds,omitempty"`

	// The number of seconds to delay power up after a power cycle or a PowerControl action.  Zero seconds indicates no delay to power up.
	PowerOnDelaySeconds float64 `json:"PowerOnDelaySeconds,omitempty"`

	// The number of seconds to delay power on after power has been restored.  Zero seconds indicates no delay.
	PowerRestoreDelaySeconds float64 `json:"PowerRestoreDelaySeconds,omitempty"`

	// The desired power state of the outlet when power is restored after a power loss.
	PowerRestorePolicy PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`

	// The power state of the outlet.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The power reading for this outlet.
	PowerWatts SensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// The rated maximum current allowed for this outlet.
	RatedCurrentAmps float64 `json:"RatedCurrentAmps,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The voltage reading for this single phase outlet.
	Voltage SensorVoltageExcerpt `json:"Voltage,omitempty"`

	// The type of voltage applied to the outlet.
	VoltageType VoltageType `json:"VoltageType,omitempty"`

}
