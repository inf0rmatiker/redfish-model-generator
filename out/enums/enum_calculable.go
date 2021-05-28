/* -----------------------------------------------------------------
 * enum_calculable.go -
 *
 * DMTF Redfish Calculable enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type Calculable string

const (
	// No calculations should be performed on the metric reading.
	Calculable_NON_CALCULATABLE Calculable = "NonCalculatable"

	// The sum of the metric reading across multiple instances is meaningful.
	Calculable_SUMMABLE Calculable = "Summable"

	// The sum of the metric reading across multiple instances is not meaningful.
	Calculable_NON_SUMMABLE Calculable = "NonSummable"

)
