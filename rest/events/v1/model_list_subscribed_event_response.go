/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSubscribedEventResponse struct for ListSubscribedEventResponse
type ListSubscribedEventResponse struct {
	Meta  ListSchemaVersionResponseMeta `json:"meta,omitempty"`
	Types []EventsV1SubscribedEvent     `json:"types,omitempty"`
}
