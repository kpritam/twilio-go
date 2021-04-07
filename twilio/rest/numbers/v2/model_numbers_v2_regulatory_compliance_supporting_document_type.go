/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// NumbersV2RegulatoryComplianceSupportingDocumentType struct for NumbersV2RegulatoryComplianceSupportingDocumentType
type NumbersV2RegulatoryComplianceSupportingDocumentType struct {
	// The required information for creating a Supporting Document
	Fields *[]map[string]interface{} `json:"fields,omitempty"`
	// A human-readable description of the Supporting Document Type resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The machine-readable description of the Supporting Document Type resource
	MachineName *string `json:"machine_name,omitempty"`
	// The unique string that identifies the Supporting Document Type resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Supporting Document Type resource
	Url *string `json:"url,omitempty"`
}
