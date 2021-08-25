/*
 * Twilio - Api
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

// Optional parameters for the method 'FetchAvailablePhoneNumberCountry'
type FetchAvailablePhoneNumberCountryParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAvailablePhoneNumberCountryParams) SetPathAccountSid(PathAccountSid string) *FetchAvailablePhoneNumberCountryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchAvailablePhoneNumberCountry(CountryCode string, params *FetchAvailablePhoneNumberCountryParams) (*ApiV2010AvailablePhoneNumberCountry, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CountryCode"+"}", CountryCode, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AvailablePhoneNumberCountry{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAvailablePhoneNumberCountry'
type ListAvailablePhoneNumberCountryParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAvailablePhoneNumberCountryParams) SetPathAccountSid(PathAccountSid string) *ListAvailablePhoneNumberCountryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAvailablePhoneNumberCountryParams) SetPageSize(PageSize int) *ListAvailablePhoneNumberCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAvailablePhoneNumberCountryParams) SetLimit(Limit int) *ListAvailablePhoneNumberCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AvailablePhoneNumberCountry records from the API. Request is executed immediately.
func (c *ApiService) PageAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams, pageToken string, pageNumber string) (*ListAvailablePhoneNumberCountryResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListAvailablePhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AvailablePhoneNumberCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams) ([]ApiV2010AvailablePhoneNumberCountry, error) {
	if params == nil {
		params = &ListAvailablePhoneNumberCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAvailablePhoneNumberCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AvailablePhoneNumberCountry

	for response != nil {
		records = append(records, response.Countries...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAvailablePhoneNumberCountryResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAvailablePhoneNumberCountryResponse)
	}

	return records, err
}

// Streams AvailablePhoneNumberCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams) (chan ApiV2010AvailablePhoneNumberCountry, error) {
	if params == nil {
		params = &ListAvailablePhoneNumberCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAvailablePhoneNumberCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AvailablePhoneNumberCountry, 1)

	go func() {
		for response != nil {
			for item := range response.Countries {
				channel <- response.Countries[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAvailablePhoneNumberCountryResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAvailablePhoneNumberCountryResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAvailablePhoneNumberCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAvailablePhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
