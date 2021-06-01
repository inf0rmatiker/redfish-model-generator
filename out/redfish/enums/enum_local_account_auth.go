/* -----------------------------------------------------------------
* enum_local_account_auth.go -
*
* DMTF Redfish LocalAccountAuth enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LocalAccountAuth string

const (
	// The service authenticates users based on the account service-defined accounts collection.
	LocalAccountAuth_ENABLED LocalAccountAuth = "Enabled"

	// The service never authenticates users based on the account service-defined accounts collection.
	LocalAccountAuth_DISABLED LocalAccountAuth = "Disabled"

	// The service authenticates users based on the account service-defined accounts collection only if any external account providers are currently unreachable.
	LocalAccountAuth_FALLBACK LocalAccountAuth = "Fallback"

)
