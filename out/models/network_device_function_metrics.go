/* -----------------------------------------------------------------
* network_device_function_metrics.go -
*
* DMTF Redfish NetworkDeviceFunctionMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The NetworkDeviceFunctionMetrics schema contains usage and health statistics for a network function of a network adapter.
type NetworkDeviceFunctionMetrics struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The network function metrics specific to Ethernet adapters.
	Ethernet Ethernet `json:"Ethernet,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The average RX queue depth as the percentage.
	RXAvgQueueDepthPercent float64 `json:"RXAvgQueueDepthPercent,omitempty"`

	// The total number of bytes received on a network function.
	RXBytes int `json:"RXBytes,omitempty"`

	// The total number of frames received on a network function.
	RXFrames int `json:"RXFrames,omitempty"`

	// The total number of good multicast frames received on a network function since reset.
	RXMulticastFrames int `json:"RXMulticastFrames,omitempty"`

	// Whether nothing is in a network function's RX queues to DMA.
	RXQueuesEmpty bool `json:"RXQueuesEmpty,omitempty"`

	// The number of RX queues that are full.
	RXQueuesFull int `json:"RXQueuesFull,omitempty"`

	// The total number of good unicast frames received on a network function since reset.
	RXUnicastFrames int `json:"RXUnicastFrames,omitempty"`

	// The average TX queue depth as the percentage.
	TXAvgQueueDepthPercent float64 `json:"TXAvgQueueDepthPercent,omitempty"`

	// The total number of bytes sent on a network function.
	TXBytes int `json:"TXBytes,omitempty"`

	// The total number of frames sent on a network function.
	TXFrames int `json:"TXFrames,omitempty"`

	// The total number of good multicast frames transmitted on a network function since reset.
	TXMulticastFrames int `json:"TXMulticastFrames,omitempty"`

	// Whether all TX queues for a network function are empty.
	TXQueuesEmpty bool `json:"TXQueuesEmpty,omitempty"`

	// The number of TX queues that are full.
	TXQueuesFull int `json:"TXQueuesFull,omitempty"`

	// The total number of good unicast frames transmitted on a network function since reset.
	TXUnicastFrames int `json:"TXUnicastFrames,omitempty"`

}
