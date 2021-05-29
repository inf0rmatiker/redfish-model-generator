/* -----------------------------------------------------------------
* nic_partitioning.go -
*
* DMTF Redfish NicPartitioning resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NIC Partitioning capability, status, and configuration for a controller.
type NicPartitioning struct {
	// An indication of whether the controller supports NIC function partitioning.
	NparCapable bool `json:"NparCapable,omitempty"`

	// An indication of whether NIC function partitioning is active on this controller.
	NparEnabled bool `json:"NparEnabled,omitempty"`

}
