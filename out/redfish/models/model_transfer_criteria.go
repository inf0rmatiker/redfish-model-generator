/* -----------------------------------------------------------------
* transfer_criteria.go -
*
* DMTF Redfish TransferCriteria resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// TransferCriteria - The criteria used to initiate a transfer for an automatic transfer switch.
// Modeled after DMTF Redfish schema TransferCriteria
type TransferCriteria struct {
	// The frequency in Hertz over the nominal value that satisfies a criterion for transfer.
	OverNominalFrequencyHz float64 `json:"OverNominalFrequencyHz,omitempty"`

	// The positive percentage of voltage RMS over the nominal value that satisfies a criterion for transfer.
	OverVoltageRMSPercentage float64 `json:"OverVoltageRMSPercentage,omitempty"`

	// The sensitivity to voltage waveform quality to satisfy the criterion for initiating a transfer.
	TransferSensitivity TransferSensitivityType `json:"TransferSensitivity,omitempty"`

	// The frequency in Hertz under the nominal value that satisfies a criterion for transfer.
	UnderNominalFrequencyHz float64 `json:"UnderNominalFrequencyHz,omitempty"`

	// The negative percentage of voltage RMS under the nominal value that satisfies a criterion for transfer.
	UnderVoltageRMSPercentage float64 `json:"UnderVoltageRMSPercentage,omitempty"`

}
