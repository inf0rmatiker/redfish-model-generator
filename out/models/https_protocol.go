/* -----------------------------------------------------------------
* https_protocol.go -
*
* DMTF Redfish HTTPSProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for a network protocol associated with a manager.
type HTTPSProtocol struct {
	// The link to a collection of certificates used for HTTPS by this manager.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The protocol port.
	Port int `json:"Port,omitempty"`

	// An indication of whether the protocol is enabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

}
