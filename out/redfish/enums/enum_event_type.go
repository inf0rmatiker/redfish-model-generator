/* -----------------------------------------------------------------
* enum_event_type.go -
*
* DMTF Redfish EventType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EventType string

const (
	// The status of a resource has changed.
	EventType_STATUS_CHANGE EventType = "StatusChange"

	// A resource has been updated.
	EventType_RESOURCE_UPDATED EventType = "ResourceUpdated"

	// A resource has been added.
	EventType_RESOURCE_ADDED EventType = "ResourceAdded"

	// A resource has been removed.
	EventType_RESOURCE_REMOVED EventType = "ResourceRemoved"

	// A condition requires attention.
	EventType_ALERT EventType = "Alert"

	// The telemetry service is sending a metric report.
	EventType_METRIC_REPORT EventType = "MetricReport"

	// Because EventType is deprecated as of Redfish Specification v1.6, the event is based on a registry or resource but not an EventType.
	EventType_OTHER EventType = "Other"

)
