/* -----------------------------------------------------------------
* schedule.go -
*
* DMTF Redfish Schedule resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Schedule - Schedule a series of occurrences.
// Modeled after DMTF Redfish schema Schedule
type Schedule struct {
	// Days of month when scheduled occurrences are enabled. Zero indicates that every day of the month is enabled.
	EnabledDaysOfMonth []int `json:"EnabledDaysOfMonth,omitempty"`

	// Days of the week when scheduled occurrences are enabled, for enabled days of month and months of year.
	EnabledDaysOfWeek []DayOfWeek `json:"EnabledDaysOfWeek,omitempty"`

	// Intervals when scheduled occurrences are enabled.
	EnabledIntervals []string `json:"EnabledIntervals,omitempty"`

	// Months of year when scheduled occurrences are enabled.
	EnabledMonthsOfYear []MonthOfYear `json:"EnabledMonthsOfYear,omitempty"`

	// Time for initial occurrence.
	InitialStartTime string `json:"InitialStartTime,omitempty"`

	// The time after provisioning when the schedule as a whole expires.
	Lifetime string `json:"Lifetime,omitempty"`

	// Maximum number of scheduled occurrences.
	MaxOccurrences int `json:"MaxOccurrences,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// Distance until the next occurrences.
	RecurrenceInterval string `json:"RecurrenceInterval,omitempty"`

}
