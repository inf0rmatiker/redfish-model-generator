/* -----------------------------------------------------------------
* enum_application_consistency_method.go -
*
* SNIA Swordfish ApplicationConsistencyMethod enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ApplicationConsistencyMethod string

const (
	// Supports consistency method commonly orchestrated using application-specific code.
	ApplicationConsistencyMethod_HOT_STANDBY ApplicationConsistencyMethod = "HotStandby"

	// Supports VMware consistency requirements, such as for VASA and VVOLs.
	ApplicationConsistencyMethod_VASA ApplicationConsistencyMethod = "VASA"

	// Supports Microsoft virtual backup device interface (VDI).
	ApplicationConsistencyMethod_VDI ApplicationConsistencyMethod = "VDI"

	// Supports Microsoft VSS.
	ApplicationConsistencyMethod_VSS ApplicationConsistencyMethod = "VSS"

	// Supports consistency method orchestrated using vendor-specific code.
	ApplicationConsistencyMethod_OTHER ApplicationConsistencyMethod = "Other"

)
