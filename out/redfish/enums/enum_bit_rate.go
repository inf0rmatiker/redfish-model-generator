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
	BitRate_1200 BitRate = "1200"
	BitRate_2400 BitRate = "2400"
	BitRate_4800 BitRate = "4800"
	BitRate_9600 BitRate = "9600"
	BitRate_19200 BitRate = "19200"
	BitRate_38400 BitRate = "38400"
	BitRate_57600 BitRate = "57600"
	BitRate_115200 BitRate = "115200"
	BitRate_230400 BitRate = "230400"
)
