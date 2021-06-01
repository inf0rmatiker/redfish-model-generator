/* -----------------------------------------------------------------
* enum_privilege_type.go -
*
* DMTF Redfish PrivilegeType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PrivilegeType string

const (
	// Able to log into the service and read resources.
	PrivilegeType_LOGIN PrivilegeType = "Login"

	// Able to configure Manager resources.
	PrivilegeType_CONFIGURE_MANAGER PrivilegeType = "ConfigureManager"

	// Able to configure Users and their Accounts.
	PrivilegeType_CONFIGURE_USERS PrivilegeType = "ConfigureUsers"

	// Able to change the password for the current user Account.
	PrivilegeType_CONFIGURE_SELF PrivilegeType = "ConfigureSelf"

	// Able to configure components managed by this service.
	PrivilegeType_CONFIGURE_COMPONENTS PrivilegeType = "ConfigureComponents"

)
