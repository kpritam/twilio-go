/*
 * Twilio - Studio
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

// Optional parameters for the method 'CreateEngagement'
type CreateEngagementParams struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}`
	From *string `json:"From,omitempty"`
	// A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
	Parameters *map[string]interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`.
	To *string `json:"To,omitempty"`
}

func (params *CreateEngagementParams) SetFrom(From string) *CreateEngagementParams {
	params.From = &From
	return params
}
func (params *CreateEngagementParams) SetParameters(Parameters map[string]interface{}) *CreateEngagementParams {
	params.Parameters = &Parameters
	return params
}
func (params *CreateEngagementParams) SetTo(To string) *CreateEngagementParams {
	params.To = &To
	return params
}

// Triggers a new Engagement for the Flow
func (c *ApiService) CreateEngagement(FlowSid string, params *CreateEngagementParams) (*StudioV1Engagement, error) {
	path := "/v1/Flows/{FlowSid}/Engagements"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Parameters != nil {
		v, err := json.Marshal(params.Parameters)

		if err != nil {
			return nil, err
		}

		data.Set("Parameters", string(v))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Engagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete this Engagement and all Steps relating to it.
func (c *ApiService) DeleteEngagement(FlowSid string, Sid string) error {
	path := "/v1/Flows/{FlowSid}/Engagements/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
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

// Retrieve an Engagement
func (c *ApiService) FetchEngagement(FlowSid string, Sid string) (*StudioV1Engagement, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Engagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEngagement'
type ListEngagementParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEngagementParams) SetPageSize(PageSize int) *ListEngagementParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEngagementParams) SetLimit(Limit int) *ListEngagementParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Engagement records from the API. Request is executed immediately.
func (c *ApiService) PageEngagement(FlowSid string, params *ListEngagementParams, pageToken, pageNumber string) (*ListEngagementResponse, error) {
	path := "/v1/Flows/{FlowSid}/Engagements"

	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

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

	ps := &ListEngagementResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Engagement records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEngagement(FlowSid string, params *ListEngagementParams) ([]StudioV1Engagement, error) {
	if params == nil {
		params = &ListEngagementParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEngagement(FlowSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []StudioV1Engagement

	for response != nil {
		records = append(records, response.Engagements...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEngagementResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListEngagementResponse)
	}

	return records, err
}

// Streams Engagement records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEngagement(FlowSid string, params *ListEngagementParams) (chan StudioV1Engagement, error) {
	if params == nil {
		params = &ListEngagementParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEngagement(FlowSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan StudioV1Engagement, 1)

	go func() {
		for response != nil {
			for item := range response.Engagements {
				channel <- response.Engagements[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEngagementResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEngagementResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEngagementResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEngagementResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
