/* -----------------------------------------------------------------
* authentication.go -
*
* DMTF Redfish Authentication resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information required to authenticate to the external service.
type Authentication struct {
	// The type of authentication used to connect to the external account provider.
	AuthenticationType AuthenticationTypes `json:"AuthenticationType,omitempty"`

	// Specifies the encryption key.
	EncryptionKey string `json:"EncryptionKey,omitempty"`

	// Indicates if the EncryptionKey property is set.
	EncryptionKeySet bool `json:"EncryptionKeySet,omitempty"`

	// The Base64-encoded version of the Kerberos keytab for this service.  A PATCH or PUT operation writes the keytab.  This property is `null` in responses.
	KerberosKeytab string `json:"KerberosKeytab,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The password for this service.  A PATCH or PUT request writes the password.  This property is `null` in responses.
	Password string `json:"Password,omitempty"`

	// The token for this service.  A PATCH or PUT operation writes the token.  This property is `null` in responses.
	Token string `json:"Token,omitempty"`

	// The user name for the service.
	Username string `json:"Username,omitempty"`

}
