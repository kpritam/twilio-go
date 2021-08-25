/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV1EngagementContext struct for StudioV1EngagementContext
type StudioV1EngagementContext struct {
	// Account SID
	AccountSid *string `json:"account_sid,omitempty"`
	// Flow state
	Context *map[string]interface{} `json:"context,omitempty"`
	// Engagement SID
	EngagementSid *string `json:"engagement_sid,omitempty"`
	// Flow SID
	FlowSid *string `json:"flow_sid,omitempty"`
	// The URL of the resource
	Url *string `json:"url,omitempty"`
}
