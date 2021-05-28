/* -----------------------------------------------------------------
 * enum_compose_request_format.go -
 *
 * DMTF Redfish ComposeRequestFormat enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ComposeRequestFormat string

const (
	// The request body contains a manifest.
	ComposeRequestFormat_MANIFEST ComposeRequestFormat = "Manifest"

)
