/* -----------------------------------------------------------------
* enum_key_usage.go -
*
* DMTF Redfish KeyUsage enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type KeyUsage string

const (
	// Verifies digital signatures, other than signatures on certificates and CRLs.
	KeyUsage_DIGITAL_SIGNATURE KeyUsage = "DigitalSignature"

	// Verifies digital signatures, other than signatures on certificates and CRLs, and provides a non-repudiation service that protects against the signing entity falsely denying some action.
	KeyUsage_NON_REPUDIATION KeyUsage = "NonRepudiation"

	// Enciphers private or secret keys.
	KeyUsage_KEY_ENCIPHERMENT KeyUsage = "KeyEncipherment"

	// Directly enciphers raw user data without an intermediate symmetric cipher.
	KeyUsage_DATA_ENCIPHERMENT KeyUsage = "DataEncipherment"

	// Key agreement.
	KeyUsage_KEY_AGREEMENT KeyUsage = "KeyAgreement"

	// Verifies signatures on public key certificates.
	KeyUsage_KEY_CERT_SIGN KeyUsage = "KeyCertSign"

	// Verifies signatures on certificate revocation lists (CRLs).
	KeyUsage_CRL_SIGNING KeyUsage = "CRLSigning"

	// Enciphers data while performing a key agreement.
	KeyUsage_ENCIPHER_ONLY KeyUsage = "EncipherOnly"

	// Deciphers data while performing a key agreement.
	KeyUsage_DECIPHER_ONLY KeyUsage = "DecipherOnly"

	// TLS WWW server authentication.
	KeyUsage_SERVER_AUTHENTICATION KeyUsage = "ServerAuthentication"

	// TLS WWW client authentication.
	KeyUsage_CLIENT_AUTHENTICATION KeyUsage = "ClientAuthentication"

	// Signs downloadable executable code.
	KeyUsage_CODE_SIGNING KeyUsage = "CodeSigning"

	// Email protection.
	KeyUsage_EMAIL_PROTECTION KeyUsage = "EmailProtection"

	// Binds the hash of an object to a time.
	KeyUsage_TIMESTAMPING KeyUsage = "Timestamping"

	// Signs OCSP responses.
	KeyUsage_OCSP_SIGNING KeyUsage = "OCSPSigning"

)
