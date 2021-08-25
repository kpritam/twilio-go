/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceParticipantConversationResponse struct for ListServiceParticipantConversationResponse
type ListServiceParticipantConversationResponse struct {
	Conversations []ConversationsV1ServiceParticipantConversation `json:"conversations,omitempty"`
	Meta          ListConversationResponseMeta                    `json:"meta,omitempty"`
}
