/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListVoiceCountryResponse struct for ListVoiceCountryResponse
type ListVoiceCountryResponse struct {
	Countries []PricingV1VoiceCountry          `json:"countries,omitempty"`
	Meta      ListMessagingCountryResponseMeta `json:"meta,omitempty"`
}
