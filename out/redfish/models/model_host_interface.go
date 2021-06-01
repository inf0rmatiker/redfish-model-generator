/* -----------------------------------------------------------------
* host_interface.go -
*
* DMTF Redfish HostInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// HostInterface - This schema defines a single Host Interface resource.
// Modeled after DMTF Redfish schema HostInterface
type HostInterface struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Indicates the authentication modes available on this interface.
	AuthenticationModes []AuthenticationMode `json:"AuthenticationModes,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Indicates whether this interface is accessible by external entities.
	ExternallyAccessible bool `json:"ExternallyAccessible,omitempty"`

	// Indicates whether this firmware authentication is enabled for this interface.
	FirmwareAuthEnabled bool `json:"FirmwareAuthEnabled,omitempty"`

	// The Role used for firmware authentication on this interface.
	FirmwareAuthRoleId string `json:"FirmwareAuthRoleId,omitempty"`

	// The Redfish link to the collection of network interface controllers or cards (NICs) that a Computer System uses to communicate with this Host Interface.
	HostEthernetInterfaces map[string]interface{} `json:"HostEthernetInterfaces,omitempty"`

	// Indicates the Host Interface type for this interface.
	HostInterfaceType HostInterfaceType `json:"HostInterfaceType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Indicates whether this interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	// Indicates whether this kernel authentication is enabled for this interface.
	KernelAuthEnabled bool `json:"KernelAuthEnabled,omitempty"`

	// The Role used for kernel authentication on this interface.
	KernelAuthRoleId string `json:"KernelAuthRoleId,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The Redfish link to a single network interface controllers or cards (NIC) that this Manager uses for network communication with this Host Interface.
	ManagerEthernetInterface map[string]interface{} `json:"ManagerEthernetInterface,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The Redfish link to the network services and their settings that the Manager controls.  It is here that clients will find network configuration options as well as network services.
	NetworkProtocol map[string]interface{} `json:"NetworkProtocol,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

}
