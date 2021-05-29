/* -----------------------------------------------------------------
* sse_filter_properties_supported.go -
*
* DMTF Redfish SSEFilterPropertiesSupported resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The set of properties that are supported in the `$filter` query parameter for the ServerSentEventUri.
type SSEFilterPropertiesSupported struct {
	// An indication of whether the service supports filtering by the EventFormatType property.
	EventFormatType bool `json:"EventFormatType,omitempty"`

	// An indication of whether the service supports filtering by the EventType property.
	EventType bool `json:"EventType,omitempty"`

	// An indication of whether the service supports filtering by the MessageId property.
	MessageId bool `json:"MessageId,omitempty"`

	// An indication of whether the service supports filtering by the MetricReportDefinition property.
	MetricReportDefinition bool `json:"MetricReportDefinition,omitempty"`

	// An indication of whether the service supports filtering by the OriginResource property.
	OriginResource bool `json:"OriginResource,omitempty"`

	// An indication of whether the service supports filtering by the RegistryPrefix property.
	RegistryPrefix bool `json:"RegistryPrefix,omitempty"`

	// An indication of whether the service supports filtering by the ResourceType property.
	ResourceType bool `json:"ResourceType,omitempty"`

	// An indication of whether the service supports filtering by the SubordinateResources property.
	SubordinateResources bool `json:"SubordinateResources,omitempty"`

}
