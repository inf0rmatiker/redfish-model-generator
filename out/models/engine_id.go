/* -----------------------------------------------------------------
* engine_id.go -
*
* DMTF Redfish EngineId resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The engine ID.
type EngineId struct {
	// The architecture identifier.
	ArchitectureId string `json:"ArchitectureId,omitempty"`

	// The enterprise specific method.
	EnterpriseSpecificMethod string `json:"EnterpriseSpecificMethod,omitempty"`

	// The private enterprise ID.
	PrivateEnterpriseId string `json:"PrivateEnterpriseId,omitempty"`

}
