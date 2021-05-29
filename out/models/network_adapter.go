/* -----------------------------------------------------------------
* network_adapter.go -
*
* DMTF Redfish NetworkAdapter resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The NetworkAdapter schema represents a physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly resource associated with this adapter.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The link to a collection of certificates for device identity and attestation.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The set of network controllers ASICs that make up this NetworkAdapter.
	Controllers array `json:"Controllers,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the environment metrics for this network adapter.
	EnvironmentMetrics EnvironmentMetrics `json:"EnvironmentMetrics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable names for the network adapter.
	Identifiers array `json:"Identifiers,omitempty"`

	// Enable or disable LLDP globally for an adapter.
	LLDPEnabled bool `json:"LLDPEnabled,omitempty"`

	// The location of the network adapter.
	Location Location `json:"Location,omitempty"`

	// The manufacturer or OEM of this network adapter.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	Measurements array `json:"Measurements,omitempty"`

	// The link to the metrics associated with this adapter.
	Metrics NetworkAdapterMetrics `json:"Metrics,omitempty"`

	// The model string for this network adapter.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The link to the collection of network device functions associated with this network adapter.
	NetworkDeviceFunctions NetworkDeviceFunctionCollection `json:"NetworkDeviceFunctions,omitempty"`

	// The link to the collection of network ports associated with this network adapter.
	NetworkPorts NetworkPortCollection `json:"NetworkPorts,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Part number for this network adapter.
	PartNumber string `json:"PartNumber,omitempty"`

	// The link to the collection of ports associated with this network adapter.
	Ports PortCollection `json:"Ports,omitempty"`

	// The manufacturer SKU for this network adapter.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this network adapter.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
