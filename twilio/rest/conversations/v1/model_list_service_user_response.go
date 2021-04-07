/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceUserResponse struct for ListServiceUserResponse
type ListServiceUserResponse struct {
	Meta  ListConversationResponseMeta        `json:"meta,omitempty"`
	Users []ConversationsV1ServiceServiceUser `json:"users,omitempty"`
}
