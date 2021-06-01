/* -----------------------------------------------------------------
* processor.go -
*
* DMTF Redfish Processor resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Processor - The Processor schema describes the information about a single processor that a system contains.  A processor includes both performance characteristics, clock speed, architecture, core count, and so on, and compatibility, such as the CPU ID instruction results.
// Modeled after DMTF Redfish schema Processor
type Processor struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The link to the collection of acceleration functions associated with this processor.
	AccelerationFunctions map[string]interface{} `json:"AccelerationFunctions,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to an assembly associated with this processor.
	Assembly map[string]interface{} `json:"Assembly,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The properties for processors of the FPGA type.
	FPGA FPGA `json:"FPGA,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The instruction set of the processor.
	InstructionSet InstructionSet `json:"InstructionSet,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the processor.
	Location map[string]interface{} `json:"Location,omitempty"`

	// The processor manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The maximum clock speed of the processor.
	MaxSpeedMHz int `json:"MaxSpeedMHz,omitempty"`

	// The maximum Thermal Design Power (TDP) in watts.
	MaxTDPWatts int `json:"MaxTDPWatts,omitempty"`

	// The link to the metrics associated with this processor.
	Metrics map[string]interface{} `json:"Metrics,omitempty"`

	// The product model number of this device.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The architecture of the processor.
	ProcessorArchitecture ProcessorArchitecture `json:"ProcessorArchitecture,omitempty"`

	// The identification information for this processor.
	ProcessorId ProcessorId `json:"ProcessorId,omitempty"`

	// The memory directly attached or integrated within this processor.
	ProcessorMemory []ProcessorMemory `json:"ProcessorMemory,omitempty"`

	// The type of processor.
	ProcessorType ProcessorType `json:"ProcessorType,omitempty"`

	// The socket or location of the processor.
	Socket string `json:"Socket,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The link to the collection of sub-processors associated with this system, such as cores or threads, that are part of a processor.
	SubProcessors map[string]interface{} `json:"SubProcessors,omitempty"`

	// The nominal Thermal Design Power (TDP) in watts.
	TDPWatts int `json:"TDPWatts,omitempty"`

	// The total number of cores that this processor contains.
	TotalCores int `json:"TotalCores,omitempty"`

	// The total number of execution threads that this processor supports.
	TotalThreads int `json:"TotalThreads,omitempty"`

	// The UUID for this processor.
	UUID map[string]interface{} `json:"UUID,omitempty"`

}
