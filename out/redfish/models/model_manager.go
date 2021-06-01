/* -----------------------------------------------------------------
* manager.go -
*
* DMTF Redfish Manager resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Manager - This is the schema definition for a Manager.  Examples of managers are BMCs, Enclosure Managers, Management Controllers and other subsystems assigned managability functions.
// Modeled after DMTF Redfish schema Manager
type Manager struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Indicates whether the manager is configured for automatic DST adjustment.
	AutoDSTEnabled bool `json:"AutoDSTEnabled,omitempty"`

	// Information about the Command Shell service provided by this manager.
	CommandShell CommandShell `json:"CommandShell,omitempty"`

	// The current DateTime (with offset) for the manager, used to set or read time.
	DateTime string `json:"DateTime,omitempty"`

	// The time offset from UTC that the DateTime property is set to in format: +06:00 .
	DateTimeLocalOffset string `json:"DateTimeLocalOffset,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// This is a reference to a collection of NICs that this manager uses for network communication.  It is here that clients will find NIC configuration options and settings.
	EthernetInterfaces map[string]interface{} `json:"EthernetInterfaces,omitempty"`

	// The firmware version of this Manager.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The value of this property shall contain the information about the Graphical Console (KVM-IP) service of this manager.
	GraphicalConsole GraphicalConsole `json:"GraphicalConsole,omitempty"`

	// This is a reference to a collection of Host Interfaces that this manager uses for local host communication.  It is here that clients will find Host Interface configuration options and settings.
	HostInterfaces map[string]interface{} `json:"HostInterfaces,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// This is a reference to a collection of Logs used by the manager.
	LogServices map[string]interface{} `json:"LogServices,omitempty"`

	// This property represents the type of manager that this resource represents.
	ManagerType ManagerType `json:"ManagerType,omitempty"`

	// The model information of this Manager as defined by the manufacturer.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// This is a reference to the network services and their settings that the manager controls.  It is here that clients will find network configuration options as well as network services.
	NetworkProtocol map[string]interface{} `json:"NetworkProtocol,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the current power state of the Manager.
	PowerState map[string]interface{} `json:"PowerState,omitempty"`

	// Redundancy information for the managers of this system.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	// Information about the Serial Console service provided by this manager.
	SerialConsole SerialConsole `json:"SerialConsole,omitempty"`

	// This is a reference to a collection of serial interfaces that this manager uses for serial and console communication.  It is here that clients will find serial configuration options and settings.
	SerialInterfaces map[string]interface{} `json:"SerialInterfaces,omitempty"`

	// The UUID of the Redfish Service provided by this manager.
	ServiceEntryPointUUID map[string]interface{} `json:"ServiceEntryPointUUID,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The Universal Unique Identifier (UUID) for this Manager.
	UUID map[string]interface{} `json:"UUID,omitempty"`

	// This is a reference to the Virtual Media services for this particular manager.
	VirtualMedia map[string]interface{} `json:"VirtualMedia,omitempty"`

}
