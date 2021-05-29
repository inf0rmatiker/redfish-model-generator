/* -----------------------------------------------------------------
* memory.go -
*
* DMTF Redfish Memory resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Memory schema represents a memory device, such as a DIMM, and its configuration.
type Memory struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The boundary that memory regions are allocated on, measured in mebibytes (MiB).
	AllocationAlignmentMiB int `json:"AllocationAlignmentMiB,omitempty"`

	// The size of the smallest unit of allocation for a memory region in mebibytes (MiB).
	AllocationIncrementMiB int `json:"AllocationIncrementMiB,omitempty"`

	// Speeds supported by this memory device.
	AllowedSpeedsMHz []integer `json:"AllowedSpeedsMHz,omitempty"`

	// The link to the assembly resource associated with this memory device.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The base module type of the memory device.
	BaseModuleType BaseModuleType `json:"BaseModuleType,omitempty"`

	// The bus width, in bits.
	BusWidthBits int `json:"BusWidthBits,omitempty"`

	// Total size of the cache portion memory in MiB.
	CacheSizeMiB int `json:"CacheSizeMiB,omitempty"`

	// Memory capacity in mebibytes (MiB).
	CapacityMiB int `json:"CapacityMiB,omitempty"`

	// An indication of whether the configuration of this memory device is locked and cannot be altered.
	ConfigurationLocked bool `json:"ConfigurationLocked,omitempty"`

	// Data width in bits.
	DataWidthBits int `json:"DataWidthBits,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Device ID.
	DeviceID string `json:"DeviceID,omitempty"`

	// Location of the memory device in the platform.
	DeviceLocator string `json:"DeviceLocator,omitempty"`

	// Error correction scheme supported for this memory device.
	ErrorCorrection ErrorCorrection `json:"ErrorCorrection,omitempty"`

	// Version of API supported by the firmware.
	FirmwareApiVersion string `json:"FirmwareApiVersion,omitempty"`

	// Revision of firmware on the memory controller.
	FirmwareRevision string `json:"FirmwareRevision,omitempty"`

	// Function classes by the memory device.
	FunctionClasses []string `json:"FunctionClasses,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether rank spare is enabled for this memory device.
	IsRankSpareEnabled bool `json:"IsRankSpareEnabled,omitempty"`

	// An indication of whether a spare device is enabled for this memory device.
	IsSpareDeviceEnabled bool `json:"IsSpareDeviceEnabled,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the memory device.
	Location Location `json:"Location,omitempty"`

	// Total size of the logical memory in MiB.
	LogicalSizeMiB int `json:"LogicalSizeMiB,omitempty"`

	// The memory device manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// Set of maximum power budgets supported by the memory device in milliwatts.
	MaxTDPMilliWatts []integer `json:"MaxTDPMilliWatts,omitempty"`

	// Type details of the memory device.
	MemoryDeviceType MemoryDeviceType `json:"MemoryDeviceType,omitempty"`

	// Memory connection information to sockets and memory controllers.
	MemoryLocation MemoryLocation `json:"MemoryLocation,omitempty"`

	// Media of this memory device.
	MemoryMedia array `json:"MemoryMedia,omitempty"`

	// The manufacturer ID of the memory subsystem controller of this memory device.
	MemorySubsystemControllerManufacturerID string `json:"MemorySubsystemControllerManufacturerID,omitempty"`

	// The product ID of the memory subsystem controller of this memory device.
	MemorySubsystemControllerProductID string `json:"MemorySubsystemControllerProductID,omitempty"`

	// The type of memory device.
	MemoryType MemoryType `json:"MemoryType,omitempty"`

	// The link to the metrics associated with this memory device.
	Metrics MemoryMetrics `json:"Metrics,omitempty"`

	// The manufacturer ID of this memory device.
	ModuleManufacturerID string `json:"ModuleManufacturerID,omitempty"`

	// The product ID of this memory device.
	ModuleProductID string `json:"ModuleProductID,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// Total size of the non-volatile portion memory in MiB.
	NonVolatileSizeMiB int `json:"NonVolatileSizeMiB,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Memory modes supported by the memory device.
	OperatingMemoryModes array `json:"OperatingMemoryModes,omitempty"`

	// Operating speed of the memory device in MHz or MT/s as appropriate.
	OperatingSpeedMhz int `json:"OperatingSpeedMhz,omitempty"`

	// The product part number of this device.
	PartNumber string `json:"PartNumber,omitempty"`

	// Total number of persistent regions this memory device can support.
	PersistentRegionNumberLimit int `json:"PersistentRegionNumberLimit,omitempty"`

	// Total size of persistent regions in mebibytes (MiB).
	PersistentRegionSizeLimitMiB int `json:"PersistentRegionSizeLimitMiB,omitempty"`

	// Maximum size of a single persistent region in mebibytes (MiB).
	PersistentRegionSizeMaxMiB int `json:"PersistentRegionSizeMaxMiB,omitempty"`

	// Power management policy information.
	PowerManagementPolicy PowerManagementPolicy `json:"PowerManagementPolicy,omitempty"`

	// Number of ranks available in the memory device.
	RankCount int `json:"RankCount,omitempty"`

	// Memory regions information within the memory device.
	Regions array `json:"Regions,omitempty"`

	// Security capabilities of the memory device.
	SecurityCapabilities SecurityCapabilities `json:"SecurityCapabilities,omitempty"`

	// The current security state of this memory device.
	SecurityState SecurityStates `json:"SecurityState,omitempty"`

	// The product serial number of this device.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// Number of unused spare devices available in the memory device.
	SpareDeviceCount int `json:"SpareDeviceCount,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// Subsystem device ID.
	SubsystemDeviceID string `json:"SubsystemDeviceID,omitempty"`

	// SubSystem vendor ID.
	SubsystemVendorID string `json:"SubsystemVendorID,omitempty"`

	// Vendor ID.
	VendorID string `json:"VendorID,omitempty"`

	// Total number of volatile regions this memory device can support.
	VolatileRegionNumberLimit int `json:"VolatileRegionNumberLimit,omitempty"`

	// Total size of volatile regions in mebibytes (MiB).
	VolatileRegionSizeLimitMiB int `json:"VolatileRegionSizeLimitMiB,omitempty"`

	// Maximum size of a single volatile region in mebibytes (MiB).
	VolatileRegionSizeMaxMiB int `json:"VolatileRegionSizeMaxMiB,omitempty"`

	// Total size of the volatile portion memory in MiB.
	VolatileSizeMiB int `json:"VolatileSizeMiB,omitempty"`

}
