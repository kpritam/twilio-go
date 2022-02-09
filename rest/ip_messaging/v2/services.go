/*
 * Twilio - Ip_messaging
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
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateService(params *CreateServiceParams) (*IpMessagingV2Service, error) {
	path := "/v2/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteService(Sid string) error {
	path := "/v2/Services/{Sid}"
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

func (c *ApiService) FetchService(Sid string) (*IpMessagingV2Service, error) {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2Service{}
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
	path := "/v2/Services"

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
func (c *ApiService) ListService(params *ListServiceParams) ([]IpMessagingV2Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV2Service

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
func (c *ApiService) StreamService(params *ListServiceParams) (chan IpMessagingV2Service, error) {
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
	channel := make(chan IpMessagingV2Service, 1)

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
	//
	ConsumptionReportInterval *int `json:"ConsumptionReportInterval,omitempty"`
	//
	DefaultChannelCreatorRoleSid *string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	//
	DefaultChannelRoleSid *string `json:"DefaultChannelRoleSid,omitempty"`
	//
	DefaultServiceRoleSid *string `json:"DefaultServiceRoleSid,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	LimitsChannelMembers *int `json:"Limits.ChannelMembers,omitempty"`
	//
	LimitsUserChannels *int `json:"Limits.UserChannels,omitempty"`
	//
	MediaCompatibilityMessage *string `json:"Media.CompatibilityMessage,omitempty"`
	//
	NotificationsAddedToChannelEnabled *bool `json:"Notifications.AddedToChannel.Enabled,omitempty"`
	//
	NotificationsAddedToChannelSound *string `json:"Notifications.AddedToChannel.Sound,omitempty"`
	//
	NotificationsAddedToChannelTemplate *string `json:"Notifications.AddedToChannel.Template,omitempty"`
	//
	NotificationsInvitedToChannelEnabled *bool `json:"Notifications.InvitedToChannel.Enabled,omitempty"`
	//
	NotificationsInvitedToChannelSound *string `json:"Notifications.InvitedToChannel.Sound,omitempty"`
	//
	NotificationsInvitedToChannelTemplate *string `json:"Notifications.InvitedToChannel.Template,omitempty"`
	//
	NotificationsLogEnabled *bool `json:"Notifications.LogEnabled,omitempty"`
	//
	NotificationsNewMessageBadgeCountEnabled *bool `json:"Notifications.NewMessage.BadgeCountEnabled,omitempty"`
	//
	NotificationsNewMessageEnabled *bool `json:"Notifications.NewMessage.Enabled,omitempty"`
	//
	NotificationsNewMessageSound *string `json:"Notifications.NewMessage.Sound,omitempty"`
	//
	NotificationsNewMessageTemplate *string `json:"Notifications.NewMessage.Template,omitempty"`
	//
	NotificationsRemovedFromChannelEnabled *bool `json:"Notifications.RemovedFromChannel.Enabled,omitempty"`
	//
	NotificationsRemovedFromChannelSound *string `json:"Notifications.RemovedFromChannel.Sound,omitempty"`
	//
	NotificationsRemovedFromChannelTemplate *string `json:"Notifications.RemovedFromChannel.Template,omitempty"`
	//
	PostWebhookRetryCount *int `json:"PostWebhookRetryCount,omitempty"`
	//
	PostWebhookUrl *string `json:"PostWebhookUrl,omitempty"`
	//
	PreWebhookRetryCount *int `json:"PreWebhookRetryCount,omitempty"`
	//
	PreWebhookUrl *string `json:"PreWebhookUrl,omitempty"`
	//
	ReachabilityEnabled *bool `json:"ReachabilityEnabled,omitempty"`
	//
	ReadStatusEnabled *bool `json:"ReadStatusEnabled,omitempty"`
	//
	TypingIndicatorTimeout *int `json:"TypingIndicatorTimeout,omitempty"`
	//
	WebhookFilters *[]string `json:"WebhookFilters,omitempty"`
	//
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
}

func (params *UpdateServiceParams) SetConsumptionReportInterval(ConsumptionReportInterval int) *UpdateServiceParams {
	params.ConsumptionReportInterval = &ConsumptionReportInterval
	return params
}
func (params *UpdateServiceParams) SetDefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid string) *UpdateServiceParams {
	params.DefaultChannelCreatorRoleSid = &DefaultChannelCreatorRoleSid
	return params
}
func (params *UpdateServiceParams) SetDefaultChannelRoleSid(DefaultChannelRoleSid string) *UpdateServiceParams {
	params.DefaultChannelRoleSid = &DefaultChannelRoleSid
	return params
}
func (params *UpdateServiceParams) SetDefaultServiceRoleSid(DefaultServiceRoleSid string) *UpdateServiceParams {
	params.DefaultServiceRoleSid = &DefaultServiceRoleSid
	return params
}
func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetLimitsChannelMembers(LimitsChannelMembers int) *UpdateServiceParams {
	params.LimitsChannelMembers = &LimitsChannelMembers
	return params
}
func (params *UpdateServiceParams) SetLimitsUserChannels(LimitsUserChannels int) *UpdateServiceParams {
	params.LimitsUserChannels = &LimitsUserChannels
	return params
}
func (params *UpdateServiceParams) SetMediaCompatibilityMessage(MediaCompatibilityMessage string) *UpdateServiceParams {
	params.MediaCompatibilityMessage = &MediaCompatibilityMessage
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsAddedToChannelEnabled = &NotificationsAddedToChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelSound(NotificationsAddedToChannelSound string) *UpdateServiceParams {
	params.NotificationsAddedToChannelSound = &NotificationsAddedToChannelSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate string) *UpdateServiceParams {
	params.NotificationsAddedToChannelTemplate = &NotificationsAddedToChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsInvitedToChannelEnabled = &NotificationsInvitedToChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelSound(NotificationsInvitedToChannelSound string) *UpdateServiceParams {
	params.NotificationsInvitedToChannelSound = &NotificationsInvitedToChannelSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate string) *UpdateServiceParams {
	params.NotificationsInvitedToChannelTemplate = &NotificationsInvitedToChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsLogEnabled(NotificationsLogEnabled bool) *UpdateServiceParams {
	params.NotificationsLogEnabled = &NotificationsLogEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageBadgeCountEnabled(NotificationsNewMessageBadgeCountEnabled bool) *UpdateServiceParams {
	params.NotificationsNewMessageBadgeCountEnabled = &NotificationsNewMessageBadgeCountEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageEnabled(NotificationsNewMessageEnabled bool) *UpdateServiceParams {
	params.NotificationsNewMessageEnabled = &NotificationsNewMessageEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageSound(NotificationsNewMessageSound string) *UpdateServiceParams {
	params.NotificationsNewMessageSound = &NotificationsNewMessageSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageTemplate(NotificationsNewMessageTemplate string) *UpdateServiceParams {
	params.NotificationsNewMessageTemplate = &NotificationsNewMessageTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelEnabled = &NotificationsRemovedFromChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelSound(NotificationsRemovedFromChannelSound string) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelSound = &NotificationsRemovedFromChannelSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate string) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelTemplate = &NotificationsRemovedFromChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetPostWebhookRetryCount(PostWebhookRetryCount int) *UpdateServiceParams {
	params.PostWebhookRetryCount = &PostWebhookRetryCount
	return params
}
func (params *UpdateServiceParams) SetPostWebhookUrl(PostWebhookUrl string) *UpdateServiceParams {
	params.PostWebhookUrl = &PostWebhookUrl
	return params
}
func (params *UpdateServiceParams) SetPreWebhookRetryCount(PreWebhookRetryCount int) *UpdateServiceParams {
	params.PreWebhookRetryCount = &PreWebhookRetryCount
	return params
}
func (params *UpdateServiceParams) SetPreWebhookUrl(PreWebhookUrl string) *UpdateServiceParams {
	params.PreWebhookUrl = &PreWebhookUrl
	return params
}
func (params *UpdateServiceParams) SetReachabilityEnabled(ReachabilityEnabled bool) *UpdateServiceParams {
	params.ReachabilityEnabled = &ReachabilityEnabled
	return params
}
func (params *UpdateServiceParams) SetReadStatusEnabled(ReadStatusEnabled bool) *UpdateServiceParams {
	params.ReadStatusEnabled = &ReadStatusEnabled
	return params
}
func (params *UpdateServiceParams) SetTypingIndicatorTimeout(TypingIndicatorTimeout int) *UpdateServiceParams {
	params.TypingIndicatorTimeout = &TypingIndicatorTimeout
	return params
}
func (params *UpdateServiceParams) SetWebhookFilters(WebhookFilters []string) *UpdateServiceParams {
	params.WebhookFilters = &WebhookFilters
	return params
}
func (params *UpdateServiceParams) SetWebhookMethod(WebhookMethod string) *UpdateServiceParams {
	params.WebhookMethod = &WebhookMethod
	return params
}

func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*IpMessagingV2Service, error) {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConsumptionReportInterval != nil {
		data.Set("ConsumptionReportInterval", fmt.Sprint(*params.ConsumptionReportInterval))
	}
	if params != nil && params.DefaultChannelCreatorRoleSid != nil {
		data.Set("DefaultChannelCreatorRoleSid", *params.DefaultChannelCreatorRoleSid)
	}
	if params != nil && params.DefaultChannelRoleSid != nil {
		data.Set("DefaultChannelRoleSid", *params.DefaultChannelRoleSid)
	}
	if params != nil && params.DefaultServiceRoleSid != nil {
		data.Set("DefaultServiceRoleSid", *params.DefaultServiceRoleSid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.LimitsChannelMembers != nil {
		data.Set("Limits.ChannelMembers", fmt.Sprint(*params.LimitsChannelMembers))
	}
	if params != nil && params.LimitsUserChannels != nil {
		data.Set("Limits.UserChannels", fmt.Sprint(*params.LimitsUserChannels))
	}
	if params != nil && params.MediaCompatibilityMessage != nil {
		data.Set("Media.CompatibilityMessage", *params.MediaCompatibilityMessage)
	}
	if params != nil && params.NotificationsAddedToChannelEnabled != nil {
		data.Set("Notifications.AddedToChannel.Enabled", fmt.Sprint(*params.NotificationsAddedToChannelEnabled))
	}
	if params != nil && params.NotificationsAddedToChannelSound != nil {
		data.Set("Notifications.AddedToChannel.Sound", *params.NotificationsAddedToChannelSound)
	}
	if params != nil && params.NotificationsAddedToChannelTemplate != nil {
		data.Set("Notifications.AddedToChannel.Template", *params.NotificationsAddedToChannelTemplate)
	}
	if params != nil && params.NotificationsInvitedToChannelEnabled != nil {
		data.Set("Notifications.InvitedToChannel.Enabled", fmt.Sprint(*params.NotificationsInvitedToChannelEnabled))
	}
	if params != nil && params.NotificationsInvitedToChannelSound != nil {
		data.Set("Notifications.InvitedToChannel.Sound", *params.NotificationsInvitedToChannelSound)
	}
	if params != nil && params.NotificationsInvitedToChannelTemplate != nil {
		data.Set("Notifications.InvitedToChannel.Template", *params.NotificationsInvitedToChannelTemplate)
	}
	if params != nil && params.NotificationsLogEnabled != nil {
		data.Set("Notifications.LogEnabled", fmt.Sprint(*params.NotificationsLogEnabled))
	}
	if params != nil && params.NotificationsNewMessageBadgeCountEnabled != nil {
		data.Set("Notifications.NewMessage.BadgeCountEnabled", fmt.Sprint(*params.NotificationsNewMessageBadgeCountEnabled))
	}
	if params != nil && params.NotificationsNewMessageEnabled != nil {
		data.Set("Notifications.NewMessage.Enabled", fmt.Sprint(*params.NotificationsNewMessageEnabled))
	}
	if params != nil && params.NotificationsNewMessageSound != nil {
		data.Set("Notifications.NewMessage.Sound", *params.NotificationsNewMessageSound)
	}
	if params != nil && params.NotificationsNewMessageTemplate != nil {
		data.Set("Notifications.NewMessage.Template", *params.NotificationsNewMessageTemplate)
	}
	if params != nil && params.NotificationsRemovedFromChannelEnabled != nil {
		data.Set("Notifications.RemovedFromChannel.Enabled", fmt.Sprint(*params.NotificationsRemovedFromChannelEnabled))
	}
	if params != nil && params.NotificationsRemovedFromChannelSound != nil {
		data.Set("Notifications.RemovedFromChannel.Sound", *params.NotificationsRemovedFromChannelSound)
	}
	if params != nil && params.NotificationsRemovedFromChannelTemplate != nil {
		data.Set("Notifications.RemovedFromChannel.Template", *params.NotificationsRemovedFromChannelTemplate)
	}
	if params != nil && params.PostWebhookRetryCount != nil {
		data.Set("PostWebhookRetryCount", fmt.Sprint(*params.PostWebhookRetryCount))
	}
	if params != nil && params.PostWebhookUrl != nil {
		data.Set("PostWebhookUrl", *params.PostWebhookUrl)
	}
	if params != nil && params.PreWebhookRetryCount != nil {
		data.Set("PreWebhookRetryCount", fmt.Sprint(*params.PreWebhookRetryCount))
	}
	if params != nil && params.PreWebhookUrl != nil {
		data.Set("PreWebhookUrl", *params.PreWebhookUrl)
	}
	if params != nil && params.ReachabilityEnabled != nil {
		data.Set("ReachabilityEnabled", fmt.Sprint(*params.ReachabilityEnabled))
	}
	if params != nil && params.ReadStatusEnabled != nil {
		data.Set("ReadStatusEnabled", fmt.Sprint(*params.ReadStatusEnabled))
	}
	if params != nil && params.TypingIndicatorTimeout != nil {
		data.Set("TypingIndicatorTimeout", fmt.Sprint(*params.TypingIndicatorTimeout))
	}
	if params != nil && params.WebhookFilters != nil {
		for _, item := range *params.WebhookFilters {
			data.Add("WebhookFilters", item)
		}
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
