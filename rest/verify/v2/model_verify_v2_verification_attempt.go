/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VerifyV2VerificationAttempt struct for VerifyV2VerificationAttempt
type VerifyV2VerificationAttempt struct {
	// The SID of the Account that created the verification.
	AccountSid *string `json:"account_sid,omitempty"`
	// Communication channel used for the attempt.
	Channel *string `json:"channel,omitempty"`
	// An object containing the channel specific information for an attempt.
	ChannelData *map[string]interface{} `json:"channel_data,omitempty"`
	// Status of the conversion for the verification.
	ConversionStatus *string `json:"conversion_status,omitempty"`
	// The date this Attempt was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Attempt was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An object containing the charge for this verification attempt.
	Price *map[string]interface{} `json:"price,omitempty"`
	// The SID of the verify service that generated this attempt.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID that uniquely identifies the verification attempt.
	Sid *string `json:"sid,omitempty"`
	Url *string `json:"url,omitempty"`
	// The SID of the verification that generated this attempt.
	VerificationSid *string `json:"verification_sid,omitempty"`
}
