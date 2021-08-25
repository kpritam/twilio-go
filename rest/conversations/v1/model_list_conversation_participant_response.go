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

// ListConversationParticipantResponse struct for ListConversationParticipantResponse
type ListConversationParticipantResponse struct {
	Meta         ListConversationResponseMeta             `json:"meta,omitempty"`
	Participants []ConversationsV1ConversationParticipant `json:"participants,omitempty"`
}
