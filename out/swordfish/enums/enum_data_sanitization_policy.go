/* -----------------------------------------------------------------
* enum_data_sanitization_policy.go -
*
* SNIA Swordfish DataSanitizationPolicy enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type DataSanitizationPolicy string

const (
	// No sanitization.
	DataSanitizationPolicy_NONE DataSanitizationPolicy = "None"

	// Sanitize data in all user-addressable storage locations for protection against simple non-invasive data recovery techniques.
	DataSanitizationPolicy_CLEAR DataSanitizationPolicy = "Clear"

	// Leverages the encryption of target data by enabling sanitization of the target data's encryption key. This leaves only the ciphertext remaining on the media, effectively sanitizing the data by preventing read-access. For more information, see NIST800-88 and ISO/IEC 27040.
	DataSanitizationPolicy_CRYPTOGRAPHIC_ERASE DataSanitizationPolicy = "CryptographicErase"

)
