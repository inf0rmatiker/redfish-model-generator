/* -----------------------------------------------------------------
 * enum_transfer_method.go -
 *
 * DMTF Redfish TransferMethod enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type TransferMethod string

const (
	// Stream image file data from the source URI.
	TransferMethod_STREAM TransferMethod = "Stream"

	// Upload the entire image file from the source URI to the service.
	TransferMethod_UPLOAD TransferMethod = "Upload"

)
