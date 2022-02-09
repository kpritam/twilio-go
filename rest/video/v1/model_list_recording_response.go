/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRecordingResponse struct for ListRecordingResponse
type ListRecordingResponse struct {
	Meta       ListCompositionHookResponseMeta `json:"meta,omitempty"`
	Recordings []VideoV1Recording              `json:"recordings,omitempty"`
}
