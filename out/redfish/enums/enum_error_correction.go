/* -----------------------------------------------------------------
* enum_error_correction.go -
*
* DMTF Redfish ErrorCorrection enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ErrorCorrection string

const (
	// No ECC available.
	ErrorCorrection_NO_ECC ErrorCorrection = "NoECC"

	// Single bit data errors can be corrected by ECC.
	ErrorCorrection_SINGLE_BIT_ECC ErrorCorrection = "SingleBitECC"

	// Multibit data errors can be corrected by ECC.
	ErrorCorrection_MULTI_BIT_ECC ErrorCorrection = "MultiBitECC"

	// Address parity errors can be corrected.
	ErrorCorrection_ADDRESS_PARITY ErrorCorrection = "AddressParity"

)
