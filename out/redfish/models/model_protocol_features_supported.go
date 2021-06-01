/* -----------------------------------------------------------------
* protocol_features_supported.go -
*
* DMTF Redfish ProtocolFeaturesSupported resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ProtocolFeaturesSupported - Contains information about protocol features supported by the service.
// Modeled after DMTF Redfish schema ProtocolFeaturesSupported
type ProtocolFeaturesSupported struct {
	// Contains information about the use of $expand in the service.
	ExpandQuery Expand `json:"ExpandQuery,omitempty"`

	// This indicates whether the filter query parameter is supported.
	FilterQuery bool `json:"FilterQuery,omitempty"`

	// This indicates whether the select query parameter is supported.
	SelectQuery bool `json:"SelectQuery,omitempty"`

}
