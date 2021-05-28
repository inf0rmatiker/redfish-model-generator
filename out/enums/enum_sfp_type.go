/* -----------------------------------------------------------------
 * enum_sfp_type.go -
 *
 * DMTF Redfish SFPType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SFPType string

const (
	// The SFP conforms to the SFF Specification for SFP.
	SFPType_SFP SFPType = "SFP"

	// The SFP conforms to the SFF Specification for SFP+.
	SFPType_SFP_PLUS SFPType = "SFPPlus"

	// The SFP conforms to the SFF Specification for SFP+ and IEEE 802.3by Specification.
	SFPType_SFP28 SFPType = "SFP28"

	// The SFP conforms to the CSFP MSA Specification.
	SFPType_C_SFP SFPType = "cSFP"

	// The SFP conforms to the SFP-DD MSA Specification.
	SFPType_SFPDD SFPType = "SFPDD"

	// The SFP conforms to the SFF Specification for QSFP.
	SFPType_QSFP SFPType = "QSFP"

	// The SFP conforms to the SFF Specification for QSFP+.
	SFPType_QSFP_PLUS SFPType = "QSFPPlus"

	// The SFP conforms to the SFF Specification for QSFP14.
	SFPType_QSFP14 SFPType = "QSFP14"

	// The SFP conforms to the SFF Specification for QSFP28.
	SFPType_QSFP28 SFPType = "QSFP28"

	// The SFP conforms to the SFF Specification for QSFP56.
	SFPType_QSFP56 SFPType = "QSFP56"

	// The SFP conforms to the SFF Specification SFF-8644.
	SFPType_MINI_SASHD SFPType = "MiniSASHD"

)
