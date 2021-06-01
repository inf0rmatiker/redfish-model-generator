/* -----------------------------------------------------------------
* voltage.go -
*
* DMTF Redfish Voltage resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Voltage - This is the base type for addressable members of an array.
// Modeled after DMTF Redfish schema Voltage
type Voltage struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// A numerical identifier to represent the voltage sensor.
	SensorNumber float64 `json:"SensorNumber,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// The present reading of the voltage sensor.
	ReadingVolts float64 `json:"ReadingVolts,omitempty"`

	// Above normal range.
	UpperThresholdNonCritical float64 `json:"UpperThresholdNonCritical,omitempty"`

	// Above normal range but not yet fatal.
	UpperThresholdCritical float64 `json:"UpperThresholdCritical,omitempty"`

	// Above normal range and is fatal.
	UpperThresholdFatal float64 `json:"UpperThresholdFatal,omitempty"`

	// Below normal range.
	LowerThresholdNonCritical float64 `json:"LowerThresholdNonCritical,omitempty"`

	// Below normal range but not yet fatal.
	LowerThresholdCritical float64 `json:"LowerThresholdCritical,omitempty"`

	// Below normal range and is fatal.
	LowerThresholdFatal float64 `json:"LowerThresholdFatal,omitempty"`

	// Minimum value for this Voltage sensor.
	MinReadingRange float64 `json:"MinReadingRange,omitempty"`

	// Maximum value for this Voltage sensor.
	MaxReadingRange float64 `json:"MaxReadingRange,omitempty"`

	// Describes the area or device to which this voltage measurement applies.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

	RelatedItemOdataCount int `json:"RelatedItem@odata.count,omitempty"`

	RelatedItem@odata.navigationLink string `json:"RelatedItem@odata.navigationLink,omitempty"`

	// Describes the areas or devices to which this voltage measurement applies.
	RelatedItem []idRef `json:"RelatedItem,omitempty"`

}
