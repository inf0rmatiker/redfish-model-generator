/* -----------------------------------------------------------------
* smtp.go -
*
* DMTF Redfish SMTP resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Settings for SMTP event delivery.
type SMTP struct {
	// The authentication method for the SMTP server.
	Authentication SMTPAuthenticationMethods `json:"Authentication,omitempty"`

	// The connection type to the outgoing SMTP server.
	ConnectionProtocol SMTPConnectionProtocol `json:"ConnectionProtocol,omitempty"`

	// The 'from' email address of the outgoing email.
	FromAddress string `json:"FromAddress,omitempty"`

	// The password for authentication with the SMTP server.  The value is `null` in responses.
	Password string `json:"Password,omitempty"`

	// The destination SMTP port.
	Port int `json:"Port,omitempty"`

	// The address of the SMTP server.
	ServerAddress string `json:"ServerAddress,omitempty"`

	// An indication if SMTP for event delivery is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The username for authentication with the SMTP server.
	Username string `json:"Username,omitempty"`

}
