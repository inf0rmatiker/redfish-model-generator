/* -----------------------------------------------------------------
* generate_csr_response.go -
*
* DMTF Redfish GenerateCSRResponse resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The response body for the GenerateCSR action.
type GenerateCSRResponse struct {
	// The string for the certificate signing request.
	CSRString string `json:"CSRString"`

	// The link to the certificate collection where the certificate is installed.
	CertificateCollection CertificateCollection `json:"CertificateCollection"`

}
