/*
 * Twilio - Ip_messaging
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
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
	Services []IpMessagingV2Service     `json:"services,omitempty"`
}
