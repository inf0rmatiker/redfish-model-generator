/* -----------------------------------------------------------------
* enum_phase_wiring_type.go -
*
* DMTF Redfish PhaseWiringType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PhaseWiringType string

const (
	// Single-phase / 3-Wire (Line1, Neutral, Protective Earth).
	PhaseWiringType_ONE_PHASE3_WIRE PhaseWiringType = "OnePhase3Wire"

	// Two-phase / 3-Wire (Line1, Line2, Protective Earth).
	PhaseWiringType_TWO_PHASE3_WIRE PhaseWiringType = "TwoPhase3Wire"

	// Single or Two-Phase / 3-Wire (Line1, Line2 or Neutral, Protective Earth).
	PhaseWiringType_ONE_OR_TWO_PHASE3_WIRE PhaseWiringType = "OneOrTwoPhase3Wire"

	// Two-phase / 4-Wire (Line1, Line2, Neutral, Protective Earth).
	PhaseWiringType_TWO_PHASE4_WIRE PhaseWiringType = "TwoPhase4Wire"

	// Three-phase / 4-Wire (Line1, Line2, Line3, Protective Earth).
	PhaseWiringType_THREE_PHASE4_WIRE PhaseWiringType = "ThreePhase4Wire"

	// Three-phase / 5-Wire (Line1, Line2, Line3, Neutral, Protective Earth).
	PhaseWiringType_THREE_PHASE5_WIRE PhaseWiringType = "ThreePhase5Wire"

)
