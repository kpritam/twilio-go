/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2VoiceVoiceCountryInstance struct for PricingV2VoiceVoiceCountryInstance
type PricingV2VoiceVoiceCountryInstance struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The list of InboundCallPrice records
	InboundCallPrices *[]PricingV2VoiceVoiceCountryInstanceInboundCallPrices `json:"inbound_call_prices,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The list of OutboundPrefixPriceWithOrigin records
	OutboundPrefixPrices *[]PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices `json:"outbound_prefix_prices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
