/* -----------------------------------------------------------------
* compose_response.go -
*
* DMTF Redfish ComposeResponse resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The response body for the Compose action.
type ComposeResponse struct {
	// The manifest containing the compose operation response.
	Manifest Manifest `json:"Manifest,omitempty"`

	// The format of the request.
	RequestFormat ComposeRequestFormat `json:"RequestFormat"`

	// The type of request.
	RequestType ComposeRequestType `json:"RequestType"`

	// The identifier of the composition reservation that was created.
	ReservationId string `json:"ReservationId,omitempty"`

}
