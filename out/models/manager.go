/* -----------------------------------------------------------------
* manager.go -
*
* DMTF Redfish Manager resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// In Redfish, a manager is a systems management entity that can implement or provide access to a Redfish service.  Examples of managers are BMCs, enclosure managers, management controllers, and other subsystems that are assigned manageability functions.  An implementation can have multiple managers, which might be directly accessible through a Redfish-defined interface.
type Manager struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An indication of whether the manager is configured for automatic Daylight Saving Time (DST) adjustment.
	AutoDSTEnabled bool `json:"AutoDSTEnabled,omitempty"`

	// The command shell service that this manager provides.
	CommandShell CommandShell `json:"CommandShell,omitempty"`

	// The current date and time with UTC offset that the manager uses to set or read time.
	DateTime string `json:"DateTime,omitempty"`

	// The time offset from UTC that the DateTime property is in `+HH:MM` format.
	DateTimeLocalOffset string `json:"DateTimeLocalOffset,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to a collection of NICs that this manager uses for network communication.
	EthernetInterfaces EthernetInterfaceCollection `json:"EthernetInterfaces,omitempty"`

	// The firmware version of this manager.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The information about the graphical console service of this manager.
	GraphicalConsole GraphicalConsole `json:"GraphicalConsole,omitempty"`

	// The link to a collection of host interfaces that this manager uses for local host communication.  Clients can find host interface configuration options and settings in this navigation property.
	HostInterfaces HostInterfaceCollection `json:"HostInterfaces,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The date and time when the manager was last reset or rebooted.
	LastResetTime string `json:"LastResetTime,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The link to a collection of logs that the manager uses.
	LogServices LogServiceCollection `json:"LogServices,omitempty"`

	// The type of manager that this resource represents.
	ManagerType ManagerType `json:"ManagerType,omitempty"`

	// The manufacturer of this manager.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model information of this manager, as defined by the manufacturer.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The link to the network services and their settings that the manager controls.
	NetworkProtocol ManagerNetworkProtocol `json:"NetworkProtocol,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number of the manager.
	PartNumber string `json:"PartNumber,omitempty"`

	// The current power state of the manager.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The redundancy information for the managers of this system.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The link to the account service resource for the remote manager that this resource represents.
	RemoteAccountService AccountService `json:"RemoteAccountService,omitempty"`

	// The URI of the Redfish service root for the remote manager that this resource represents.
	RemoteRedfishServiceUri string `json:"RemoteRedfishServiceUri,omitempty"`

	// The serial console service that this manager provides.
	SerialConsole SerialConsole `json:"SerialConsole,omitempty"`

	// The link to a collection of serial interfaces that this manager uses for serial and console communication.
	SerialInterfaces SerialInterfaceCollection `json:"SerialInterfaces,omitempty"`

	// The serial number of the manager.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The UUID of the Redfish service that is hosted by this manager.
	ServiceEntryPointUUID UUID `json:"ServiceEntryPointUUID,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The UUID for this manager.
	UUID UUID `json:"UUID,omitempty"`

	// The link to the Virtual Media services for this particular manager.
	VirtualMedia VirtualMediaCollection `json:"VirtualMedia,omitempty"`

}
