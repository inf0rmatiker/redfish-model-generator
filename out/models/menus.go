/* -----------------------------------------------------------------
* menus.go -
*
* DMTF Redfish Menus resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// An attribute's menu and its hierarchy.
type Menus struct {
	// The user-readable display string of this menu in the defined language.
	DisplayName string `json:"DisplayName,omitempty"`

	// The ascending order, as a number, in which this menu appears relative to other menus.
	DisplayOrder int `json:"DisplayOrder,omitempty"`

	// An indication of whether this menu is grayed out.  A grayed-only menu is not accessible in user interfaces.
	GrayOut bool `json:"GrayOut,omitempty"`

	// An indication of whether this menu is hidden in user interfaces.
	Hidden bool `json:"Hidden,omitempty"`

	// The unique name string of this menu.
	MenuName string `json:"MenuName,omitempty"`

	// The path to the menu names that describes this menu hierarchy relative to other menus.
	MenuPath string `json:"MenuPath,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this menu is read-only.  A read-only menu, its properties, and sub-menus are not accessible in user interfaces.
	ReadOnly bool `json:"ReadOnly,omitempty"`

}
