/* -----------------------------------------------------------------
 * enum_interface_type.go -
 *
 * DMTF Redfish InterfaceType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type InterfaceType string

const (
	// Trusted Platform Module (TPM) 1.2.
	InterfaceType_TPM1_2 InterfaceType = "TPM1_2"

	// Trusted Platform Module (TPM) 2.0.
	InterfaceType_TPM2_0 InterfaceType = "TPM2_0"

	// Trusted Cryptography Module (TCM) 1.0.
	InterfaceType_TCM1_0 InterfaceType = "TCM1_0"

)
