/* -----------------------------------------------------------------
* enum_line_input_voltage_type.go -
*
* DMTF Redfish LineInputVoltageType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LineInputVoltageType string

const (
	// The power supply line input voltage type cannot be determined.
	LineInputVoltageType_UNKNOWN LineInputVoltageType = "Unknown"

	// 100-127V AC input. Deprecated: Use AC120V.
	LineInputVoltageType_AC_LOW_LINE LineInputVoltageType = "ACLowLine"

	// 200-240V AC input. Deprecated: Use AC240V.
	LineInputVoltageType_AC_MID_LINE LineInputVoltageType = "ACMidLine"

	// 277V AC input. Deprecated: Use AC277V.
	LineInputVoltageType_AC_HIGH_LINE LineInputVoltageType = "ACHighLine"

	// -48V DC input.
	LineInputVoltageType_DC_NEG48_V LineInputVoltageType = "DCNeg48V"

	// High Voltage DC input (380V).
	LineInputVoltageType_DC380_V LineInputVoltageType = "DC380V"

	// AC 120V nominal input.
	LineInputVoltageType_AC120_V LineInputVoltageType = "AC120V"

	// AC 240V nominal input.
	LineInputVoltageType_AC240_V LineInputVoltageType = "AC240V"

	// AC 277V nominal input.
	LineInputVoltageType_AC277_V LineInputVoltageType = "AC277V"

	// Wide range AC or DC input.
	LineInputVoltageType_A_CAND_DC_WIDE_RANGE LineInputVoltageType = "ACandDCWideRange"

	// Wide range AC input.
	LineInputVoltageType_AC_WIDE_RANGE LineInputVoltageType = "ACWideRange"

	// DC 240V nominal input.
	LineInputVoltageType_DC240_V LineInputVoltageType = "DC240V"

)
