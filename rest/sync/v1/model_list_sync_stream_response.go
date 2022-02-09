/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncStreamResponse struct for ListSyncStreamResponse
type ListSyncStreamResponse struct {
	Meta    ListServiceResponseMeta `json:"meta,omitempty"`
	Streams []SyncV1SyncStream      `json:"streams,omitempty"`
}
