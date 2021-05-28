/* -----------------------------------------------------------------
 * enum_account_provider_types.go -
 *
 * DMTF Redfish AccountProviderTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type AccountProviderTypes string

const (
	// An external Redfish service.
	AccountProviderTypes_REDFISH_SERVICE AccountProviderTypes = "RedfishService"

	// An external Active Directory service.
	AccountProviderTypes_ACTIVE_DIRECTORY_SERVICE AccountProviderTypes = "ActiveDirectoryService"

	// A generic external LDAP service.
	AccountProviderTypes_LDAP_SERVICE AccountProviderTypes = "LDAPService"

	// An OEM-specific external authentication or directory service.
	AccountProviderTypes_OEM AccountProviderTypes = "OEM"

	// An external TACACS+ service.
	AccountProviderTypes_TACAC_SPLUS AccountProviderTypes = "TACACSplus"

)
