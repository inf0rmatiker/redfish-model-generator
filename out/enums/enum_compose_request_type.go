/* -----------------------------------------------------------------
 * enum_compose_request_type.go -
 *
 * DMTF Redfish ComposeRequestType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ComposeRequestType string

const (
	// Preview the outcome of the operations specified by the manifest.
	ComposeRequestType_PREVIEW ComposeRequestType = "Preview"

	// Preview the outcome of the operations specified by the manifest and reserve resources.
	ComposeRequestType_PREVIEW_RESERVE ComposeRequestType = "PreviewReserve"

	// Perform the requested operations specified by the manifest and modify resources as needed.
	ComposeRequestType_APPLY ComposeRequestType = "Apply"

)
