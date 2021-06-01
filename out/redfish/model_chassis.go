/* -----------------------------------------------------------------
* chassis.go -
*
* DMTF Redfish Chassis resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Chassis - A Chassis represents the physical components for any system.  This resource represents the sheet-metal confined spaces and logical zones like racks, enclosures, chassis and all other containers. Subsystems (like sensors), which operate outside of a system's data plane (meaning the resources are not accessible to software running on the system) are linked either directly or indirectly through this resource.
// Modeled after DMTF Redfish schema Chassis
type Chassis struct {
}
