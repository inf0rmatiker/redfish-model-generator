/* -----------------------------------------------------------------
 * enum_bit_rate.go -
 *
 * DMTF Redfish BitRate enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type BitRate string

const (
	// A bit rate of 1200 bit/s.
	BitRate_1200 BitRate = "1200"

	// A bit rate of 2400 bit/s.
	BitRate_2400 BitRate = "2400"

	// A bit rate of 4800 bit/s.
	BitRate_4800 BitRate = "4800"

	// A bit rate of 9600 bit/s.
	BitRate_9600 BitRate = "9600"

	// A bit rate of 19200 bit/s.
	BitRate_19200 BitRate = "19200"

	// A bit rate of 38400 bit/s.
	BitRate_38400 BitRate = "38400"

	// A bit rate of 57600 bit/s.
	BitRate_57600 BitRate = "57600"

	// A bit rate of 115200 bit/s.
	BitRate_115200 BitRate = "115200"

	// A bit rate of 230400 bit/s.
	BitRate_230400 BitRate = "230400"

)
