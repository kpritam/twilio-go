/*
 * Twilio - Supersim
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

// SupersimV1Command struct for SupersimV1Command
type SupersimV1Command struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The message body of the command sent to or from the SIM
	Command *string `json:"command,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The direction of the Command
	Direction *string `json:"direction,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the SIM that this Command was sent to or from
	SimSid *string `json:"sim_sid,omitempty"`
	// The status of the Command
	Status *string `json:"status,omitempty"`
	// The absolute URL of the Command resource
	Url *string `json:"url,omitempty"`
}
