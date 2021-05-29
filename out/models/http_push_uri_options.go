/* -----------------------------------------------------------------
* http_push_uri_options.go -
*
* DMTF Redfish HttpPushUriOptions resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for HttpPushUri-provided software updates.
type HttpPushUriOptions struct {
	// The settings for when to apply HttpPushUri-provided firmware.
	HttpPushUriApplyTime HttpPushUriApplyTime `json:"HttpPushUriApplyTime,omitempty"`

}
