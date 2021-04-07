/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VerifyV2ServiceMessagingConfiguration struct for VerifyV2ServiceMessagingConfiguration
type VerifyV2ServiceMessagingConfiguration struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO-3166-1 country code of the country or `all`.
	Country *string `json:"country,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Messaging Service used for this configuration.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
