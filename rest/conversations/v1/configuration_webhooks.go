/*
 * Twilio - Conversations
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
)

func (c *ApiService) FetchConfigurationWebhook() (*ConversationsV1ConfigurationWebhook, error) {
	path := "/v1/Configuration/Webhooks"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateConfigurationWebhook'
type UpdateConfigurationWebhookParams struct {
	// The list of webhook event triggers that are enabled for this Service: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`
	Filters *[]string `json:"Filters,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	Method *string `json:"Method,omitempty"`
	// The absolute url the post-event webhook request should be sent to.
	PostWebhookUrl *string `json:"PostWebhookUrl,omitempty"`
	// The absolute url the pre-event webhook request should be sent to.
	PreWebhookUrl *string `json:"PreWebhookUrl,omitempty"`
	// The routing target of the webhook.
	Target *string `json:"Target,omitempty"`
}

func (params *UpdateConfigurationWebhookParams) SetFilters(Filters []string) *UpdateConfigurationWebhookParams {
	params.Filters = &Filters
	return params
}
func (params *UpdateConfigurationWebhookParams) SetMethod(Method string) *UpdateConfigurationWebhookParams {
	params.Method = &Method
	return params
}
func (params *UpdateConfigurationWebhookParams) SetPostWebhookUrl(PostWebhookUrl string) *UpdateConfigurationWebhookParams {
	params.PostWebhookUrl = &PostWebhookUrl
	return params
}
func (params *UpdateConfigurationWebhookParams) SetPreWebhookUrl(PreWebhookUrl string) *UpdateConfigurationWebhookParams {
	params.PreWebhookUrl = &PreWebhookUrl
	return params
}
func (params *UpdateConfigurationWebhookParams) SetTarget(Target string) *UpdateConfigurationWebhookParams {
	params.Target = &Target
	return params
}

func (c *ApiService) UpdateConfigurationWebhook(params *UpdateConfigurationWebhookParams) (*ConversationsV1ConfigurationWebhook, error) {
	path := "/v1/Configuration/Webhooks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Filters != nil {
		for _, item := range *params.Filters {
			data.Add("Filters", item)
		}
	}
	if params != nil && params.Method != nil {
		data.Set("Method", *params.Method)
	}
	if params != nil && params.PostWebhookUrl != nil {
		data.Set("PostWebhookUrl", *params.PostWebhookUrl)
	}
	if params != nil && params.PreWebhookUrl != nil {
		data.Set("PreWebhookUrl", *params.PreWebhookUrl)
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
