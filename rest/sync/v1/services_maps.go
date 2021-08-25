/*
 * Twilio - Sync
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

// Optional parameters for the method 'CreateSyncMap'
type CreateSyncMapParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateSyncMapParams) SetCollectionTtl(CollectionTtl int) *CreateSyncMapParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *CreateSyncMapParams) SetTtl(Ttl int) *CreateSyncMapParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateSyncMapParams) SetUniqueName(UniqueName string) *CreateSyncMapParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateSyncMap(ServiceSid string, params *CreateSyncMapParams) (*SyncV1SyncMap, error) {
	path := "/v1/Services/{ServiceSid}/Maps"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMap{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteSyncMap(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Maps/{Sid}"
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

func (c *ApiService) FetchSyncMap(ServiceSid string, Sid string) (*SyncV1SyncMap, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMap{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncMap'
type ListSyncMapParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncMapParams) SetPageSize(PageSize int) *ListSyncMapParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncMapParams) SetLimit(Limit int) *ListSyncMapParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncMap records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMap(ServiceSid string, params *ListSyncMapParams, pageToken string, pageNumber string) (*ListSyncMapResponse, error) {
	path := "/v1/Services/{ServiceSid}/Maps"

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

	ps := &ListSyncMapResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncMap records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMap(ServiceSid string, params *ListSyncMapParams) ([]SyncV1SyncMap, error) {
	if params == nil {
		params = &ListSyncMapParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncMap(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1SyncMap

	for response != nil {
		records = append(records, response.Maps...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncMapResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSyncMapResponse)
	}

	return records, err
}

// Streams SyncMap records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMap(ServiceSid string, params *ListSyncMapParams) (chan SyncV1SyncMap, error) {
	if params == nil {
		params = &ListSyncMapParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncMap(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncMap, 1)

	go func() {
		for response != nil {
			for item := range response.Maps {
				channel <- response.Maps[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncMapResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncMapResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncMapResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncMapResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncMap'
type UpdateSyncMapParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateSyncMapParams) SetCollectionTtl(CollectionTtl int) *UpdateSyncMapParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *UpdateSyncMapParams) SetTtl(Ttl int) *UpdateSyncMapParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) UpdateSyncMap(ServiceSid string, Sid string, params *UpdateSyncMapParams) (*SyncV1SyncMap, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMap{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
