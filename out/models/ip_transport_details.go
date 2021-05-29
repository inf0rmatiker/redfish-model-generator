/* -----------------------------------------------------------------
* ip_transport_details.go -
*
* DMTF Redfish IPTransportDetails resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type specifies the details of the transport supported by the endpoint.  The properties that are present are dependent on the type of transport supported by the endpoint.
type IPTransportDetails struct {
	// The IPv4 addresses assigned to the endpoint.
	IPv4Address IPv4Address `json:"IPv4Address,omitempty"`

	// The IPv6 addresses assigned to the endpoint.
	IPv6Address IPv6Address `json:"IPv6Address,omitempty"`

	// The UDP or TCP port number used by the endpoint.
	Port float64 `json:"Port,omitempty"`

	// The protocol used by the connection entity.
	TransportProtocol Protocol `json:"TransportProtocol,omitempty"`

}
