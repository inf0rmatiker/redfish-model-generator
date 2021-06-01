/* -----------------------------------------------------------------
* enum_nominal_voltage_type.go -
*
* DMTF Redfish NominalVoltageType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type NominalVoltageType string

const (
	// AC 100-240V nominal.
	NominalVoltageType_AC100_TO240_V NominalVoltageType = "AC100To240V"

	// AC 100-277V nominal.
	NominalVoltageType_AC100_TO277_V NominalVoltageType = "AC100To277V"

	// AC 120V nominal.
	NominalVoltageType_AC120_V NominalVoltageType = "AC120V"

	// AC 200-240V nominal.
	NominalVoltageType_AC200_TO240_V NominalVoltageType = "AC200To240V"

	// AC 200-277V nominal.
	NominalVoltageType_AC200_TO277_V NominalVoltageType = "AC200To277V"

	// AC 208V nominal.
	NominalVoltageType_AC208_V NominalVoltageType = "AC208V"

	// AC 230V nominal.
	NominalVoltageType_AC230_V NominalVoltageType = "AC230V"

	// AC 240V nominal.
	NominalVoltageType_AC240_V NominalVoltageType = "AC240V"

	// AC 200-240V and DC 380V.
	NominalVoltageType_AC240_AND_DC380_V NominalVoltageType = "AC240AndDC380V"

	// AC 277V nominal.
	NominalVoltageType_AC277_V NominalVoltageType = "AC277V"

	// AC 200-277V and DC 380V.
	NominalVoltageType_AC277_AND_DC380_V NominalVoltageType = "AC277AndDC380V"

	// AC 400V or 415V nominal.
	NominalVoltageType_AC400_V NominalVoltageType = "AC400V"

	// AC 480V nominal.
	NominalVoltageType_AC480_V NominalVoltageType = "AC480V"

	// DC 48V nominal.
	NominalVoltageType_DC48_V NominalVoltageType = "DC48V"

	// DC 240V nominal.
	NominalVoltageType_DC240_V NominalVoltageType = "DC240V"

	// High Voltage DC (380V).
	NominalVoltageType_DC380_V NominalVoltageType = "DC380V"

	// -48V DC.
	NominalVoltageType_DC_NEG48_V NominalVoltageType = "DCNeg48V"

)
