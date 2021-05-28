/* -----------------------------------------------------------------
 * enum_map_to_property.go -
 *
 * DMTF Redfish MapToProperty enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MapToProperty string

const (
	// The dependency that affects an attribute's CurrentValue.
	MapToProperty_CURRENT_VALUE MapToProperty = "CurrentValue"

	// The dependency that affects an attribute's DefaultValue.
	MapToProperty_DEFAULT_VALUE MapToProperty = "DefaultValue"

	// The dependency that affects an attribute's ReadOnly state.
	MapToProperty_READ_ONLY MapToProperty = "ReadOnly"

	// The dependency that affects an attribute's WriteOnly state.
	MapToProperty_WRITE_ONLY MapToProperty = "WriteOnly"

	// The dependency that affects an attribute's GrayOut state.
	MapToProperty_GRAY_OUT MapToProperty = "GrayOut"

	// The dependency that affects an attribute's Hidden state.
	MapToProperty_HIDDEN MapToProperty = "Hidden"

	// The dependency that affects an attribute's Immutable state.
	MapToProperty_IMMUTABLE MapToProperty = "Immutable"

	// The dependency that affects an attribute's HelpText.
	MapToProperty_HELP_TEXT MapToProperty = "HelpText"

	// The dependency that affects an attribute's WarningText.
	MapToProperty_WARNING_TEXT MapToProperty = "WarningText"

	// The dependency that affects an attribute's DisplayName.
	MapToProperty_DISPLAY_NAME MapToProperty = "DisplayName"

	// The dependency that affects an attribute's DisplayName.
	MapToProperty_DISPLAY_ORDER MapToProperty = "DisplayOrder"

	// The dependency that affects an attribute's LowerBound.
	MapToProperty_LOWER_BOUND MapToProperty = "LowerBound"

	// The dependency that affects an attribute's UpperBound.
	MapToProperty_UPPER_BOUND MapToProperty = "UpperBound"

	// The dependency that affects an attribute's MinLength.
	MapToProperty_MIN_LENGTH MapToProperty = "MinLength"

	// The dependency that affects an attribute's MaxLength.
	MapToProperty_MAX_LENGTH MapToProperty = "MaxLength"

	// The dependency that affects an attribute's ScalarIncrement.
	MapToProperty_SCALAR_INCREMENT MapToProperty = "ScalarIncrement"

	// The dependency that affects an attribute's ValueExpression.
	MapToProperty_VALUE_EXPRESSION MapToProperty = "ValueExpression"

)
