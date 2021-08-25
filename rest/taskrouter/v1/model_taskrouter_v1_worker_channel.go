/*
 * Twilio - Taskrouter
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

// TaskrouterV1WorkerChannel struct for TaskrouterV1WorkerChannel
type TaskrouterV1WorkerChannel struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The total number of Tasks assigned to Worker for the TaskChannel type
	AssignedTasks *int `json:"assigned_tasks,omitempty"`
	// Whether the Worker should receive Tasks of the TaskChannel type
	Available *bool `json:"available,omitempty"`
	// The current available capacity between 0 to 100 for the TaskChannel
	AvailableCapacityPercentage *int `json:"available_capacity_percentage,omitempty"`
	// The current configured capacity for the WorkerChannel
	ConfiguredCapacity *int `json:"configured_capacity,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the TaskChannel
	TaskChannelSid *string `json:"task_channel_sid,omitempty"`
	// The unique name of the TaskChannel, such as 'voice' or 'sms'
	TaskChannelUniqueName *string `json:"task_channel_unique_name,omitempty"`
	// The absolute URL of the WorkerChannel resource
	Url *string `json:"url,omitempty"`
	// The SID of the Worker that contains the WorkerChannel
	WorkerSid *string `json:"worker_sid,omitempty"`
	// The SID of the Workspace that contains the WorkerChannel
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
