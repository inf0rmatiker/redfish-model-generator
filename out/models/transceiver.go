/* -----------------------------------------------------------------
* transceiver.go -
*
* DMTF Redfish Transceiver resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The transceiver metrics.
type Transceiver struct {
	// The RX input power value of a small form-factor pluggable (SFP) transceiver.
	RXInputPowerMilliWatts float64 `json:"RXInputPowerMilliWatts,omitempty"`

	// The supply voltage of a small form-factor pluggable (SFP) transceiver.
	SupplyVoltage float64 `json:"SupplyVoltage,omitempty"`

	// The TX bias current value of a small form-factor pluggable (SFP) transceiver.
	TXBiasCurrentMilliAmps float64 `json:"TXBiasCurrentMilliAmps,omitempty"`

	// The TX output power value of a small form-factor pluggable (SFP) transceiver.
	TXOutputPowerMilliWatts float64 `json:"TXOutputPowerMilliWatts,omitempty"`

}
