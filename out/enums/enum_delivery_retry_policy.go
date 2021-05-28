/* -----------------------------------------------------------------
 * enum_delivery_retry_policy.go -
 *
 * DMTF Redfish DeliveryRetryPolicy enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DeliveryRetryPolicy string

const (
	// The subscription is terminated after the maximum number of retries is reached.
	DeliveryRetryPolicy_TERMINATE_AFTER_RETRIES DeliveryRetryPolicy = "TerminateAfterRetries"

	// The subscription is suspended after the maximum number of retries is reached.
	DeliveryRetryPolicy_SUSPEND_RETRIES DeliveryRetryPolicy = "SuspendRetries"

	// The subscription is not suspended or terminated, and attempts at delivery of future events shall continue regardless of the number of retries.
	DeliveryRetryPolicy_RETRY_FOREVER DeliveryRetryPolicy = "RetryForever"

)
