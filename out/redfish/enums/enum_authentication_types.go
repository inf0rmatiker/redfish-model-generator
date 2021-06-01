/* -----------------------------------------------------------------
* enum_authentication_types.go -
*
* DMTF Redfish AuthenticationTypes enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AuthenticationTypes string

const (
	// An opaque authentication token.
	AuthenticationTypes_TOKEN AuthenticationTypes = "Token"

	// A kerberos keytab.
	AuthenticationTypes_KERBEROS_KEYTAB AuthenticationTypes = "KerberosKeytab"

	// Username and password combination.
	AuthenticationTypes_USERNAME_AND_PASSWORD AuthenticationTypes = "UsernameAndPassword"

	// An OEM specific authentication mechanism.
	AuthenticationTypes_OEM AuthenticationTypes = "OEM"

)
