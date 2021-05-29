/* -----------------------------------------------------------------
* protocol_features_supported.go -
*
* DMTF Redfish ProtocolFeaturesSupported resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information about protocol features that the service supports.
type ProtocolFeaturesSupported struct {
	// The information about deep operations that the service supports.
	DeepOperations DeepOperations `json:"DeepOperations,omitempty"`

	// An indication of whether the service supports the excerpt query parameter.
	ExcerptQuery bool `json:"ExcerptQuery,omitempty"`

	// The information about the use of $expand in the service.
	ExpandQuery Expand `json:"ExpandQuery,omitempty"`

	// An indication of whether the service supports the $filter query parameter.
	FilterQuery bool `json:"FilterQuery,omitempty"`

	// An indication of whether the service supports the only query parameter.
	OnlyMemberQuery bool `json:"OnlyMemberQuery,omitempty"`

	// An indication of whether the service supports the $select query parameter.
	SelectQuery bool `json:"SelectQuery,omitempty"`

}
