/* -----------------------------------------------------------------
* generate_csr_response.go -
*
* DMTF Redfish GenerateCSRResponse resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// GenerateCSRResponse - The response body for the GenerateCSR action.
// Modeled after DMTF Redfish schema GenerateCSRResponse
type GenerateCSRResponse struct {
	// The string for the certificate signing request.
	CSRString string `json:"CSRString"`

	// A link to the certificate collection where the certificate will be installed.
	CertificateCollection map[string]interface{} `json:"CertificateCollection"`

}
