/* -----------------------------------------------------------------
 * enum_collection_time_scope.go -
 *
 * DMTF Redfish CollectionTimeScope enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type CollectionTimeScope string

const (
	// The corresponding metric values apply to a point in time.  On the corresponding metric value instances, the Timestamp property value in the metric report specifies the point in time.
	CollectionTimeScope_POINT CollectionTimeScope = "Point"

	// The corresponding metric values apply to a time interval.  On the corresponding metric value instances, the Timestamp property value in the metric report specifies the end of the time interval and the CollectionDuration property specifies its duration.
	CollectionTimeScope_INTERVAL CollectionTimeScope = "Interval"

	// The corresponding metric values apply to a time interval that began at the startup of the measured resource.  On the corresponding metric value instances, the Timestamp property value in the metric report shall specifies the end of the time interval.  The CollectionDuration property value specifies the duration between the startup of resource and timestamp.
	CollectionTimeScope_STARTUP_INTERVAL CollectionTimeScope = "StartupInterval"

)
