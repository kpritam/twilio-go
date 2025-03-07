/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010CallFeedbackSummary struct for ApiV2010CallFeedbackSummary
type ApiV2010CallFeedbackSummary struct {
	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`
	// The total number of calls
	CallCount *int `json:"call_count,omitempty"`
	// The total number of calls with a feedback entry
	CallFeedbackCount *int `json:"call_feedback_count,omitempty"`
	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The latest feedback entry date in the summary
	EndDate *string `json:"end_date,omitempty"`
	// Whether the feedback summary includes subaccounts
	IncludeSubaccounts *bool `json:"include_subaccounts,omitempty"`
	// Issues experienced during the call
	Issues *[]map[string]interface{} `json:"issues,omitempty"`
	// The average QualityScore of the feedback entries
	QualityScoreAverage *float32 `json:"quality_score_average,omitempty"`
	// The median QualityScore of the feedback entries
	QualityScoreMedian *float32 `json:"quality_score_median,omitempty"`
	// The standard deviation of the quality scores
	QualityScoreStandardDeviation *float32 `json:"quality_score_standard_deviation,omitempty"`
	// A string that uniquely identifies this feedback entry
	Sid *string `json:"sid,omitempty"`
	// The earliest feedback entry date in the summary
	StartDate *string `json:"start_date,omitempty"`
	// The status of the feedback summary
	Status *string `json:"status,omitempty"`
}
