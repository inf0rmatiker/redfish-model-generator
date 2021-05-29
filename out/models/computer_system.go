/* -----------------------------------------------------------------
* computer_system.go -
*
* DMTF Redfish ComputerSystem resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ComputerSystem schema represents a computer or system instance and the software-visible resources, or items within the data plane, such as memory, CPU, and other devices that it can access.  Details of those resources or subsystems are also linked through this resource.
type ComputerSystem struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The user-definable tag that can track this computer system for inventory or other client purposes.
	AssetTag string `json:"AssetTag,omitempty"`

	// The link to the BIOS settings associated with this system.
	Bios Bios `json:"Bios,omitempty"`

	// The version of the system BIOS or primary system firmware.
	BiosVersion string `json:"BiosVersion,omitempty"`

	// The boot settings for this system.
	Boot Boot `json:"Boot,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the collection of Ethernet interfaces associated with this system.
	EthernetInterfaces EthernetInterfaceCollection `json:"EthernetInterfaces,omitempty"`

	// The DNS host name, without any domain information.
	HostName string `json:"HostName,omitempty"`

	// The host watchdog timer functionality for this system.
	HostWatchdogTimer WatchdogTimer `json:"HostWatchdogTimer,omitempty"`

	// The services that this computer system supports.
	HostedServices HostedServices `json:"HostedServices,omitempty"`

	// The hosting roles that this computer system supports.
	HostingRoles array `json:"HostingRoles,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the indicator LED, which identifies the system.
	IndicatorLED IndicatorLED `json:"IndicatorLED,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The link to the collection of log services associated with this system.
	LogServices LogServiceCollection `json:"LogServices,omitempty"`

	// The manufacturer or OEM of this system.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The link to the collection of memory associated with this system.
	Memory MemoryCollection `json:"Memory,omitempty"`

	// The link to the collection of memory domains associated with this system.
	MemoryDomains MemoryDomainCollection `json:"MemoryDomains,omitempty"`

	// The central memory of the system in general detail.
	MemorySummary MemorySummary `json:"MemorySummary,omitempty"`

	// The product name for this system, without the manufacturer name.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The link to the collection of Network Interfaces associated with this system.
	NetworkInterfaces NetworkInterfaceCollection `json:"NetworkInterfaces,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to a collection of PCIe devices that this computer system uses.
	PCIeDevices array `json:"PCIeDevices,omitempty"`

	PCIeDevices@odata.count count `json:"PCIeDevices@odata.count,omitempty"`

	// The link to a collection of PCIe functions that this computer system uses.
	PCIeFunctions array `json:"PCIeFunctions,omitempty"`

	PCIeFunctions@odata.count count `json:"PCIeFunctions@odata.count,omitempty"`

	// The part number for this system.
	PartNumber string `json:"PartNumber,omitempty"`

	// The desired power state of the system when power is restored after a power loss.
	PowerRestorePolicy PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`

	// The current power state of the system.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The central processors of the system in general detail.
	ProcessorSummary ProcessorSummary `json:"ProcessorSummary,omitempty"`

	// The link to the collection of processors associated with this system.
	Processors ProcessorCollection `json:"Processors,omitempty"`

	// The link to a collection of redundancy entities.  Each entity specifies a kind and level of redundancy and a collection, or redundancy set, of other computer systems that provide the specified redundancy to this computer system.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The manufacturer SKU for this system.
	SKU string `json:"SKU,omitempty"`

	// The link to the UEFI Secure Boot associated with this system.
	SecureBoot SecureBoot `json:"SecureBoot,omitempty"`

	// The serial number for this system.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The link to the collection of storage devices associated with this system.
	SimpleStorage SimpleStorageCollection `json:"SimpleStorage,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The link to the collection of storage devices associated with this system.
	Storage StorageCollection `json:"Storage,omitempty"`

	// The sub-model for this system.
	SubModel string `json:"SubModel,omitempty"`

	// The type of computer system that this resource represents.
	SystemType SystemType `json:"SystemType,omitempty"`

	// An array of trusted modules in the system.
	TrustedModules array `json:"TrustedModules,omitempty"`

	// The UUID for this system.
	UUID UUID `json:"UUID,omitempty"`

}
