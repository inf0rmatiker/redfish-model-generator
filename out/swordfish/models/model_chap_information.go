/* -----------------------------------------------------------------
* chap_information.go -
*
* SNIA Swordfish CHAPInformation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CHAPInformation - User name and password for CHAP authentication.
// Modeled after SNIA Swordfish schema CHAPInformation
type CHAPInformation struct {
	// The shared secret for Mutual (2-way) CHAP authentication by the initiator.
	InitiatorCHAPPassword string `json:"InitiatorCHAPPassword,omitempty"`

	// The Initiator username for Mutual (2-way) CHAP authentication by the initiator.
	InitiatorCHAPUser string `json:"InitiatorCHAPUser,omitempty"`

	// The Target CHAP Username for Mutual (2-way) CHAP authentication by the target.
	TargetCHAPUser string `json:"TargetCHAPUser,omitempty"`

	// This property is deprecated in favor of TargetCHAPPassword.
	TargetPassword string `json:"TargetPassword,omitempty"`

}
