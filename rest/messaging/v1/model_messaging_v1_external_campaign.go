/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// MessagingV1ExternalCampaign struct for MessagingV1ExternalCampaign
type MessagingV1ExternalCampaign struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// ID of the preregistered campaign.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The SID of the Messaging Service the resource is associated with
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The unique string that identifies a US A2P Compliance resource
	Sid *string `json:"sid,omitempty"`
}
