/* -----------------------------------------------------------------
* menus.go -
*
* DMTF Redfish Menus resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Menus - A menu and its hierarchy.
// Modeled after DMTF Redfish schema Menus
type Menus struct {
	// The user-readable display string of this menu in the defined 'Language'.
	DisplayName string `json:"DisplayName,omitempty"`

	// The numeric value describing the ascending order in which this menu is displayed relative to other menus.
	DisplayOrder int `json:"DisplayOrder,omitempty"`

	// The gray-out state of this menu. A grayed-only menu is not accessible in user interfaces.
	GrayOut bool `json:"GrayOut,omitempty"`

	// The unique name string of this menu.
	MenuName string `json:"MenuName,omitempty"`

	// A path that describes this menu hierarchy relative to other menus.
	MenuPath string `json:"MenuPath,omitempty"`

	// The read-only state of this menu.
	ReadOnly bool `json:"ReadOnly,omitempty"`

}
