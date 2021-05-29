/* -----------------------------------------------------------------
* temperature.go -
*
* DMTF Redfish Temperature resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Temperature struct {
	OdataId string `json:"@odata.id"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Adjusted maximum allowable operating temperature for this equipment based on the current environmental conditions present.
	AdjustedMaxAllowableOperatingValue int `json:"AdjustedMaxAllowableOperatingValue,omitempty"`

	// Adjusted minimum allowable operating temperature for this equipment based on the current environmental conditions present.
	AdjustedMinAllowableOperatingValue int `json:"AdjustedMinAllowableOperatingValue,omitempty"`

	// The area or device to which the DeltaReadingCelsius temperature measurement applies, relative to PhysicalContext.
	DeltaPhysicalContext PhysicalContext `json:"DeltaPhysicalContext,omitempty"`

	// The delta temperature reading.
	DeltaReadingCelsius float64 `json:"DeltaReadingCelsius,omitempty"`

	// The value at which the reading is below normal range but not yet fatal.
	LowerThresholdCritical float64 `json:"LowerThresholdCritical,omitempty"`

	// The value at which the reading is below normal range and fatal.
	LowerThresholdFatal float64 `json:"LowerThresholdFatal,omitempty"`

	// The value at which the reading is below normal range.
	LowerThresholdNonCritical float64 `json:"LowerThresholdNonCritical,omitempty"`

	// The value at which the reading is below the user-defined range.
	LowerThresholdUser int `json:"LowerThresholdUser,omitempty"`

	// Maximum allowable operating temperature for this equipment.
	MaxAllowableOperatingValue int `json:"MaxAllowableOperatingValue,omitempty"`

	// Maximum value for this sensor.
	MaxReadingRangeTemp float64 `json:"MaxReadingRangeTemp,omitempty"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// Minimum allowable operating temperature for this equipment.
	MinAllowableOperatingValue int `json:"MinAllowableOperatingValue,omitempty"`

	// Minimum value for this sensor.
	MinReadingRangeTemp float64 `json:"MinReadingRangeTemp,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The area or device to which this temperature measurement applies.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// The temperature in degrees Celsius.
	ReadingCelsius float64 `json:"ReadingCelsius,omitempty"`

	// An array of links to resources or objects that represent areas or devices to which this temperature applies.
	RelatedItem []map[string]string `json:"RelatedItem,omitempty"`

	RelatedItem@odata.count count `json:"RelatedItem@odata.count,omitempty"`

	// The numerical identifier of the temperature sensor.
	SensorNumber int `json:"SensorNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The value at which the reading is above normal range but not yet fatal.
	UpperThresholdCritical float64 `json:"UpperThresholdCritical,omitempty"`

	// The value at which the reading is above normal range and fatal.
	UpperThresholdFatal float64 `json:"UpperThresholdFatal,omitempty"`

	// The value at which the reading is above normal range.
	UpperThresholdNonCritical float64 `json:"UpperThresholdNonCritical,omitempty"`

	// The value at which the reading is above the user-defined range.
	UpperThresholdUser int `json:"UpperThresholdUser,omitempty"`

}
