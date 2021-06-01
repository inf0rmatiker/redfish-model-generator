/* -----------------------------------------------------------------
* enum_address_state.go -
*
* DMTF Redfish AddressState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AddressState string

const (
	// This address is currently within both it's valid and preferred lifetimes as defined in RFC 4862.
	AddressState_PREFERRED AddressState = "Preferred"

	// This address is currently within it's valid lifetime, but is now outside of it's preferred lifetime as defined in RFC 4862.
	AddressState_DEPRECATED AddressState = "Deprecated"

	// This address is currently undergoing Duplicate Address Detection testing as defined in RFC 4862 section 5.4.
	AddressState_TENTATIVE AddressState = "Tentative"

	// This address has failed Duplicate Address Detection testing as defined in RFC 4862 section 5.4 and is not currently in use.
	AddressState_FAILED AddressState = "Failed"

)
