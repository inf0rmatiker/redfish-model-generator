/* -----------------------------------------------------------------
* port_metrics.go -
*
* DMTF Redfish PortMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The PortMetrics schema contains usage and health statistics for a switch device or component port summary.
type PortMetrics struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The port metrics specific to Gen-Z ports.
	GenZ GenZ `json:"GenZ,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The port metrics for network ports, including Ethernet, Fibre Channel, and InfiniBand, that are not specific to one of these protocols.
	Networking Networking `json:"Networking,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The total number of bytes received on a port since reset.
	RXBytes int `json:"RXBytes,omitempty"`

	// The total number of received errors on a port since reset.
	RXErrors int `json:"RXErrors,omitempty"`

	// The physical (phy) metrics for Serial Attached SCSI (SAS).  Each member represents a single phy.
	SAS array `json:"SAS,omitempty"`

	// The total number of bytes transmitted on a port since reset.
	TXBytes int `json:"TXBytes,omitempty"`

	// The total number of transmission errors on a port since reset.
	TXErrors int `json:"TXErrors,omitempty"`

	// The metrics for the transceivers in this port.  Each member represents a single transceiver.
	Transceivers array `json:"Transceivers,omitempty"`

}
