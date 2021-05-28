/* -----------------------------------------------------------------
 * enum_tacac_splus_password_exchange_protocol.go -
 *
 * DMTF Redfish TACACSplusPasswordExchangeProtocol enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type TACACSplusPasswordExchangeProtocol string

const (
	// The ASCII Login method.
	TACACSplusPasswordExchangeProtocol_ASCII TACACSplusPasswordExchangeProtocol = "ASCII"

	// The PAP Login method.
	TACACSplusPasswordExchangeProtocol_PAP TACACSplusPasswordExchangeProtocol = "PAP"

	// The CHAP Login method.
	TACACSplusPasswordExchangeProtocol_CHAP TACACSplusPasswordExchangeProtocol = "CHAP"

	// The MS-CHAP v1 Login method.
	TACACSplusPasswordExchangeProtocol_MSCHA_PV1 TACACSplusPasswordExchangeProtocol = "MSCHAPv1"

	// The MS-CHAP v2 Login method.
	TACACSplusPasswordExchangeProtocol_MSCHA_PV2 TACACSplusPasswordExchangeProtocol = "MSCHAPv2"

)
