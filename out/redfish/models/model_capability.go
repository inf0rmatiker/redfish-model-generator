/* -----------------------------------------------------------------
* capability.go -
*
* DMTF Redfish Capability resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Capability - This type describes a specific capability of a collection for a given use case.
// Modeled after DMTF Redfish schema Capability
type Capability struct {
	// Reference to the resource the client may GET to in order to understand how to form a POST request for a given collection.
	CapabilitiesObject map[string]interface{} `json:"CapabilitiesObject"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links"`

	// This property represents the use case in which a client may issue a POST request to the collection.
	UseCase UseCase `json:"UseCase"`

}
