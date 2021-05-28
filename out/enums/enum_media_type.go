/* -----------------------------------------------------------------
 * enum_media_type.go -
 *
 * DMTF Redfish MediaType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MediaType string

const (
	// A CD-ROM format (ISO) image.
	MediaType_CD MediaType = "CD"

	// A floppy disk image.
	MediaType_FLOPPY MediaType = "Floppy"

	// An emulation of a USB storage device.
	MediaType_USB_STICK MediaType = "USBStick"

	// A DVD-ROM format image.
	MediaType_DVD MediaType = "DVD"

)
