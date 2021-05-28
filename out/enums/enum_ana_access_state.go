/* -----------------------------------------------------------------
 * enum_ana_access_state.go -
 *
 * DMTF Redfish ANAAccessState enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ANAAccessState string

const (
	// Commands processed by a controller provide optimized access to any namespace in the ANA group.
	ANAAccessState_OPTIMIZED ANAAccessState = "Optimized"

	// Commands processed by a controller that reports this state for an ANA Group provide non-optimized access characteristics, such as lower performance or non-optimal use of subsystem resources, to any namespace in the ANA Group.
	ANAAccessState_NON_OPTIMIZED ANAAccessState = "NonOptimized"

	// Namespaces in this group are inaccessible.  Commands are not able to access user data of namespaces in the ANA Group.
	ANAAccessState_INACCESSIBLE ANAAccessState = "Inaccessible"

	// The group is persistently inaccessible.  Commands are persistently not able to access user data of namespaces in the ANA Group.
	ANAAccessState_PERSISTENT_LOSS ANAAccessState = "PersistentLoss"

)
