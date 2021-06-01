/* -----------------------------------------------------------------
* rekey_response.go -
*
* DMTF Redfish RekeyResponse resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// RekeyResponse - The response body for the Rekey action.
// Modeled after DMTF Redfish schema RekeyResponse
type RekeyResponse struct {
	// The string for the certificate signing request.
	CSRString string `json:"CSRString"`

	// The link to the certificate being rekeyed.
	Certificate map[string]interface{} `json:"Certificate"`

}
