/* -----------------------------------------------------------------
* fan.go -
*
* DMTF Redfish Fan resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Fan - This is the base type for addressable members of an array.
// Modeled after DMTF Redfish schema Fan
type Fan struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId,omitempty"`

	// Name of the fan.
	FanName string `json:"FanName,omitempty"`

	// Describes the area or device associated with this fan.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// Current fan speed.
	Reading float64 `json:"Reading,omitempty"`

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

	// Minimum value for Reading.
	MinReadingRange float64 `json:"MinReadingRange,omitempty"`

	// Maximum value for Reading.
	MaxReadingRange float64 `json:"MaxReadingRange,omitempty"`

	RelatedItemOdataCount int `json:"RelatedItem@odata.count,omitempty"`

	RelatedItem@odata.navigationLink string `json:"RelatedItem@odata.navigationLink,omitempty"`

	// The ID(s) of the resources serviced with this fan.
	RelatedItem []idRef `json:"RelatedItem,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	Redundancy@odata.navigationLink string `json:"Redundancy@odata.navigationLink,omitempty"`

	// This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	// Units in which the reading and thresholds are measured.
	ReadingUnits ReadingUnits `json:"ReadingUnits,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// This is the manufacturer of this Fan.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model number for this Fan.
	Model string `json:"Model,omitempty"`

	// The serial number for this Fan.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The part number for this Fan.
	PartNumber string `json:"PartNumber,omitempty"`

	// The spare part number for this Fan.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The state of the indicator LED, used to identify this Fan.
	IndicatorLED map[string]interface{} `json:"IndicatorLED,omitempty"`

}
