/* -----------------------------------------------------------------
* bios.go -
*
* DMTF Redfish Bios resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Bios - The Bios schema contains properties related to the BIOS Attribute Registry.  The Attribute Registry describes the system-specific BIOS attributes and Actions for changing to BIOS settings.  Changes to the BIOS typically require a system reset before they take effect.  It is likely that a client will find the @Redfish.Settings term in this resource, and if it is found, the client makes requests to change BIOS settings by modifying the resource identified by the @Redfish.Settings term.
// Modeled after DMTF Redfish schema Bios
type Bios struct {
}
