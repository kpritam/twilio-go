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

// ListCredentialResponse struct for ListCredentialResponse
type ListCredentialResponse struct {
	Credentials []ConversationsV1Credential  `json:"credentials,omitempty"`
	Meta        ListConversationResponseMeta `json:"meta,omitempty"`
}
