/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListVerificationTemplate'
type ListVerificationTemplateParams struct {
	// String filter used to query templates with a given friendly name
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListVerificationTemplateParams) SetFriendlyName(FriendlyName string) *ListVerificationTemplateParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListVerificationTemplateParams) SetPageSize(PageSize int) *ListVerificationTemplateParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListVerificationTemplateParams) SetLimit(Limit int) *ListVerificationTemplateParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of VerificationTemplate records from the API. Request is executed immediately.
func (c *ApiService) PageVerificationTemplate(params *ListVerificationTemplateParams, pageToken, pageNumber string) (*ListVerificationTemplateResponse, error) {
	path := "/v2/Templates"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVerificationTemplateResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists VerificationTemplate records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListVerificationTemplate(params *ListVerificationTemplateParams) ([]VerifyV2VerificationTemplate, error) {
	response, err := c.StreamVerificationTemplate(params)
	if err != nil {
		return nil, err
	}

	records := make([]VerifyV2VerificationTemplate, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams VerificationTemplate records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamVerificationTemplate(params *ListVerificationTemplateParams) (chan VerifyV2VerificationTemplate, error) {
	if params == nil {
		params = &ListVerificationTemplateParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVerificationTemplate(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2VerificationTemplate, 1)

	go func() {
		for response != nil {
			responseRecords := response.Templates
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListVerificationTemplateResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListVerificationTemplateResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListVerificationTemplateResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVerificationTemplateResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
