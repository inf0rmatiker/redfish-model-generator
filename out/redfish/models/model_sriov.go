/* -----------------------------------------------------------------
* sriov.go -
*
* DMTF Redfish SRIOV resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SRIOV - Single-root input/output virtualization (SR-IOV) capabilities.
// Modeled after DMTF Redfish schema SRIOV
type SRIOV struct {
	// An indication of whether this controller supports single root input/output virtualization (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA) mode.
	SRIOVVEPACapable bool `json:"SRIOVVEPACapable,omitempty"`

}
