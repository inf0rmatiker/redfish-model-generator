/* -----------------------------------------------------------------
 * enum_data_bits.go -
 *
 * DMTF Redfish DataBits enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DataBits string

const (
	// Five bits of data following the start bit.
	DataBits_5 DataBits = "5"

	// Six bits of data following the start bit.
	DataBits_6 DataBits = "6"

	// Seven bits of data following the start bit.
	DataBits_7 DataBits = "7"

	// Eight bits of data following the start bit.
	DataBits_8 DataBits = "8"

)
