/* -----------------------------------------------------------------
* port.go -
*
* DMTF Redfish Port resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Port - The Port schema contains properties that describe a port of a switch, controller, chassis, or any other device that could be connected to another entity.
// Modeled after DMTF Redfish schema Port
type Port struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The number of active lanes for this interface.
	ActiveWidth int `json:"ActiveWidth,omitempty"`

	// The current speed of this port.
	CurrentSpeedGbps float64 `json:"CurrentSpeedGbps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Gen-Z specific properties.
	GenZ GenZ `json:"GenZ,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether the interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	// The link network technology capabilities of this port.
	LinkNetworkTechnology LinkNetworkTechnology `json:"LinkNetworkTechnology,omitempty"`

	// The desired link state for this interface.
	LinkState LinkState `json:"LinkState,omitempty"`

	// The desired link status for this interface.
	LinkStatus LinkStatus `json:"LinkStatus,omitempty"`

	// The number of link state transitions for this interface.
	LinkTransitionIndicator int `json:"LinkTransitionIndicator,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the port.
	Location map[string]interface{} `json:"Location,omitempty"`

	// The maximum speed of this port as currently configured.
	MaxSpeedGbps float64 `json:"MaxSpeedGbps,omitempty"`

	// The link to the metrics associated with this port.
	Metrics map[string]interface{} `json:"Metrics,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The label of this port on the physical package for this port.
	PortId string `json:"PortId,omitempty"`

	// The physical connection medium for this port.
	PortMedium PortMedium `json:"PortMedium,omitempty"`

	// The protocol being sent over this port.
	PortProtocol map[string]interface{} `json:"PortProtocol,omitempty"`

	// The type of this port.
	PortType PortType `json:"PortType,omitempty"`

	// An indication of whether a signal is detected on this interface.
	SignalDetected bool `json:"SignalDetected,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The number of lanes, phys, or other physical transport links that this port contains.
	Width int `json:"Width,omitempty"`

}
