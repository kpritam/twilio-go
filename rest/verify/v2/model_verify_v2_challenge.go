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

// VerifyV2Challenge struct for VerifyV2Challenge
type VerifyV2Challenge struct {
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this Challenge was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Challenge was responded
	DateResponded *time.Time `json:"date_responded,omitempty"`
	// The date this Challenge was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Details about the Challenge.
	Details *map[string]interface{} `json:"details,omitempty"`
	// Entity Sid.
	EntitySid *string `json:"entity_sid,omitempty"`
	// The date-time when this Challenge expires
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// Factor Sid.
	FactorSid *string `json:"factor_sid,omitempty"`
	// The Factor Type of this Challenge
	FactorType *string `json:"factor_type,omitempty"`
	// Hidden details about the Challenge
	HiddenDetails *map[string]interface{} `json:"hidden_details,omitempty"`
	// Unique external identifier of the Entity
	Identity *string `json:"identity,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The Reason of this Challenge `status`
	RespondedReason *string `json:"responded_reason,omitempty"`
	// Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// A string that uniquely identifies this Challenge.
	Sid *string `json:"sid,omitempty"`
	// The Status of this Challenge
	Status *string `json:"status,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
