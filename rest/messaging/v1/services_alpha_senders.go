/*
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateAlphaSender'
type CreateAlphaSenderParams struct {
	// The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, and hyphen `-`. This value cannot contain only numbers.
	AlphaSender *string `json:"AlphaSender,omitempty"`
}

func (params *CreateAlphaSenderParams) SetAlphaSender(AlphaSender string) *CreateAlphaSenderParams {
	params.AlphaSender = &AlphaSender
	return params
}

func (c *ApiService) CreateAlphaSender(ServiceSid string, params *CreateAlphaSenderParams) (*MessagingV1AlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AlphaSender != nil {
		data.Set("AlphaSender", *params.AlphaSender)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1AlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteAlphaSender(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/AlphaSenders/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchAlphaSender(ServiceSid string, Sid string) (*MessagingV1AlphaSender, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1AlphaSender{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAlphaSender'
type ListAlphaSenderParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAlphaSenderParams) SetPageSize(PageSize int) *ListAlphaSenderParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAlphaSenderParams) SetLimit(Limit int) *ListAlphaSenderParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AlphaSender records from the API. Request is executed immediately.
func (c *ApiService) PageAlphaSender(ServiceSid string, params *ListAlphaSenderParams, pageToken string, pageNumber string) (*ListAlphaSenderResponse, error) {
	path := "/v1/Services/{ServiceSid}/AlphaSenders"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AlphaSender records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAlphaSender(ServiceSid string, params *ListAlphaSenderParams) ([]MessagingV1AlphaSender, error) {
	if params == nil {
		params = &ListAlphaSenderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAlphaSender(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []MessagingV1AlphaSender

	for response != nil {
		records = append(records, response.AlphaSenders...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAlphaSenderResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAlphaSenderResponse)
	}

	return records, err
}

// Streams AlphaSender records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAlphaSender(ServiceSid string, params *ListAlphaSenderParams) (chan MessagingV1AlphaSender, error) {
	if params == nil {
		params = &ListAlphaSenderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAlphaSender(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1AlphaSender, 1)

	go func() {
		for response != nil {
			for item := range response.AlphaSenders {
				channel <- response.AlphaSenders[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAlphaSenderResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAlphaSenderResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAlphaSenderResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAlphaSenderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
