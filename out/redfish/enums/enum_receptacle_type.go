/* -----------------------------------------------------------------
* enum_receptacle_type.go -
*
* DMTF Redfish ReceptacleType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReceptacleType string

const (
	// NEMA 5-15R (120V; 15A).
	ReceptacleType_NEMA_5_15_R ReceptacleType = "NEMA_5_15R"

	// NEMA 5-20R (120V; 20A).
	ReceptacleType_NEMA_5_20_R ReceptacleType = "NEMA_5_20R"

	// NEMA L5-20R (120V; 20A).
	ReceptacleType_NEMA_L5_20_R ReceptacleType = "NEMA_L5_20R"

	// NEMA L5-30R (120V; 30A).
	ReceptacleType_NEMA_L5_30_R ReceptacleType = "NEMA_L5_30R"

	// NEMA L6-20R (250V; 20A).
	ReceptacleType_NEMA_L6_20_R ReceptacleType = "NEMA_L6_20R"

	// NEMA L6-30R (250V; 30A).
	ReceptacleType_NEMA_L6_30_R ReceptacleType = "NEMA_L6_30R"

	// IEC C13 (250V; 10A or 15A).
	ReceptacleType_IEC_60320_C13 ReceptacleType = "IEC_60320_C13"

	// IEC C19 (250V; 16A or 20A).
	ReceptacleType_IEC_60320_C19 ReceptacleType = "IEC_60320_C19"

	// CEE 7/7 Type E (250V; 16A).
	ReceptacleType_CEE_7_TYPE_E ReceptacleType = "CEE_7_Type_E"

	// CEE 7/7 Type F (250V; 16A).
	ReceptacleType_CEE_7_TYPE_F ReceptacleType = "CEE_7_Type_F"

	// SEV 1011 Type 12 (250V; 10A).
	ReceptacleType_SEV_1011_TYPE_12 ReceptacleType = "SEV_1011_TYPE_12"

	// SEV 1011 Type 23 (250V; 16A).
	ReceptacleType_SEV_1011_TYPE_23 ReceptacleType = "SEV_1011_TYPE_23"

	// BS 1363 Type G (250V; 13A).
	ReceptacleType_BS_1363_TYPE_G ReceptacleType = "BS_1363_Type_G"

)
