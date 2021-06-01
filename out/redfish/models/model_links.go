/* -----------------------------------------------------------------
* links.go -
*
* DMTF Redfish Links resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Links - Contains references to other resources that are related to this resource.
// Modeled after DMTF Redfish schema Links
type Links struct {
	// This property is an array of references to the chassis that this manager has control over.
	ManagerForChassis []Chassis `json:"ManagerForChassis,omitempty"`

	ManagerForChassisOdataCount int `json:"ManagerForChassis@odata.count,omitempty"`

	// This property is an array of references to the systems that this manager has control over.
	ManagerForServers []ComputerSystem `json:"ManagerForServers,omitempty"`

	ManagerForServersOdataCount int `json:"ManagerForServers@odata.count,omitempty"`

	// This property is an array of references to the switches that this manager has control over.
	ManagerForSwitches []Switch `json:"ManagerForSwitches,omitempty"`

	ManagerForSwitchesOdataCount int `json:"ManagerForSwitches@odata.count,omitempty"`

	// This property is a reference to the chassis that this manager is located in.
	ManagerInChassis map[string]interface{} `json:"ManagerInChassis,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
