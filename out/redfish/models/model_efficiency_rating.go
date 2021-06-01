/* -----------------------------------------------------------------
* efficiency_rating.go -
*
* DMTF Redfish EfficiencyRating resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EfficiencyRating - This type describes an efficiency rating for a power supply.
// Modeled after DMTF Redfish schema EfficiencyRating
type EfficiencyRating struct {
	// The rated efficiency of this power supply at the specified load.
	EfficiencyPercent float64 `json:"EfficiencyPercent,omitempty"`

	// The electrical load for this rating.
	LoadPercent float64 `json:"LoadPercent,omitempty"`

}
