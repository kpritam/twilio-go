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

// ListMemberResponse struct for ListMemberResponse
type ListMemberResponse struct {
	Members []IpMessagingV2Member      `json:"members,omitempty"`
	Meta    ListCredentialResponseMeta `json:"meta,omitempty"`
}
