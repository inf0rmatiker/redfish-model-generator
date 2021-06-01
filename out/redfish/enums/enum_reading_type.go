/* -----------------------------------------------------------------
* enum_reading_type.go -
*
* DMTF Redfish ReadingType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReadingType string

const (
	// Temperature.
	ReadingType_TEMPERATURE ReadingType = "Temperature"

	// Relative Humidity.
	ReadingType_HUMIDITY ReadingType = "Humidity"

	// Power.
	ReadingType_POWER ReadingType = "Power"

	// Energy (kWh).
	ReadingType_ENERGYK_WH ReadingType = "EnergykWh"

	// Energy (Joules).
	ReadingType_ENERGY_JOULES ReadingType = "EnergyJoules"

	// Voltage (AC or DC).
	ReadingType_VOLTAGE ReadingType = "Voltage"

	// Current.
	ReadingType_CURRENT ReadingType = "Current"

	// Frequency.
	ReadingType_FREQUENCY ReadingType = "Frequency"

	// Pressure.
	ReadingType_PRESSURE ReadingType = "Pressure"

	// Liquid level.
	ReadingType_LIQUID_LEVEL ReadingType = "LiquidLevel"

	// Rotational.
	ReadingType_ROTATIONAL ReadingType = "Rotational"

	// Airflow.
	ReadingType_AIR_FLOW ReadingType = "AirFlow"

	// Liquid flow.
	ReadingType_LIQUID_FLOW ReadingType = "LiquidFlow"

	// Barometric pressure.
	ReadingType_BAROMETRIC ReadingType = "Barometric"

	// Altitude.
	ReadingType_ALTITUDE ReadingType = "Altitude"

)
