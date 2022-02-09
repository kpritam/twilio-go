/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListChannelResponse struct for ListChannelResponse
type ListChannelResponse struct {
	Channels []IpMessagingV2Channel     `json:"channels,omitempty"`
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
}
