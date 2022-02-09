/*
 * Twilio - Taskrouter
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

// TaskrouterV1Workflow struct for TaskrouterV1Workflow
type TaskrouterV1Workflow struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The URL that we call when a task managed by the Workflow is assigned to a Worker
	AssignmentCallbackUrl *string `json:"assignment_callback_url,omitempty"`
	// A JSON string that contains the Workflow's configuration
	Configuration *string `json:"configuration,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The MIME type of the document
	DocumentContentType *string `json:"document_content_type,omitempty"`
	// The URL that we call when a call to the `assignment_callback_url` fails
	FallbackAssignmentCallbackUrl *string `json:"fallback_assignment_callback_url,omitempty"`
	// The string that you assigned to describe the Workflow resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker
	TaskReservationTimeout *int `json:"task_reservation_timeout,omitempty"`
	// The absolute URL of the Workflow resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the Workflow
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
