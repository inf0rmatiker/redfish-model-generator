/* -----------------------------------------------------------------
* enum_plug_type.go -
*
* DMTF Redfish PlugType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PlugType string

const (
	// NEMA 5-15P (Single-phase 125V; 15A; 1P3W).
	PlugType_NEMA_5_15_P PlugType = "NEMA_5_15P"

	// NEMA L5-15P (Single-phase 125V; 15A; 1P3W).
	PlugType_NEMA_L5_15_P PlugType = "NEMA_L5_15P"

	// NEMA 5-20P (Single-phase 125V; 20A; 1P3W).
	PlugType_NEMA_5_20_P PlugType = "NEMA_5_20P"

	// NEMA L5-20P (Single-phase 125V; 20A; 1P3W).
	PlugType_NEMA_L5_20_P PlugType = "NEMA_L5_20P"

	// NEMA L5-30P (Single-phase 125V; 30A; 1P3W).
	PlugType_NEMA_L5_30_P PlugType = "NEMA_L5_30P"

	// NEMA 6-15P (Single-phase 250V; 15A; 2P3W).
	PlugType_NEMA_6_15_P PlugType = "NEMA_6_15P"

	// NEMA L6-15P (Single-phase 250V; 15A; 2P3W).
	PlugType_NEMA_L6_15_P PlugType = "NEMA_L6_15P"

	// NEMA 6-20P (Single-phase 250V; 20A; 2P3W).
	PlugType_NEMA_6_20_P PlugType = "NEMA_6_20P"

	// NEMA L6-20P (Single-phase 250V; 20A; 2P3W).
	PlugType_NEMA_L6_20_P PlugType = "NEMA_L6_20P"

	// NEMA L6-30P (Single-phase 250V; 30A; 2P3W).
	PlugType_NEMA_L6_30_P PlugType = "NEMA_L6_30P"

	// NEMA L14-20P (Split-phase 125/250V; 20A; 2P4W).
	PlugType_NEMA_L14_20_P PlugType = "NEMA_L14_20P"

	// NEMA L14-30P (Split-phase 125/250V; 30A; 2P4W).
	PlugType_NEMA_L14_30_P PlugType = "NEMA_L14_30P"

	// NEMA L15-20P (Three-phase 250V; 20A; 3P4W).
	PlugType_NEMA_L15_20_P PlugType = "NEMA_L15_20P"

	// NEMA L15-30P (Three-phase 250V; 30A; 3P4W).
	PlugType_NEMA_L15_30_P PlugType = "NEMA_L15_30P"

	// NEMA L21-20P (Three-phase 120/208V; 20A; 3P5W).
	PlugType_NEMA_L21_20_P PlugType = "NEMA_L21_20P"

	// NEMA L21-30P (Three-phase 120/208V; 30A; 3P5W).
	PlugType_NEMA_L21_30_P PlugType = "NEMA_L21_30P"

	// NEMA L22-20P (Three-phase 277/480V; 20A; 3P5W).
	PlugType_NEMA_L22_20_P PlugType = "NEMA_L22_20P"

	// NEMA L22-30P (Three-phase 277/480V; 30A; 3P5W).
	PlugType_NEMA_L22_30_P PlugType = "NEMA_L22_30P"

	// California Standard CS8265 (Single-phase 250V; 50A; 2P3W).
	PlugType_CALIFORNIA_CS8265 PlugType = "California_CS8265"

	// California Standard CS8365 (Three-phase 250V; 50A; 3P4W).
	PlugType_CALIFORNIA_CS8365 PlugType = "California_CS8365"

	// IEC C14 (Single-phase 250V; 10A; 1P3W).
	PlugType_IEC_60320_C14 PlugType = "IEC_60320_C14"

	// IEC C20 (Single-phase 250V; 16A; 1P3W).
	PlugType_IEC_60320_C20 PlugType = "IEC_60320_C20"

	// IEC 60309 316P6 (Single-phase 200-250V; 16A; 1P3W; Blue, 6-hour).
	PlugType_IEC_60309_316_P6 PlugType = "IEC_60309_316P6"

	// IEC 60309 332P6 (Single-phase 200-250V; 32A; 1P3W; Blue, 6-hour).
	PlugType_IEC_60309_332_P6 PlugType = "IEC_60309_332P6"

	// IEC 60309 363P6 (Single-phase 200-250V; 63A; 1P3W; Blue, 6-hour).
	PlugType_IEC_60309_363_P6 PlugType = "IEC_60309_363P6"

	// IEC 60309 516P6 (Three-phase 200-240/346-415V; 16A; 3P5W; Red; 6-hour).
	PlugType_IEC_60309_516_P6 PlugType = "IEC_60309_516P6"

	// IEC 60309 532P6 (Three-phase 200-240/346-415V; 32A; 3P5W; Red; 6-hour).
	PlugType_IEC_60309_532_P6 PlugType = "IEC_60309_532P6"

	// IEC 60309 563P6 (Three-phase 200-240/346-415V; 63A; 3P5W; Red; 6-hour).
	PlugType_IEC_60309_563_P6 PlugType = "IEC_60309_563P6"

	// IEC 60309 460P9 (Three-phase 200-250V; 60A; 3P4W; Blue; 9-hour).
	PlugType_IEC_60309_460_P9 PlugType = "IEC_60309_460P9"

	// IEC 60309 560P9 (Three-phase 120-144/208-250V; 60A; 3P5W; Blue; 9-hour).
	PlugType_IEC_60309_560_P9 PlugType = "IEC_60309_560P9"

	// Field-wired; Three-phase 200-250V; 60A; 3P4W.
	PlugType_FIELD_208_V_3_P4_W_60_A PlugType = "Field_208V_3P4W_60A"

	// Field-wired; Three-phase 200-240/346-415V; 32A; 3P5W.
	PlugType_FIELD_400_V_3_P5_W_32_A PlugType = "Field_400V_3P5W_32A"

)
