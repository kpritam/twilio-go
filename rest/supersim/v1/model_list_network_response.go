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

// ListNetworkResponse struct for ListNetworkResponse
type ListNetworkResponse struct {
	Meta     ListCommandResponseMeta `json:"meta,omitempty"`
	Networks []SupersimV1Network     `json:"networks,omitempty"`
}
