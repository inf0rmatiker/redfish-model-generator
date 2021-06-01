/* -----------------------------------------------------------------
* enum_certificate_type.go -
*
* DMTF Redfish CertificateType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type CertificateType string

const (
	// A Privacy Enhanced Mail (PEM)-encoded certificate.
	CertificateType_PEM CertificateType = "PEM"

	// A Privacy Enhanced Mail (PEM)-encoded PKCS7 certificate.
	CertificateType_PKCS7 CertificateType = "PKCS7"

)
