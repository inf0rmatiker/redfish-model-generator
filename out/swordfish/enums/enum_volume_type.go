/* -----------------------------------------------------------------
* enum_volume_type.go -
*
* SNIA Swordfish VolumeType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type VolumeType string

const (
	// The volume is a raw physical device without any RAID or other virtualization applied.
	VolumeType_RAW_DEVICE VolumeType = "RawDevice"

	// The volume is a non-redundant storage device.
	VolumeType_NON_REDUNDANT VolumeType = "NonRedundant"

	// The volume is a mirrored device.
	VolumeType_MIRRORED VolumeType = "Mirrored"

	// The volume is a device which uses parity to retain redundant information.
	VolumeType_STRIPED_WITH_PARITY VolumeType = "StripedWithParity"

	// The volume is a spanned set of mirrored devices.
	VolumeType_SPANNED_MIRRORS VolumeType = "SpannedMirrors"

	// The volume is a spanned set of devices which uses parity to retain redundant information.
	VolumeType_SPANNED_STRIPES_WITH_PARITY VolumeType = "SpannedStripesWithParity"

)
