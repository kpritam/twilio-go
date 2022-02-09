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

// Optional parameters for the method 'CreateCommand'
type CreateCommandParams struct {
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` after we have sent the command.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The message body of the command.
	Command *string `json:"Command,omitempty"`
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
	Sim *string `json:"Sim,omitempty"`
}

func (params *CreateCommandParams) SetCallbackMethod(CallbackMethod string) *CreateCommandParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateCommandParams) SetCallbackUrl(CallbackUrl string) *CreateCommandParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateCommandParams) SetCommand(Command string) *CreateCommandParams {
	params.Command = &Command
	return params
}
func (params *CreateCommandParams) SetSim(Sim string) *CreateCommandParams {
	params.Sim = &Sim
	return params
}

// Send a Command to a Sim.
func (c *ApiService) CreateCommand(params *CreateCommandParams) (*SupersimV1Command, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Command != nil {
		data.Set("Command", *params.Command)
	}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a Command instance from your account.
func (c *ApiService) FetchCommand(Sid string) (*SupersimV1Command, error) {
	path := "/v1/Commands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCommand'
type ListCommandParams struct {
	// The SID or unique name of the Sim that Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`
	// The status of the Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each.
	Status *string `json:"Status,omitempty"`
	// The direction of the Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCommandParams) SetSim(Sim string) *ListCommandParams {
	params.Sim = &Sim
	return params
}
func (params *ListCommandParams) SetStatus(Status string) *ListCommandParams {
	params.Status = &Status
	return params
}
func (params *ListCommandParams) SetDirection(Direction string) *ListCommandParams {
	params.Direction = &Direction
	return params
}
func (params *ListCommandParams) SetPageSize(PageSize int) *ListCommandParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCommandParams) SetLimit(Limit int) *ListCommandParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Command records from the API. Request is executed immediately.
func (c *ApiService) PageCommand(params *ListCommandParams, pageToken, pageNumber string) (*ListCommandResponse, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
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

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Command records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCommand(params *ListCommandParams) ([]SupersimV1Command, error) {
	if params == nil {
		params = &ListCommandParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCommand(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SupersimV1Command

	for response != nil {
		records = append(records, response.Commands...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCommandResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCommandResponse)
	}

	return records, err
}

// Streams Command records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCommand(params *ListCommandParams) (chan SupersimV1Command, error) {
	if params == nil {
		params = &ListCommandParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCommand(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SupersimV1Command, 1)

	go func() {
		for response != nil {
			for item := range response.Commands {
				channel <- response.Commands[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCommandResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCommandResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCommandResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
