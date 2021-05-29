/* -----------------------------------------------------------------
* schedule.go -
*
* DMTF Redfish Schedule resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Schedule a series of occurrences.
type Schedule struct {
	// Days of the month when scheduled occurrences are enabled.  `0` indicates that every day of the month is enabled.
	EnabledDaysOfMonth []integer `json:"EnabledDaysOfMonth,omitempty"`

	// Days of the week when scheduled occurrences are enabled, for enabled days of the month and months of the year.  If not present, all days of the week are enabled.
	EnabledDaysOfWeek array `json:"EnabledDaysOfWeek,omitempty"`

	// Intervals when scheduled occurrences are enabled.
	EnabledIntervals []string `json:"EnabledIntervals,omitempty"`

	// The months of the year when scheduled occurrences are enabled.  If not present, all months of the year are enabled.
	EnabledMonthsOfYear array `json:"EnabledMonthsOfYear,omitempty"`

	// The date and time when the initial occurrence is scheduled to occur.
	InitialStartTime string `json:"InitialStartTime,omitempty"`

	// The time after provisioning when the schedule as a whole expires.
	Lifetime string `json:"Lifetime,omitempty"`

	// The maximum number of scheduled occurrences.
	MaxOccurrences int `json:"MaxOccurrences,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The amount of time until the next occurrence occurs.
	RecurrenceInterval string `json:"RecurrenceInterval,omitempty"`

}
