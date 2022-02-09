/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSupportingDocumentTypeResponse struct for ListSupportingDocumentTypeResponse
type ListSupportingDocumentTypeResponse struct {
	Meta                    ListCustomerProfileResponseMeta    `json:"meta,omitempty"`
	SupportingDocumentTypes []TrusthubV1SupportingDocumentType `json:"supporting_document_types,omitempty"`
}
