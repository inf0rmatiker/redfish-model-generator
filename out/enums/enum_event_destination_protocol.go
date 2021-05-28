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
	// The destination follows the Redfish Specification for event notifications.
	EventDestinationProtocol_REDFISH EventDestinationProtocol = "Redfish"

	// The destination follows the SNMPv1 protocol for event notifications.
	EventDestinationProtocol_SNM_PV1 EventDestinationProtocol = "SNMPv1"

	// The destination follows the SNMPv2c protocol for event notifications.
	EventDestinationProtocol_SNM_PV2C EventDestinationProtocol = "SNMPv2c"

	// The destination follows the SNMPv3 protocol for event notifications.
	EventDestinationProtocol_SNM_PV3 EventDestinationProtocol = "SNMPv3"

	// The destination follows the SMTP specification for event notifications.
	EventDestinationProtocol_SMTP EventDestinationProtocol = "SMTP"

	// The destination follows syslog TLS-based for event notifications.
	EventDestinationProtocol_SYSLOG_TLS EventDestinationProtocol = "SyslogTLS"

	// The destination follows syslog TCP-based for event notifications.
	EventDestinationProtocol_SYSLOG_TCP EventDestinationProtocol = "SyslogTCP"

	// The destination follows syslog UDP-based for event notifications.
	EventDestinationProtocol_SYSLOG_UDP EventDestinationProtocol = "SyslogUDP"

	// The destination follows syslog RELP for event notifications.
	EventDestinationProtocol_SYSLOG_RELP EventDestinationProtocol = "SyslogRELP"

	// The destination follows an OEM protocol for event notifications.
	EventDestinationProtocol_OEM EventDestinationProtocol = "OEM"

)
