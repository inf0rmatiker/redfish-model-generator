/* -----------------------------------------------------------------
 * enum_map_from_property.go -
 *
 * DMTF Redfish MapFromProperty enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MapFromProperty string

const (
	// The dependency on an attribute's CurrentValue.
	MapFromProperty_CURRENT_VALUE MapFromProperty = "CurrentValue"

	// The dependency on an attribute's DefaultValue.
	MapFromProperty_DEFAULT_VALUE MapFromProperty = "DefaultValue"

	// The dependency on an attribute's ReadOnly state.
	MapFromProperty_READ_ONLY MapFromProperty = "ReadOnly"

	// The dependency on an attribute's WriteOnly state.
	MapFromProperty_WRITE_ONLY MapFromProperty = "WriteOnly"

	// The dependency on an attribute's GrayOut state.
	MapFromProperty_GRAY_OUT MapFromProperty = "GrayOut"

	// The dependency on an attribute's Hidden state.
	MapFromProperty_HIDDEN MapFromProperty = "Hidden"

	// The dependency on an attribute's LowerBound.
	MapFromProperty_LOWER_BOUND MapFromProperty = "LowerBound"

	// The dependency on an attribute's UpperBound.
	MapFromProperty_UPPER_BOUND MapFromProperty = "UpperBound"

	// The dependency on an attribute's MinLength.
	MapFromProperty_MIN_LENGTH MapFromProperty = "MinLength"

	// The dependency on an attribute's MaxLength.
	MapFromProperty_MAX_LENGTH MapFromProperty = "MaxLength"

	// The dependency on an attribute's ScalarIncrement.
	MapFromProperty_SCALAR_INCREMENT MapFromProperty = "ScalarIncrement"

)
