/* -----------------------------------------------------------------
* enum_event_destination_protocol.go -
*
* DMTF Redfish EventDestinationProtocol enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EventDestinationProtocol string

const (
	// The destination follows the Redfish specification for event notifications.
	EventDestinationProtocol_REDFISH EventDestinationProtocol = "Redfish"

)
