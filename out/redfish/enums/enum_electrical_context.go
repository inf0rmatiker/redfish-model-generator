/* -----------------------------------------------------------------
* enum_electrical_context.go -
*
* DMTF Redfish ElectricalContext enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ElectricalContext string

const (
	// The circuits that share the L1 current-carrying conductor.
	ElectricalContext_LINE1 ElectricalContext = "Line1"

	// The circuits that share the L2 current-carrying conductor when PhaseWiringType.ThreePhase4Wire, TwoPhase4Wire, or ThreePhase5Wire.
	ElectricalContext_LINE2 ElectricalContext = "Line2"

	// The circuits that share the L3 current-carrying conductor when PhaseWiringType.ThreePhase4Wire or ThreePhase5Wire.
	ElectricalContext_LINE3 ElectricalContext = "Line3"

	// The grounded current-carrying return circuit of current-carrying conductors when PhaseWiringType.OnePhase3Wire, TwoPhase4Wire, or ThreePhase5Wire.
	ElectricalContext_NEUTRAL ElectricalContext = "Neutral"

	// The circuit formed by two current-carrying conductors when PhaseWiringType.TwoPhase3Wire, TwoPhase4Wire, ThreePhase4Wire, or ThreePhase5Wire.
	ElectricalContext_LINE_TO_LINE ElectricalContext = "LineToLine"

	// The circuit formed by L1 and L2 current-carrying conductors when PhaseWiringType.TwoPhase3Wire, TwoPhase4Wire, ThreePhase4Wire, or ThreePhase5Wire.
	ElectricalContext_LINE1_TO_LINE2 ElectricalContext = "Line1ToLine2"

	// The circuit formed by L2 and L3 current-carrying conductors when PhaseWiringType.ThreePhase4Wire or ThreePhase5Wire.
	ElectricalContext_LINE2_TO_LINE3 ElectricalContext = "Line2ToLine3"

	// The circuit formed by L3 and L1 current-carrying conductors when PhaseWiringType.ThreePhase4Wire or ThreePhase5Wire.
	ElectricalContext_LINE3_TO_LINE1 ElectricalContext = "Line3ToLine1"

	// The circuit formed by a line and Neutral current-carrying conductor when PhaseWiringType.OnePhase3Wire, TwoPhase4Wire, ThreePhase4Wire, or ThreePhase5Wire.
	ElectricalContext_LINE_TO_NEUTRAL ElectricalContext = "LineToNeutral"

	// The circuit formed by L1 and Neutral current-carrying conductors when PhaseWiringType.OnePhase3Wire, TwoPhase4Wire, ThreePhase4Wire, or ThreePhase5Wire.
	ElectricalContext_LINE1_TO_NEUTRAL ElectricalContext = "Line1ToNeutral"

	// The circuit formed by L2 and Neutral current-carrying conductors when PhaseWiringType.TwoPhase4Wire or ThreePhase5Wire.
	ElectricalContext_LINE2_TO_NEUTRAL ElectricalContext = "Line2ToNeutral"

	// The circuit formed by L3 and Neutral current-carrying conductors when PhaseWiringType.ThreePhase5Wire.
	ElectricalContext_LINE3_TO_NEUTRAL ElectricalContext = "Line3ToNeutral"

	// The circuits formed by L1, L2, and Neutral current-carrying conductors when PhaseWiringType.TwoPhase4Wire or ThreePhase5Wire.
	ElectricalContext_LINE1_TO_NEUTRAL_AND_L1_L2 ElectricalContext = "Line1ToNeutralAndL1L2"

	// The circuits formed by L1, L2, and Neutral current-carrying conductors when PhaseWiringType.TwoPhase4Wire or ThreePhase5Wire.
	ElectricalContext_LINE2_TO_NEUTRAL_AND_L1_L2 ElectricalContext = "Line2ToNeutralAndL1L2"

	// The circuits formed by L2, L3, and Neutral current-carrying conductors when PhaseWiringType.ThreePhase5Wire.
	ElectricalContext_LINE2_TO_NEUTRAL_AND_L2_L3 ElectricalContext = "Line2ToNeutralAndL2L3"

	// The circuits formed by L3, L1, and Neutral current-carrying conductors when PhaseWiringType.ThreePhase5Wire.
	ElectricalContext_LINE3_TO_NEUTRAL_AND_L3_L1 ElectricalContext = "Line3ToNeutralAndL3L1"

	// The circuits formed by all current-carrying conductors for any PhaseWiringType.
	ElectricalContext_TOTAL ElectricalContext = "Total"

)
