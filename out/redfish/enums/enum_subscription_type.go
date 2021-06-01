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
	// The subscription follows the Redfish specification for event notifications, which is done by a service sending an HTTP POST to the subscriber's destination URI.
	SubscriptionType_REDFISH_EVENT SubscriptionType = "RedfishEvent"

	// The subscription follows the HTML5 Server-Sent Event definition for event notifications.
	SubscriptionType_SSE SubscriptionType = "SSE"

)
