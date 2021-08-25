/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VerifyV2Form struct for VerifyV2Form
type VerifyV2Form struct {
	// Additional information for the available forms for this type.
	FormMeta *map[string]interface{} `json:"form_meta,omitempty"`
	// The Type of this Form
	FormType *string `json:"form_type,omitempty"`
	// Object that contains the available forms for this type.
	Forms *map[string]interface{} `json:"forms,omitempty"`
	// The URL to access the forms for this type.
	Url *string `json:"url,omitempty"`
}
