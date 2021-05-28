/* -----------------------------------------------------------------
 * enum_day_of_week.go -
 *
 * DMTF Redfish DayOfWeek enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DayOfWeek string

const (
	// Monday.
	DayOfWeek_MONDAY DayOfWeek = "Monday"

	// Tuesday.
	DayOfWeek_TUESDAY DayOfWeek = "Tuesday"

	// Wednesday.
	DayOfWeek_WEDNESDAY DayOfWeek = "Wednesday"

	// Thursday.
	DayOfWeek_THURSDAY DayOfWeek = "Thursday"

	// Friday.
	DayOfWeek_FRIDAY DayOfWeek = "Friday"

	// Saturday.
	DayOfWeek_SATURDAY DayOfWeek = "Saturday"

	// Sunday.
	DayOfWeek_SUNDAY DayOfWeek = "Sunday"

	// Every day of the week.
	DayOfWeek_EVERY DayOfWeek = "Every"

)
