/* -----------------------------------------------------------------
* settings.go -
*
* DMTF Redfish Settings resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The resource settings.
type Settings struct {
	// The entity tag (ETag) of the resource to which the settings were applied, after the application.
	ETag string `json:"ETag,omitempty"`

	// The location of the maintenance window settings.
	MaintenanceWindowResource idRef `json:"MaintenanceWindowResource,omitempty"`

	// An array of messages associated with the settings.
	Messages array `json:"Messages,omitempty"`

	// The link to the resource that the client can PUT or PATCH to modify the resource.
	SettingsObject idRef `json:"SettingsObject,omitempty"`

	// The time when the settings can be applied.
	SupportedApplyTimes array `json:"SupportedApplyTimes,omitempty"`

	// The time when the settings were applied.
	Time string `json:"Time,omitempty"`

}
