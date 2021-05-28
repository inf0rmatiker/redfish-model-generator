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
	// This address is currently within both its RFC4862-defined valid and preferred lifetimes.
	AddressState_PREFERRED AddressState = "Preferred"

	// This address is currently within its valid lifetime but is now outside its RFC4862-defined preferred lifetime.
	AddressState_DEPRECATED AddressState = "Deprecated"

	// This address is currently undergoing Duplicate Address Detection (DAD) testing, as defined in RFC4862, section 5.4.
	AddressState_TENTATIVE AddressState = "Tentative"

	// This address has failed Duplicate Address Detection (DAD) testing, as defined in RFC4862, section 5.4, and is not currently in use.
	AddressState_FAILED AddressState = "Failed"

)
