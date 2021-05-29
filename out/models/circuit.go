/* -----------------------------------------------------------------
* circuit.go -
*
* DMTF Redfish Circuit resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This is the schema definition for an electrical circuit.
type Circuit struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The state of the over current protection device.
	BreakerState BreakerStates `json:"BreakerState,omitempty"`

	// The type of circuit.
	CircuitType CircuitType `json:"CircuitType,omitempty"`

	// Designates if this is a critical circuit.
	CriticalCircuit bool `json:"CriticalCircuit,omitempty"`

	// The current reading for this single phase circuit.
	CurrentAmps SensorCurrentExcerpt `json:"CurrentAmps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The combination of current-carrying conductors.
	ElectricalContext ElectricalContext `json:"ElectricalContext,omitempty"`

	// The energy reading for this circuit.
	EnergykWh SensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	// The frequency reading for this circuit.
	FrequencyHz SensorExcerpt `json:"FrequencyHz,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the indicator LED, which identifies the circuit.
	IndicatorLED IndicatorLED `json:"IndicatorLED,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The nominal voltage for this circuit.
	NominalVoltage NominalVoltageType `json:"NominalVoltage,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The number of ungrounded current-carrying conductors (phases) and the total number of conductors (wires).
	PhaseWiringType PhaseWiringType `json:"PhaseWiringType,omitempty"`

	// The type of plug according to NEMA, IEC, or regional standards.
	PlugType PlugType `json:"PlugType,omitempty"`

	// The current readings for this circuit.
	PolyPhaseCurrentAmps CurrentSensors `json:"PolyPhaseCurrentAmps,omitempty"`

	// The energy readings for this circuit.
	PolyPhaseEnergykWh EnergySensors `json:"PolyPhaseEnergykWh,omitempty"`

	// The power readings for this circuit.
	PolyPhasePowerWatts PowerSensors `json:"PolyPhasePowerWatts,omitempty"`

	// The voltage readings for this circuit.
	PolyPhaseVoltage VoltageSensors `json:"PolyPhaseVoltage,omitempty"`

	// The number of seconds to delay power on after a PowerControl action to cycle power.  Zero seconds indicates no delay.
	PowerCycleDelaySeconds float64 `json:"PowerCycleDelaySeconds,omitempty"`

	// Indicates if the circuit can be powered.
	PowerEnabled bool `json:"PowerEnabled,omitempty"`

	// The number of seconds to delay power off after a PowerControl action.  Zero seconds indicates no delay to power off.
	PowerOffDelaySeconds float64 `json:"PowerOffDelaySeconds,omitempty"`

	// The number of seconds to delay power up after a power cycle or a PowerControl action.  Zero seconds indicates no delay to power up.
	PowerOnDelaySeconds float64 `json:"PowerOnDelaySeconds,omitempty"`

	// The number of seconds to delay power on after power has been restored.  Zero seconds indicates no delay.
	PowerRestoreDelaySeconds float64 `json:"PowerRestoreDelaySeconds,omitempty"`

	// The desired power state of the circuit when power is restored after a power loss.
	PowerRestorePolicy PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`

	// The power state of the circuit.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The power reading for this circuit.
	PowerWatts SensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// The rated maximum current allowed for this circuit.
	RatedCurrentAmps float64 `json:"RatedCurrentAmps,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The voltage reading for this single phase circuit.
	Voltage SensorVoltageExcerpt `json:"Voltage,omitempty"`

	// The type of voltage applied to the circuit.
	VoltageType VoltageType `json:"VoltageType,omitempty"`

}