/* -----------------------------------------------------------------
* manager_account.go -
*
* DMTF Redfish ManagerAccount resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ManagerAccount schema defines the user accounts that are owned by a manager.  Changes to a manager account might affect the current Redfish service connection if this manager is responsible for the Redfish service.
type ManagerAccount struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// Indicates the date and time when this account expires.  If `null`, the account never expires.
	AccountExpiration string `json:"AccountExpiration,omitempty"`

	// The account types.
	AccountTypes array `json:"AccountTypes"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to a collection of certificates used for this account.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether an account is enabled.  An administrator can disable it without deleting the user information.  If `true`, the account is enabled and the user can log in.  If `false`, the account is disabled and, in the future, the user cannot log in.
	Enabled bool `json:"Enabled,omitempty"`

	// An indication of whether this account is a bootstrap account for the host interface.
	HostBootstrapAccount bool `json:"HostBootstrapAccount,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An indication of whether the account service automatically locked the account because the lockout threshold was exceeded.  To manually unlock the account before the lockout duration period, an administrator can change the property to `false` to clear the lockout condition.
	Locked bool `json:"Locked,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM account types.
	OEMAccountTypes []string `json:"OEMAccountTypes,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The password.  Use this property with a PATCH or PUT to write the password for the account.  This property is `null` in responses.
	Password string `json:"Password,omitempty"`

	// An indication of whether the service requires that the password for this account be changed before further access to the account is allowed.
	PasswordChangeRequired bool `json:"PasswordChangeRequired,omitempty"`

	// Indicates the date and time when this account password expires.  If `null`, the account password never expires.
	PasswordExpiration string `json:"PasswordExpiration,omitempty"`

	// The role for this account.
	RoleId string `json:"RoleId,omitempty"`

	// The SNMP settings for this account.
	SNMP SNMPUserInfo `json:"SNMP,omitempty"`

	// Indicates if the service needs to use the account types exactly as specified when the account is created or updated.
	StrictAccountTypes bool `json:"StrictAccountTypes,omitempty"`

	// The user name for the account.
	UserName string `json:"UserName,omitempty"`

}
