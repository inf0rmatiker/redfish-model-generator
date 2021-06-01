/* -----------------------------------------------------------------
* settings.go -
*
* DMTF Redfish Settings resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Settings - The Resource settings.
// Modeled after DMTF Redfish schema Settings
type Settings struct {
	// The entity tag (ETag) of the Resource to which the settings were applied, after the application.
	ETag string `json:"ETag,omitempty"`

	// The location of the maintenance window settings.
	MaintenanceWindowResource map[string]interface{} `json:"MaintenanceWindowResource,omitempty"`

	// An array of messages associated with the settings.
	Messages []Message `json:"Messages,omitempty"`

	// The link to the Resource that the client may PUT or PATCH to modify the Resource.
	SettingsObject map[string]interface{} `json:"SettingsObject,omitempty"`

	// The time when the settings can be applied.
	SupportedApplyTimes []ApplyTime `json:"SupportedApplyTimes,omitempty"`

	// The time when the settings were applied.
	Time string `json:"Time,omitempty"`

}
