/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListDataSessionResponse struct for ListDataSessionResponse
type ListDataSessionResponse struct {
	DataSessions []WirelessV1SimDataSession `json:"data_sessions,omitempty"`
	Meta         ListCommandResponseMeta    `json:"meta,omitempty"`
}
