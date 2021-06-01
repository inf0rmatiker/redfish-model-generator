/* -----------------------------------------------------------------
* transfer_configuration.go -
*
* DMTF Redfish TransferConfiguration resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// TransferConfiguration - The configuration settings for an automatic transfer switch.
// Modeled after DMTF Redfish schema TransferConfiguration
type TransferConfiguration struct {
	// The mains circuit that is switched on and qualified to supply power to the output circuit.
	ActiveMainsId string `json:"ActiveMainsId,omitempty"`

	// Indicates if the qualified alternate mains circuit is automatically switched on when the preferred mains circuit becomes unqualified and is automatically switched off.
	AutoTransferEnabled bool `json:"AutoTransferEnabled,omitempty"`

	// Indicates if a make-before-break switching sequence of the mains circuits is permitted when they are both qualified and in synchronization.
	ClosedTransitionAllowed bool `json:"ClosedTransitionAllowed,omitempty"`

	// The time in seconds to wait for a closed transition to occur.
	ClosedTransitionTimeoutSeconds int `json:"ClosedTransitionTimeoutSeconds,omitempty"`

	// The preferred source for the mains circuit to this equipment.
	PreferredMainsId string `json:"PreferredMainsId,omitempty"`

	// The time in seconds to delay the automatic transfer from the alternate mains circuit back to the preferred mains circuit.
	RetransferDelaySeconds int `json:"RetransferDelaySeconds,omitempty"`

	// Indicates if the automatic transfer is permitted from the alternate mains circuit back to the preferred mains circuit after the preferred mains circuit is qualified again and the Retransfer Delay time has expired.
	RetransferEnabled bool `json:"RetransferEnabled,omitempty"`

	// The time in seconds to delay the automatic transfer from the preferred mains circuit to the alternate mains circuit when the preferred mains circuit is disqualified.
	TransferDelaySeconds int `json:"TransferDelaySeconds,omitempty"`

	// Indicates if any transfer is inhibited.
	TransferInhibit bool `json:"TransferInhibit,omitempty"`

}
