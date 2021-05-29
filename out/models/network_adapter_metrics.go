/* -----------------------------------------------------------------
* network_adapter_metrics.go -
*
* DMTF Redfish NetworkAdapterMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The NetworkAdapterMetrics schema contains usage and health statistics for a network adapter.
type NetworkAdapterMetrics struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The device CPU core utilization as a percentage.
	CPUCorePercent float64 `json:"CPUCorePercent,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The host bus, such as PCIe, RX utilization as a percentage.
	HostBusRXPercent float64 `json:"HostBusRXPercent,omitempty"`

	// The host bus, such as PCIe, TX utilization as a percentage.
	HostBusTXPercent float64 `json:"HostBusTXPercent,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The total number of NC-SI bytes received since reset.
	NCSIRXBytes int `json:"NCSIRXBytes,omitempty"`

	// The total number of NC-SI frames received since reset.
	NCSIRXFrames int `json:"NCSIRXFrames,omitempty"`

	// The total number of NC-SI bytes sent since reset.
	NCSITXBytes int `json:"NCSITXBytes,omitempty"`

	// The total number of NC-SI frames sent since reset.
	NCSITXFrames int `json:"NCSITXFrames,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The total number of bytes received since reset.
	RXBytes int `json:"RXBytes,omitempty"`

	// The total number of good multicast frames received since reset.
	RXMulticastFrames int `json:"RXMulticastFrames,omitempty"`

	// The total number of good unicast frames received since reset.
	RXUnicastFrames int `json:"RXUnicastFrames,omitempty"`

	// The total number of bytes transmitted since reset.
	TXBytes int `json:"TXBytes,omitempty"`

	// The total number of good multicast frames transmitted since reset.
	TXMulticastFrames int `json:"TXMulticastFrames,omitempty"`

	// The total number of good unicast frames transmitted since reset.
	TXUnicastFrames int `json:"TXUnicastFrames,omitempty"`

}
