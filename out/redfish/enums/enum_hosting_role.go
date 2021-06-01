/* -----------------------------------------------------------------
* enum_hosting_role.go -
*
* DMTF Redfish HostingRole enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type HostingRole string

const (
	// The system hosts functionality that supports general purpose applications.
	HostingRole_APPLICATION_SERVER HostingRole = "ApplicationServer"

	// The system hosts functionality that supports the system acting as a storage server.
	HostingRole_STORAGE_SERVER HostingRole = "StorageServer"

	// The system hosts functionality that supports the system acting as a switch.
	HostingRole_SWITCH HostingRole = "Switch"

)
