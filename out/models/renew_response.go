/* -----------------------------------------------------------------
* renew_response.go -
*
* DMTF Redfish RenewResponse resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The response body for the Renew action.
type RenewResponse struct {
	// The string for the certificate signing request.
	CSRString string `json:"CSRString"`

	// The link to the certificate being renewed.
	Certificate Certificate `json:"Certificate"`

}
