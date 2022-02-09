/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListBillingPeriod'
type ListBillingPeriodParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBillingPeriodParams) SetPageSize(PageSize int) *ListBillingPeriodParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBillingPeriodParams) SetLimit(Limit int) *ListBillingPeriodParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of BillingPeriod records from the API. Request is executed immediately.
func (c *ApiService) PageBillingPeriod(SimSid string, params *ListBillingPeriodParams, pageToken, pageNumber string) (*ListBillingPeriodResponse, error) {
	path := "/v1/Sims/{SimSid}/BillingPeriods"

	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BillingPeriod records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBillingPeriod(SimSid string, params *ListBillingPeriodParams) ([]SupersimV1BillingPeriod, error) {
	if params == nil {
		params = &ListBillingPeriodParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBillingPeriod(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SupersimV1BillingPeriod

	for response != nil {
		records = append(records, response.BillingPeriods...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListBillingPeriodResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListBillingPeriodResponse)
	}

	return records, err
}

// Streams BillingPeriod records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBillingPeriod(SimSid string, params *ListBillingPeriodParams) (chan SupersimV1BillingPeriod, error) {
	if params == nil {
		params = &ListBillingPeriodParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBillingPeriod(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SupersimV1BillingPeriod, 1)

	go func() {
		for response != nil {
			for item := range response.BillingPeriods {
				channel <- response.BillingPeriods[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListBillingPeriodResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBillingPeriodResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBillingPeriodResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
