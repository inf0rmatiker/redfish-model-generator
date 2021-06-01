/* -----------------------------------------------------------------
* renew_response.go -
*
* DMTF Redfish RenewResponse resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// RenewResponse - The response body for the Renew action.
// Modeled after DMTF Redfish schema RenewResponse
type RenewResponse struct {
	// The string for the certificate signing request.
	CSRString string `json:"CSRString"`

	// The link to the certificate being renewed.
	Certificate map[string]interface{} `json:"Certificate"`

}