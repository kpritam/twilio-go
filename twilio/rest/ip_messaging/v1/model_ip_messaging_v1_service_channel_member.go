/*
 * Twilio - Ip_messaging
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

// IpMessagingV1ServiceChannelMember struct for IpMessagingV1ServiceChannelMember
type IpMessagingV1ServiceChannelMember struct {
	AccountSid               *string    `json:"account_sid,omitempty"`
	ChannelSid               *string    `json:"channel_sid,omitempty"`
	DateCreated              *time.Time `json:"date_created,omitempty"`
	DateUpdated              *time.Time `json:"date_updated,omitempty"`
	Identity                 *string    `json:"identity,omitempty"`
	LastConsumedMessageIndex *int32     `json:"last_consumed_message_index,omitempty"`
	LastConsumptionTimestamp *time.Time `json:"last_consumption_timestamp,omitempty"`
	RoleSid                  *string    `json:"role_sid,omitempty"`
	ServiceSid               *string    `json:"service_sid,omitempty"`
	Sid                      *string    `json:"sid,omitempty"`
	Url                      *string    `json:"url,omitempty"`
}
