/*
 * Twilio - Voice
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

// VoiceV1ByocTrunk struct for VoiceV1ByocTrunk
type VoiceV1ByocTrunk struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk
	CnamLookupEnabled *bool `json:"cnam_lookup_enabled,omitempty"`
	// Origination Connection Policy (to your Carrier)
	ConnectionPolicySid *string `json:"connection_policy_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The SID of the SIP Domain that should be used in the `From` header of originating calls
	FromDomainSid *string `json:"from_domain_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The HTTP method we use to call status_callback_url
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// The URL that we call with status updates
	StatusCallbackUrl *string `json:"status_callback_url,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// The HTTP method used with voice_fallback_url
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
	// The URL we call when an error occurs while executing TwiML
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
	// The HTTP method to use with voice_url
	VoiceMethod *string `json:"voice_method,omitempty"`
	// The URL we call when receiving a call
	VoiceUrl *string `json:"voice_url,omitempty"`
}
