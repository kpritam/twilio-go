/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceResponse struct for ListServiceResponse
type ListServiceResponse struct {
	Meta     ListServiceResponseMeta `json:"meta,omitempty"`
	Services []SyncV1Service         `json:"services,omitempty"`
}
