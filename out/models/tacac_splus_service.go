/* -----------------------------------------------------------------
* tacac_splus_service.go -
*
* DMTF Redfish TACACSplusService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Various settings to parse a TACACS+ service.
type TACACSplusService struct {
	// Indicates the allowed TACACS+ password exchange protocols.
	PasswordExchangeProtocols array `json:"PasswordExchangeProtocols,omitempty"`

	// Indicates the name of the TACACS+ argument name in an authorization request.
	PrivilegeLevelArgument string `json:"PrivilegeLevelArgument,omitempty"`

}
