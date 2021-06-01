/* -----------------------------------------------------------------
* manager_account.go -
*
* DMTF Redfish ManagerAccount resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ManagerAccount - The user accounts, owned by a manager, are defined in this Resource.  Changes to a manager account may affect the current Redfish Service connection if this manager is responsible for the Redfish Service.
// Modeled after DMTF Redfish schema ManagerAccount
type ManagerAccount struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to a collection of certificates used for this account.
	Certificates map[string]interface{} `json:"Certificates,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether an account is enabled.  An administrator can disable it without deleting the user information.  If `true`, the account is enabled and the user can log in.  If `false`, the account is disabled and, in the future, the user cannot log in.  This property overrides the Locked property.
	Enabled bool `json:"Enabled,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// An indication of whether the Account Service automatically locked the account because the lockout threshold was exceeded.  To manually unlock the account before the lockout duration period, an administrator can change the property to `false` to clear the lockout condition.
	Locked bool `json:"Locked,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The password.  Use this property with a PATCH or PUT to write the password for the account.  This property is `null` in responses.
	Password string `json:"Password,omitempty"`

	// The Role for this account.
	RoleId string `json:"RoleId,omitempty"`

	// The user name for the account.
	UserName string `json:"UserName,omitempty"`

}
