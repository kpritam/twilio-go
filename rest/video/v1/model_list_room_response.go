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

// ListRoomResponse struct for ListRoomResponse
type ListRoomResponse struct {
	Meta  ListCompositionHookResponseMeta `json:"meta,omitempty"`
	Rooms []VideoV1Room                   `json:"rooms,omitempty"`
}
