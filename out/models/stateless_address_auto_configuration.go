/* -----------------------------------------------------------------
* stateless_address_auto_configuration.go -
*
* DMTF Redfish StatelessAddressAutoConfiguration resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Stateless address autoconfiguration (SLAAC) parameters for this interface.
type StatelessAddressAutoConfiguration struct {
	// An indication of whether IPv4 stateless address autoconfiguration (SLAAC) is enabled for this interface.
	IPv4AutoConfigEnabled bool `json:"IPv4AutoConfigEnabled,omitempty"`

	// An indication of whether IPv6 stateless address autoconfiguration (SLAAC) is enabled for this interface.
	IPv6AutoConfigEnabled bool `json:"IPv6AutoConfigEnabled,omitempty"`

}
