/* -----------------------------------------------------------------
* enum_calculation_algorithm_enum.go -
*
* DMTF Redfish CalculationAlgorithmEnum enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type CalculationAlgorithmEnum string

const (
	// The metric is calculated as the average metric reading over a sliding time interval.
	CalculationAlgorithmEnum_AVERAGE CalculationAlgorithmEnum = "Average"

	// The metric is calculated as the maximum metric reading over during a time interval.
	CalculationAlgorithmEnum_MAXIMUM CalculationAlgorithmEnum = "Maximum"

	// The metric is calculated as the minimum metric reading over a sliding time interval.
	CalculationAlgorithmEnum_MINIMUM CalculationAlgorithmEnum = "Minimum"

	// The metric is calculated as specified by an OEM.
	CalculationAlgorithmEnum_OEM CalculationAlgorithmEnum = "OEM"

)
