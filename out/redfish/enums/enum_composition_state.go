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

	// Indicates the Resource Block is currently participating in one or more compositions, and is available to be used in more compositions.
	CompositionState_COMPOSED_AND_AVAILABLE CompositionState = "ComposedAndAvailable"

	// Final successful state of a Resource Block which has participated in composition.
	CompositionState_COMPOSED CompositionState = "Composed"

	// Indicates the Resource Block is free and can participate in composition.
	CompositionState_UNUSED CompositionState = "Unused"

	// The final composition resulted in failure and manual intervention may be required to fix it.
	CompositionState_FAILED CompositionState = "Failed"

	// Indicates the Resource Block has been made unavailable by the service, such as due to maintenance being performed on the Resource Block.
	CompositionState_UNAVAILABLE CompositionState = "Unavailable"

)
