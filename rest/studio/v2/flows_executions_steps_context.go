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
	"net/url"
	"strings"
)

// Retrieve the context for an Execution Step.
func (c *ApiService) FetchExecutionStepContext(FlowSid string, ExecutionSid string, StepSid string) (*StudioV2ExecutionStepContext, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
	path = strings.Replace(path, "{"+"StepSid"+"}", StepSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2ExecutionStepContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
