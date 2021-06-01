/* -----------------------------------------------------------------
* io_workload_component.go -
*
* SNIA Swordfish IOWorkloadComponent resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOWorkloadComponent - Describe a component of a IO workload.
// Modeled after SNIA Swordfish schema IOWorkloadComponent
type IOWorkloadComponent struct {
	// Average I/O Size for this component.
	AverageIOBytes int `json:"AverageIOBytes,omitempty"`

	// Duration that this component is active.
	Duration string `json:"Duration,omitempty"`

	// Expected access pattern for this component.
	IOAccessPattern IOAccessPattern `json:"IOAccessPattern,omitempty"`

	// Percent of data for this workload component.
	PercentOfData int `json:"PercentOfData,omitempty"`

	// Percent of total IOPS for this workload component.
	PercentOfIOPS int `json:"PercentOfIOPS,omitempty"`

	// Specifies when to apply this workload component.
	Schedule map[string]interface{} `json:"Schedule,omitempty"`

}
