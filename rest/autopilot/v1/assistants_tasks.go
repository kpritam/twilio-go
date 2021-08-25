/*
 * Twilio - Autopilot
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

// Optional parameters for the method 'CreateTask'
type CreateTaskParams struct {
	// The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique.
	Actions *map[string]interface{} `json:"Actions,omitempty"`
	// The URL from which the Assistant can fetch actions.
	ActionsUrl *string `json:"ActionsUrl,omitempty"`
	// A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateTaskParams) SetActions(Actions map[string]interface{}) *CreateTaskParams {
	params.Actions = &Actions
	return params
}
func (params *CreateTaskParams) SetActionsUrl(ActionsUrl string) *CreateTaskParams {
	params.ActionsUrl = &ActionsUrl
	return params
}
func (params *CreateTaskParams) SetFriendlyName(FriendlyName string) *CreateTaskParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTaskParams) SetUniqueName(UniqueName string) *CreateTaskParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateTask(AssistantSid string, params *CreateTaskParams) (*AutopilotV1Task, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Actions != nil {
		v, err := json.Marshal(params.Actions)

		if err != nil {
			return nil, err
		}

		data.Set("Actions", string(v))
	}
	if params != nil && params.ActionsUrl != nil {
		data.Set("ActionsUrl", *params.ActionsUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteTask(AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
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

func (c *ApiService) FetchTask(AssistantSid string, Sid string) (*AutopilotV1Task, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTask'
type ListTaskParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTaskParams) SetPageSize(PageSize int) *ListTaskParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTaskParams) SetLimit(Limit int) *ListTaskParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Task records from the API. Request is executed immediately.
func (c *ApiService) PageTask(AssistantSid string, params *ListTaskParams, pageToken string, pageNumber string) (*ListTaskResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

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

	ps := &ListTaskResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Task records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTask(AssistantSid string, params *ListTaskParams) ([]AutopilotV1Task, error) {
	if params == nil {
		params = &ListTaskParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTask(AssistantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AutopilotV1Task

	for response != nil {
		records = append(records, response.Tasks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListTaskResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListTaskResponse)
	}

	return records, err
}

// Streams Task records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTask(AssistantSid string, params *ListTaskParams) (chan AutopilotV1Task, error) {
	if params == nil {
		params = &ListTaskParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTask(AssistantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AutopilotV1Task, 1)

	go func() {
		for response != nil {
			for item := range response.Tasks {
				channel <- response.Tasks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListTaskResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTaskResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTaskResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateTask'
type UpdateTaskParams struct {
	// The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.
	Actions *map[string]interface{} `json:"Actions,omitempty"`
	// The URL from which the Assistant can fetch actions.
	ActionsUrl *string `json:"ActionsUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateTaskParams) SetActions(Actions map[string]interface{}) *UpdateTaskParams {
	params.Actions = &Actions
	return params
}
func (params *UpdateTaskParams) SetActionsUrl(ActionsUrl string) *UpdateTaskParams {
	params.ActionsUrl = &ActionsUrl
	return params
}
func (params *UpdateTaskParams) SetFriendlyName(FriendlyName string) *UpdateTaskParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTaskParams) SetUniqueName(UniqueName string) *UpdateTaskParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateTask(AssistantSid string, Sid string, params *UpdateTaskParams) (*AutopilotV1Task, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Actions != nil {
		v, err := json.Marshal(params.Actions)

		if err != nil {
			return nil, err
		}

		data.Set("Actions", string(v))
	}
	if params != nil && params.ActionsUrl != nil {
		data.Set("ActionsUrl", *params.ActionsUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
