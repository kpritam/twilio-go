/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
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

func (c *ApiService) FetchMessagingCountry(IsoCountry string) (*PricingV1MessagingCountryInstance, error) {
	path := "/v1/Messaging/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1MessagingCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessagingCountry'
type ListMessagingCountryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMessagingCountryParams) SetPageSize(PageSize int) *ListMessagingCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMessagingCountryParams) SetLimit(Limit int) *ListMessagingCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of MessagingCountry records from the API. Request is executed immediately.
func (c *ApiService) PageMessagingCountry(params *ListMessagingCountryParams, pageToken string, pageNumber string) (*ListMessagingCountryResponse, error) {
	path := "/v1/Messaging/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessagingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists MessagingCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessagingCountry(params *ListMessagingCountryParams) ([]PricingV1MessagingCountry, error) {
	if params == nil {
		params = &ListMessagingCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessagingCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []PricingV1MessagingCountry

	for response != nil {
		records = append(records, response.Countries...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessagingCountryResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListMessagingCountryResponse)
	}

	return records, err
}

// Streams MessagingCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessagingCountry(params *ListMessagingCountryParams) (chan PricingV1MessagingCountry, error) {
	if params == nil {
		params = &ListMessagingCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessagingCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan PricingV1MessagingCountry, 1)

	go func() {
		for response != nil {
			for item := range response.Countries {
				channel <- response.Countries[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessagingCountryResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMessagingCountryResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMessagingCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessagingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
