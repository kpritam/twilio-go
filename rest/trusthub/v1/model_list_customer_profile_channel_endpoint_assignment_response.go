/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCustomerProfileChannelEndpointAssignmentResponse struct for ListCustomerProfileChannelEndpointAssignmentResponse
type ListCustomerProfileChannelEndpointAssignmentResponse struct {
	Meta    ListCustomerProfileResponseMeta                      `json:"meta,omitempty"`
	Results []TrusthubV1CustomerProfileChannelEndpointAssignment `json:"results,omitempty"`
}
