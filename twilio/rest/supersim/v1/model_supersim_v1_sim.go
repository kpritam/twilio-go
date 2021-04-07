/*
 * Twilio - Supersim
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

// SupersimV1Sim struct for SupersimV1Sim
type SupersimV1Sim struct {
	// The SID of the Account that the Super SIM belongs to
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique ID of the Fleet configured for this SIM
	FleetSid *string `json:"fleet_sid,omitempty"`
	// The ICCID associated with the SIM
	Iccid *string `json:"iccid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the Super SIM
	Status *string `json:"status,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Sim Resource
	Url *string `json:"url,omitempty"`
}
