/* -----------------------------------------------------------------
* processor.go -
*
* DMTF Redfish Processor resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Processor schema describes the information about a single processor that a system contains.  A processor includes both performance characteristics, clock speed, architecture, core count, and so on, and compatibility, such as the CPU ID instruction results.
type Processor struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The link to the collection of acceleration functions associated with this processor.
	AccelerationFunctions AccelerationFunctionCollection `json:"AccelerationFunctions,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the operating configuration that is applied to this processor.
	AppliedOperatingConfig OperatingConfig `json:"AppliedOperatingConfig,omitempty"`

	// The link to an assembly associated with this processor.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The state of the base frequency settings of the operation configuration applied to this processor.
	BaseSpeedPriorityState BaseSpeedPriorityState `json:"BaseSpeedPriorityState,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The properties for processors of the FPGA type.
	FPGA FPGA `json:"FPGA,omitempty"`

	// The firmware version of the processor.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The list of core identifiers corresponding to the cores that have been configured with the higher clock speed from the operating configuration applied to this processor.
	HighSpeedCoreIDs []integer `json:"HighSpeedCoreIDs,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The instruction set of the processor.
	InstructionSet InstructionSet `json:"InstructionSet,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the processor.
	Location Location `json:"Location,omitempty"`

	// The processor manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The maximum clock speed of the processor.
	MaxSpeedMHz int `json:"MaxSpeedMHz,omitempty"`

	// The maximum Thermal Design Power (TDP) in watts.
	MaxTDPWatts int `json:"MaxTDPWatts,omitempty"`

	// The link to the metrics associated with this processor.
	Metrics ProcessorMetrics `json:"Metrics,omitempty"`

	// The minimum clock speed of the processor in MHz.
	MinSpeedMHz int `json:"MinSpeedMHz,omitempty"`

	// The product model number of this device.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to the collection operating configurations that can be applied to this processor.
	OperatingConfigs OperatingConfigCollection `json:"OperatingConfigs,omitempty"`

	// Operating speed of the processor in MHz.
	OperatingSpeedMHz int `json:"OperatingSpeedMHz,omitempty"`

	// The part number of the processor.
	PartNumber string `json:"PartNumber,omitempty"`

	// The architecture of the processor.
	ProcessorArchitecture ProcessorArchitecture `json:"ProcessorArchitecture,omitempty"`

	// The identification information for this processor.
	ProcessorId ProcessorId `json:"ProcessorId,omitempty"`

	// The memory directly attached or integrated within this processor.
	ProcessorMemory array `json:"ProcessorMemory,omitempty"`

	// The type of processor.
	ProcessorType ProcessorType `json:"ProcessorType,omitempty"`

	// The serial number of the processor.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The socket or location of the processor.
	Socket string `json:"Socket,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The link to the collection of sub-processors associated with this system, such as cores or threads, that are part of a processor.
	SubProcessors ProcessorCollection `json:"SubProcessors,omitempty"`

	// The interface between the system and the processor.
	SystemInterface ProcessorInterface `json:"SystemInterface,omitempty"`

	// The nominal Thermal Design Power (TDP) in watts.
	TDPWatts int `json:"TDPWatts,omitempty"`

	// The total number of cores that this processor contains.
	TotalCores int `json:"TotalCores,omitempty"`

	// The total number of enabled cores that this processor contains.
	TotalEnabledCores int `json:"TotalEnabledCores,omitempty"`

	// The total number of execution threads that this processor supports.
	TotalThreads int `json:"TotalThreads,omitempty"`

	// The state of the turbo for this processor.
	TurboState TurboState `json:"TurboState,omitempty"`

	// The UUID for this processor.
	UUID UUID `json:"UUID,omitempty"`

	// The hardware version of the processor.
	Version string `json:"Version,omitempty"`

}
