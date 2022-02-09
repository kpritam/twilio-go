/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1ExportConfiguration struct for BulkexportsV1ExportConfiguration
type BulkexportsV1ExportConfiguration struct {
	// Whether files are automatically generated
	Enabled *bool `json:"enabled,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Whether to GET or POST to the webhook url
	WebhookMethod *string `json:"webhook_method,omitempty"`
	// URL targeted at export
	WebhookUrl *string `json:"webhook_url,omitempty"`
}
