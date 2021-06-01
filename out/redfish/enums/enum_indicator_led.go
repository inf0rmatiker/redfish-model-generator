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
	// The Indicator LED is lit.
	IndicatorLED_LIT IndicatorLED = "Lit"

	// The Indicator LED is blinking.
	IndicatorLED_BLINKING IndicatorLED = "Blinking"

	// The Indicator LED is off.
	IndicatorLED_OFF IndicatorLED = "Off"

)
