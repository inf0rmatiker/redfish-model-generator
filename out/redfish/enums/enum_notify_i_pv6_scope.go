/* -----------------------------------------------------------------
* enum_notify_i_pv6_scope.go -
*
* DMTF Redfish NotifyIPv6Scope enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type NotifyIPv6Scope string

const (
	// SSDP Notify messages are sent to addresses in the IPv6 Local Link scope.
	NotifyIPv6Scope_LINK NotifyIPv6Scope = "Link"

	// SSDP Notify messages are sent to addresses in the IPv6 Local Site scope.
	NotifyIPv6Scope_SITE NotifyIPv6Scope = "Site"

	// SSDP Notify messages are sent to addresses in the IPv6 Local Organization scope.
	NotifyIPv6Scope_ORGANIZATION NotifyIPv6Scope = "Organization"

)
