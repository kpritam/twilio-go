/*
 * Twilio - Serverless
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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to inject Account credentials into a function invocation context. The default value is `true`.
	IncludeCredentials *bool `json:"IncludeCredentials,omitempty"`
	// Whether the Service's properties and subresources can be edited via the UI. The default value is `false`.
	UiEditable *bool `json:"UiEditable,omitempty"`
	// A user-defined string that uniquely identifies the Service resource. It can be used as an alternative to the `sid` in the URL path to address the Service resource. This value must be 50 characters or less in length and be unique.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceParams) SetIncludeCredentials(IncludeCredentials bool) *CreateServiceParams {
	params.IncludeCredentials = &IncludeCredentials
	return params
}
func (params *CreateServiceParams) SetUiEditable(UiEditable bool) *CreateServiceParams {
	params.UiEditable = &UiEditable
	return params
}
func (params *CreateServiceParams) SetUniqueName(UniqueName string) *CreateServiceParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new Service resource.
func (c *ApiService) CreateService(params *CreateServiceParams) (*ServerlessV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IncludeCredentials != nil {
		data.Set("IncludeCredentials", fmt.Sprint(*params.IncludeCredentials))
	}
	if params != nil && params.UiEditable != nil {
		data.Set("UiEditable", fmt.Sprint(*params.UiEditable))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Service resource.
func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
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

// Retrieve a specific Service resource.
func (c *ApiService) FetchService(Sid string) (*ServerlessV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParams) SetLimit(Limit int) *ListServiceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken, pageNumber string) (*ListServiceResponse, error) {
	path := "/v1/Services"

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

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]ServerlessV1Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ServerlessV1Service

	for response != nil {
		records = append(records, response.Services...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListServiceResponse)
	}

	return records, err
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan ServerlessV1Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ServerlessV1Service, 1)

	go func() {
		for response != nil {
			for item := range response.Services {
				channel <- response.Services[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	// A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to inject Account credentials into a function invocation context.
	IncludeCredentials *bool `json:"IncludeCredentials,omitempty"`
	// Whether the Service resource's properties and subresources can be edited via the UI. The default value is `false`.
	UiEditable *bool `json:"UiEditable,omitempty"`
}

func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetIncludeCredentials(IncludeCredentials bool) *UpdateServiceParams {
	params.IncludeCredentials = &IncludeCredentials
	return params
}
func (params *UpdateServiceParams) SetUiEditable(UiEditable bool) *UpdateServiceParams {
	params.UiEditable = &UiEditable
	return params
}

// Update a specific Service resource.
func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*ServerlessV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IncludeCredentials != nil {
		data.Set("IncludeCredentials", fmt.Sprint(*params.IncludeCredentials))
	}
	if params != nil && params.UiEditable != nil {
		data.Set("UiEditable", fmt.Sprint(*params.UiEditable))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
