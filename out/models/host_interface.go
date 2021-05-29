/* -----------------------------------------------------------------
* host_interface.go -
*
* DMTF Redfish HostInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The properties associated with a Host Interface.  A Host Interface is a connection between host software and a Redfish Service.
type HostInterface struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The role when no authentication on this interface is used.
	AuthNoneRoleId string `json:"AuthNoneRoleId,omitempty"`

	// The authentication modes available on this interface.
	AuthenticationModes array `json:"AuthenticationModes,omitempty"`

	// The credential bootstrapping settings for this interface.
	CredentialBootstrapping CredentialBootstrapping `json:"CredentialBootstrapping,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether external entities can access this interface.  External entities are non-host entities.  For example, if the host and manager are connected through a switch and the switch also exposes an external port on the system, external clients can also use the interface, and this property value is `true`.
	ExternallyAccessible bool `json:"ExternallyAccessible,omitempty"`

	// An indication of whether this firmware authentication is enabled for this interface.
	FirmwareAuthEnabled bool `json:"FirmwareAuthEnabled,omitempty"`

	// The Role used for firmware authentication on this interface.
	FirmwareAuthRoleId string `json:"FirmwareAuthRoleId,omitempty"`

	// A link to the collection of network interface controllers or cards (NICs) that a computer system uses to communicate with this Host Interface.
	HostEthernetInterfaces EthernetInterfaceCollection `json:"HostEthernetInterfaces,omitempty"`

	// The Host Interface type for this interface.
	HostInterfaceType HostInterfaceType `json:"HostInterfaceType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether this interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	// An indication of whether this kernel authentication is enabled for this interface.
	KernelAuthEnabled bool `json:"KernelAuthEnabled,omitempty"`

	// The Role used for kernel authentication on this interface.
	KernelAuthRoleId string `json:"KernelAuthRoleId,omitempty"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// A link to a single network interface controllers or cards (NIC) that this manager uses for network communication with this Host Interface.
	ManagerEthernetInterface EthernetInterface `json:"ManagerEthernetInterface,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// A link to the network services and their settings that the manager controls.  In this property, clients find configuration options for the network and network services.
	NetworkProtocol ManagerNetworkProtocol `json:"NetworkProtocol,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

}
