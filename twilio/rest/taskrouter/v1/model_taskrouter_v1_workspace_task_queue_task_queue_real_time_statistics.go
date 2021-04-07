/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics struct for TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics
type TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The number of current Workers by Activity
	ActivityStatistics *[]map[string]interface{} `json:"activity_statistics,omitempty"`
	// The relative age in the TaskQueue for the longest waiting Task.
	LongestRelativeTaskAgeInQueue *int32 `json:"longest_relative_task_age_in_queue,omitempty"`
	// The SID of the Task waiting in the TaskQueue the longest.
	LongestRelativeTaskSidInQueue *string `json:"longest_relative_task_sid_in_queue,omitempty"`
	// The age of the longest waiting Task
	LongestTaskWaitingAge *int32 `json:"longest_task_waiting_age,omitempty"`
	// The SID of the longest waiting Task
	LongestTaskWaitingSid *string `json:"longest_task_waiting_sid,omitempty"`
	// The SID of the TaskQueue from which these statistics were calculated
	TaskQueueSid *string `json:"task_queue_sid,omitempty"`
	// The number of Tasks by priority
	TasksByPriority *map[string]interface{} `json:"tasks_by_priority,omitempty"`
	// The number of Tasks by their current status
	TasksByStatus *map[string]interface{} `json:"tasks_by_status,omitempty"`
	// The total number of Workers available for Tasks in the TaskQueue
	TotalAvailableWorkers *int32 `json:"total_available_workers,omitempty"`
	// The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state
	TotalEligibleWorkers *int32 `json:"total_eligible_workers,omitempty"`
	// The total number of Tasks
	TotalTasks *int32 `json:"total_tasks,omitempty"`
	// The absolute URL of the TaskQueue statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the TaskQueue
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
