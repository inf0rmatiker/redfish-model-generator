/* -----------------------------------------------------------------
 * enum_subscription_type.go -
 *
 * DMTF Redfish SubscriptionType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SubscriptionType string

const (
	// The subscription follows the Redfish Specification for event notifications.  To send an event notification, a service sends an HTTP POST to the subscriber's destination URI.
	SubscriptionType_REDFISH_EVENT SubscriptionType = "RedfishEvent"

	// The subscription follows the HTML5 Server-Sent Event definition for event notifications.
	SubscriptionType_SSE SubscriptionType = "SSE"

	// The subscription follows the various versions of SNMP Traps for event notifications.
	SubscriptionType_SNMP_TRAP SubscriptionType = "SNMPTrap"

	// The subscription follows versions 2 and 3 of SNMP Inform for event notifications.
	SubscriptionType_SNMP_INFORM SubscriptionType = "SNMPInform"

	// The subscription sends Syslog messages for event notifications.
	SubscriptionType_SYSLOG SubscriptionType = "Syslog"

	// The subscription is an OEM subscription.
	SubscriptionType_OEM SubscriptionType = "OEM"

)
