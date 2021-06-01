/* -----------------------------------------------------------------
* enum_raid_type.go -
*
* DMTF Redfish RAIDType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type RAIDType string

const (
	// A placement policy where consecutive logical blocks of data are uniformly distributed across a set of independent storage devices without offering any form of redundancy.
	RAIDType_RAID0 RAIDType = "RAID0"

	// A placement policy where each logical block of data is stored on more than one independent storage device.
	RAIDType_RAID1 RAIDType = "RAID1"

	// A placement policy using parity-based protection where logical bytes of data are uniformly distributed across a set of independent storage devices and where the parity is stored on a dedicated independent storage device.
	RAIDType_RAID3 RAIDType = "RAID3"

	// A placement policy using parity-based protection where logical blocks of data are uniformly distributed across a set of independent storage devices and where the parity is stored on a dedicated independent storage device.
	RAIDType_RAID4 RAIDType = "RAID4"

	// A placement policy using parity-based protection for storing stripes of 'n' logical blocks of data and one logical block of parity across a set of 'n+1' independent storage devices where the parity and data blocks are interleaved across the storage devices.
	RAIDType_RAID5 RAIDType = "RAID5"

	// A placement policy using parity-based protection for storing stripes of 'n' logical blocks of data and two logical blocks of independent parity across a set of 'n+2' independent storage devices where the parity and data blocks are interleaved across the storage devices.
	RAIDType_RAID6 RAIDType = "RAID6"

	// A placement policy that creates a striped device (RAID 0) over a set of mirrored devices (RAID 1).
	RAIDType_RAID10 RAIDType = "RAID10"

	// A data placement policy that creates a mirrored device (RAID 1) over a set of striped devices (RAID 0).
	RAIDType_RAID01 RAIDType = "RAID01"

	// A placement policy that uses parity-based protection for storing stripes of 'n' logical blocks of data and three logical blocks of independent parity across a set of 'n+3' independent storage devices where the parity and data blocks are interleaved across the storage devices.
	RAIDType_RAID6_TP RAIDType = "RAID6TP"

	// A placement policy that uses a form of mirroring implemented over a set of independent storage devices where logical blocks are duplicated on a pair of independent storage devices so that data is uniformly distributed across the storage devices.
	RAIDType_RAID1_E RAIDType = "RAID1E"

	// A placement policy that uses a RAID 0 stripe set over two or more RAID 5 sets of independent storage devices.
	RAIDType_RAID50 RAIDType = "RAID50"

	// A placement policy that uses a RAID 0 stripe set over two or more RAID 6 sets of independent storage devices.
	RAIDType_RAID60 RAIDType = "RAID60"

	// A placement policy that creates a RAID 0 stripe set over two or more RAID 0 sets.
	RAIDType_RAID00 RAIDType = "RAID00"

	// A placement policy that uses a RAID 0 stripe set over two or more RAID 10 sets.
	RAIDType_RAID10_E RAIDType = "RAID10E"

	// A placement policy where each logical block of data is mirrored three times across a set of three independent storage devices.
	RAIDType_RAID1_TRIPLE RAIDType = "RAID1Triple"

	// A placement policy that uses a striped device (RAID 0) over a set of triple mirrored devices (RAID 1Triple).
	RAIDType_RAID10_TRIPLE RAIDType = "RAID10Triple"

	// A placement policy with no redundancy at the device level.
	RAIDType_NONE RAIDType = "None"

)
