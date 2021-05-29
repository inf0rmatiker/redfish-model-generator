/* -----------------------------------------------------------------
* capability.go -
*
* DMTF Redfish Capability resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes a capability of a collection for a specific use case.
type Capability struct {
	// The link to the resource the client can issue a GET request against to understand how to form a POST request for a collection.
	CapabilitiesObject idRef `json:"CapabilitiesObject"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links"`

	// The use case in which a client can issue a POST request to the collection.
	UseCase UseCase `json:"UseCase"`

}
