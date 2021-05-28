/* -----------------------------------------------------------------
 * enum_indicator_led.go -
 *
 * DMTF Redfish IndicatorLED enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IndicatorLED string

const (
	// The state of the indicator LED cannot be determined.
	IndicatorLED_UNKNOWN IndicatorLED = "Unknown"

	// The indicator LED is lit.
	IndicatorLED_LIT IndicatorLED = "Lit"

	// The indicator LED is blinking.
	IndicatorLED_BLINKING IndicatorLED = "Blinking"

	// The indicator LED is off.
	IndicatorLED_OFF IndicatorLED = "Off"

)
