/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConversationResponse struct for ListConversationResponse
type ListConversationResponse struct {
	Conversations []ConversationsV1Conversation        `json:"conversations,omitempty"`
	Meta          ListConfigurationAddressResponseMeta `json:"meta,omitempty"`
}
