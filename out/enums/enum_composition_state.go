/* -----------------------------------------------------------------
 * enum_composition_state.go -
 *
 * DMTF Redfish CompositionState enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type CompositionState string

const (
	// Intermediate state indicating composition is in progress.
	CompositionState_COMPOSING CompositionState = "Composing"

	// The resource block is currently participating in one or more compositions, and is available to use in more compositions.
	CompositionState_COMPOSED_AND_AVAILABLE CompositionState = "ComposedAndAvailable"

	// Final successful state of a resource block that has participated in composition.
	CompositionState_COMPOSED CompositionState = "Composed"

	// The resource block is free and can participate in composition.
	CompositionState_UNUSED CompositionState = "Unused"

	// The final composition resulted in failure and manual intervention might be required to fix it.
	CompositionState_FAILED CompositionState = "Failed"

	// The resource block has been made unavailable by the service, such as due to maintenance being performed on the resource block.
	CompositionState_UNAVAILABLE CompositionState = "Unavailable"

)
