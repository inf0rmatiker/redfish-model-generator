/* -----------------------------------------------------------------
* ip_transport_details.go -
*
* DMTF Redfish IPTransportDetails resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IPTransportDetails - This object contains properties which specify the details of the transport supported by the endpoint. The properties which are present, is dependemt on the type of transport supported by the endpoint.
// Modeled after DMTF Redfish schema IPTransportDetails
type IPTransportDetails struct {
	// The IPv4 addresses assigned to the Endpoint.
	IPv4Address map[string]interface{} `json:"IPv4Address,omitempty"`

	// The IPv6 addresses assigned to the Endpoint.
	IPv6Address map[string]interface{} `json:"IPv6Address,omitempty"`

	// The UDP or TCP port number used by the Endpoint.
	Port float64 `json:"Port,omitempty"`

	// The protocol used by the connection entity.
	TransportProtocol map[string]interface{} `json:"TransportProtocol,omitempty"`

}
