/* -----------------------------------------------------------------
 * enum_authentication_mode.go -
 *
 * DMTF Redfish AuthenticationMode enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type AuthenticationMode string

const (
	// Requests without any sort of authentication are allowed.
	AuthenticationMode_AUTH_NONE AuthenticationMode = "AuthNone"

	// Requests using HTTP Basic Authentication are allowed.
	AuthenticationMode_BASIC_AUTH AuthenticationMode = "BasicAuth"

	// Requests using Redfish Session Authentication are allowed.
	AuthenticationMode_REDFISH_SESSION_AUTH AuthenticationMode = "RedfishSessionAuth"

	// Requests using OEM authentication mechanisms are allowed.
	AuthenticationMode_OEM_AUTH AuthenticationMode = "OemAuth"

)
