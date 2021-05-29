/* -----------------------------------------------------------------
* voltage.go -
*
* DMTF Redfish Voltage resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Voltage struct {
	OdataId string `json:"@odata.id"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The value at which the reading is below normal range but not yet fatal.
	LowerThresholdCritical float64 `json:"LowerThresholdCritical,omitempty"`

	// The value at which the reading is below normal range and fatal.
	LowerThresholdFatal float64 `json:"LowerThresholdFatal,omitempty"`

	// The value at which the reading is below normal range.
	LowerThresholdNonCritical float64 `json:"LowerThresholdNonCritical,omitempty"`

	// Maximum value for this sensor.
	MaxReadingRange float64 `json:"MaxReadingRange,omitempty"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// Minimum value for this sensor.
	MinReadingRange float64 `json:"MinReadingRange,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The area or device to which this voltage measurement applies.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// The reading of the voltage sensor.
	ReadingVolts float64 `json:"ReadingVolts,omitempty"`

	// An array of links to resources or objects to which this voltage measurement applies.
	RelatedItem []map[string]string `json:"RelatedItem,omitempty"`

	RelatedItem@odata.count count `json:"RelatedItem@odata.count,omitempty"`

	// A numerical identifier to represent the voltage sensor.
	SensorNumber int `json:"SensorNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The value at which the reading is above normal range but not yet fatal.
	UpperThresholdCritical float64 `json:"UpperThresholdCritical,omitempty"`

	// The value at which the reading is above normal range and fatal.
	UpperThresholdFatal float64 `json:"UpperThresholdFatal,omitempty"`

	// The value at which the reading is above normal range.
	UpperThresholdNonCritical float64 `json:"UpperThresholdNonCritical,omitempty"`

}
