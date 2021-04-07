/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://trunking.twilio.com",
	}
}

// CreateCredentialListParams Optional parameters for the method 'CreateCredentialList'
type CreateCredentialListParams struct {
	CredentialListSid *string `json:"CredentialListSid,omitempty"`
}

/*
* CreateCredentialList Method for CreateCredentialList
* @param TrunkSid The SID of the Trunk to associate the credential list with.
* @param optional nil or *CreateCredentialListParams - Optional Parameters:
* @param "CredentialListSid" (string) - The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.
* @return TrunkingV1TrunkCredentialList
 */
func (c *DefaultApiService) CreateCredentialList(TrunkSid string, params *CreateCredentialListParams) (*TrunkingV1TrunkCredentialList, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.CredentialListSid != nil {
		data.Set("CredentialListSid", *params.CredentialListSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateIpAccessControlListParams Optional parameters for the method 'CreateIpAccessControlList'
type CreateIpAccessControlListParams struct {
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
}

/*
* CreateIpAccessControlList Method for CreateIpAccessControlList
* Associate an IP Access Control List with a Trunk
* @param TrunkSid The SID of the Trunk to associate the IP Access Control List with.
* @param optional nil or *CreateIpAccessControlListParams - Optional Parameters:
* @param "IpAccessControlListSid" (string) - The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.
* @return TrunkingV1TrunkIpAccessControlList
 */
func (c *DefaultApiService) CreateIpAccessControlList(TrunkSid string, params *CreateIpAccessControlListParams) (*TrunkingV1TrunkIpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.IpAccessControlListSid != nil {
		data.Set("IpAccessControlListSid", *params.IpAccessControlListSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateOriginationUrlParams Optional parameters for the method 'CreateOriginationUrl'
type CreateOriginationUrlParams struct {
	Enabled      *bool   `json:"Enabled,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	Priority     *int32  `json:"Priority,omitempty"`
	SipUrl       *string `json:"SipUrl,omitempty"`
	Weight       *int32  `json:"Weight,omitempty"`
}

/*
* CreateOriginationUrl Method for CreateOriginationUrl
* @param TrunkSid The SID of the Trunk to associate the resource with.
* @param optional nil or *CreateOriginationUrlParams - Optional Parameters:
* @param "Enabled" (bool) - Whether the URL is enabled. The default is `true`.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
* @param "Priority" (int32) - The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
* @param "SipUrl" (string) - The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema.
* @param "Weight" (int32) - The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
* @return TrunkingV1TrunkOriginationUrl
 */
func (c *DefaultApiService) CreateOriginationUrl(TrunkSid string, params *CreateOriginationUrlParams) (*TrunkingV1TrunkOriginationUrl, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.SipUrl != nil {
		data.Set("SipUrl", *params.SipUrl)
	}
	if params != nil && params.Weight != nil {
		data.Set("Weight", fmt.Sprint(*params.Weight))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkOriginationUrl{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreatePhoneNumberParams Optional parameters for the method 'CreatePhoneNumber'
type CreatePhoneNumberParams struct {
	PhoneNumberSid *string `json:"PhoneNumberSid,omitempty"`
}

/*
* CreatePhoneNumber Method for CreatePhoneNumber
* @param TrunkSid The SID of the Trunk to associate the phone number with.
* @param optional nil or *CreatePhoneNumberParams - Optional Parameters:
* @param "PhoneNumberSid" (string) - The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk.
* @return TrunkingV1TrunkPhoneNumber
 */
func (c *DefaultApiService) CreatePhoneNumber(TrunkSid string, params *CreatePhoneNumberParams) (*TrunkingV1TrunkPhoneNumber, error) {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PhoneNumberSid != nil {
		data.Set("PhoneNumberSid", *params.PhoneNumberSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateTrunkParams Optional parameters for the method 'CreateTrunk'
type CreateTrunkParams struct {
	CnamLookupEnabled      *bool   `json:"CnamLookupEnabled,omitempty"`
	DisasterRecoveryMethod *string `json:"DisasterRecoveryMethod,omitempty"`
	DisasterRecoveryUrl    *string `json:"DisasterRecoveryUrl,omitempty"`
	DomainName             *string `json:"DomainName,omitempty"`
	FriendlyName           *string `json:"FriendlyName,omitempty"`
	Secure                 *bool   `json:"Secure,omitempty"`
	TransferMode           *string `json:"TransferMode,omitempty"`
}

/*
* CreateTrunk Method for CreateTrunk
* @param optional nil or *CreateTrunkParams - Optional Parameters:
* @param "CnamLookupEnabled" (bool) - Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
* @param "DisasterRecoveryMethod" (string) - The HTTP method we should use to call the `disaster_recovery_url`. Can be: `GET` or `POST`.
* @param "DisasterRecoveryUrl" (string) - The URL we should call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
* @param "DomainName" (string) - The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
* @param "Secure" (bool) - Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
* @param "TransferMode" (string) - The call transfer settings for the trunk. Can be: `enable-all`, `sip-only` and `disable-all`. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.
* @return TrunkingV1Trunk
 */
func (c *DefaultApiService) CreateTrunk(params *CreateTrunkParams) (*TrunkingV1Trunk, error) {
	path := "/v1/Trunks"

	data := url.Values{}
	headers := 0

	if params != nil && params.CnamLookupEnabled != nil {
		data.Set("CnamLookupEnabled", fmt.Sprint(*params.CnamLookupEnabled))
	}
	if params != nil && params.DisasterRecoveryMethod != nil {
		data.Set("DisasterRecoveryMethod", *params.DisasterRecoveryMethod)
	}
	if params != nil && params.DisasterRecoveryUrl != nil {
		data.Set("DisasterRecoveryUrl", *params.DisasterRecoveryUrl)
	}
	if params != nil && params.DomainName != nil {
		data.Set("DomainName", *params.DomainName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Secure != nil {
		data.Set("Secure", fmt.Sprint(*params.Secure))
	}
	if params != nil && params.TransferMode != nil {
		data.Set("TransferMode", *params.TransferMode)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Trunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* DeleteCredentialList Method for DeleteCredentialList
* @param TrunkSid The SID of the Trunk from which to delete the credential list.
* @param Sid The unique string that we created to identify the CredentialList resource to delete.
 */
func (c *DefaultApiService) DeleteCredentialList(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeleteIpAccessControlList Method for DeleteIpAccessControlList
* Remove an associated IP Access Control List from a Trunk
* @param TrunkSid The SID of the Trunk from which to delete the IP Access Control List.
* @param Sid The unique string that we created to identify the IpAccessControlList resource to delete.
 */
func (c *DefaultApiService) DeleteIpAccessControlList(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeleteOriginationUrl Method for DeleteOriginationUrl
* @param TrunkSid The SID of the Trunk from which to delete the OriginationUrl.
* @param Sid The unique string that we created to identify the OriginationUrl resource to delete.
 */
func (c *DefaultApiService) DeleteOriginationUrl(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeletePhoneNumber Method for DeletePhoneNumber
* @param TrunkSid The SID of the Trunk from which to delete the PhoneNumber resource.
* @param Sid The unique string that we created to identify the PhoneNumber resource to delete.
 */
func (c *DefaultApiService) DeletePhoneNumber(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* DeleteTrunk Method for DeleteTrunk
* @param Sid The unique string that we created to identify the Trunk resource to delete.
 */
func (c *DefaultApiService) DeleteTrunk(Sid string) error {
	path := "/v1/Trunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* FetchCredentialList Method for FetchCredentialList
* @param TrunkSid The SID of the Trunk from which to fetch the credential list.
* @param Sid The unique string that we created to identify the CredentialList resource to fetch.
* @return TrunkingV1TrunkCredentialList
 */
func (c *DefaultApiService) FetchCredentialList(TrunkSid string, Sid string) (*TrunkingV1TrunkCredentialList, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchIpAccessControlList Method for FetchIpAccessControlList
* @param TrunkSid The SID of the Trunk from which to fetch the IP Access Control List.
* @param Sid The unique string that we created to identify the IpAccessControlList resource to fetch.
* @return TrunkingV1TrunkIpAccessControlList
 */
func (c *DefaultApiService) FetchIpAccessControlList(TrunkSid string, Sid string) (*TrunkingV1TrunkIpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchOriginationUrl Method for FetchOriginationUrl
* @param TrunkSid The SID of the Trunk from which to fetch the OriginationUrl.
* @param Sid The unique string that we created to identify the OriginationUrl resource to fetch.
* @return TrunkingV1TrunkOriginationUrl
 */
func (c *DefaultApiService) FetchOriginationUrl(TrunkSid string, Sid string) (*TrunkingV1TrunkOriginationUrl, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkOriginationUrl{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchPhoneNumber Method for FetchPhoneNumber
* @param TrunkSid The SID of the Trunk from which to fetch the PhoneNumber resource.
* @param Sid The unique string that we created to identify the PhoneNumber resource to fetch.
* @return TrunkingV1TrunkPhoneNumber
 */
func (c *DefaultApiService) FetchPhoneNumber(TrunkSid string, Sid string) (*TrunkingV1TrunkPhoneNumber, error) {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchRecording Method for FetchRecording
* @param TrunkSid The SID of the Trunk from which to fetch the recording settings.
* @return TrunkingV1TrunkRecording
 */
func (c *DefaultApiService) FetchRecording(TrunkSid string) (*TrunkingV1TrunkRecording, error) {
	path := "/v1/Trunks/{TrunkSid}/Recording"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchTrunk Method for FetchTrunk
* @param Sid The unique string that we created to identify the Trunk resource to fetch.
* @return TrunkingV1Trunk
 */
func (c *DefaultApiService) FetchTrunk(Sid string) (*TrunkingV1Trunk, error) {
	path := "/v1/Trunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Trunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListCredentialListParams Optional parameters for the method 'ListCredentialList'
type ListCredentialListParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListCredentialList Method for ListCredentialList
* @param TrunkSid The SID of the Trunk from which to read the credential lists.
* @param optional nil or *ListCredentialListParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListCredentialListResponse
 */
func (c *DefaultApiService) ListCredentialList(TrunkSid string, params *ListCredentialListParams) (*ListCredentialListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListIpAccessControlListParams Optional parameters for the method 'ListIpAccessControlList'
type ListIpAccessControlListParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListIpAccessControlList Method for ListIpAccessControlList
* List all IP Access Control Lists for a Trunk
* @param TrunkSid The SID of the Trunk from which to read the IP Access Control Lists.
* @param optional nil or *ListIpAccessControlListParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListIpAccessControlListResponse
 */
func (c *DefaultApiService) ListIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams) (*ListIpAccessControlListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListOriginationUrlParams Optional parameters for the method 'ListOriginationUrl'
type ListOriginationUrlParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListOriginationUrl Method for ListOriginationUrl
* @param TrunkSid The SID of the Trunk from which to read the OriginationUrl.
* @param optional nil or *ListOriginationUrlParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListOriginationUrlResponse
 */
func (c *DefaultApiService) ListOriginationUrl(TrunkSid string, params *ListOriginationUrlParams) (*ListOriginationUrlResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListOriginationUrlResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListPhoneNumberParams Optional parameters for the method 'ListPhoneNumber'
type ListPhoneNumberParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListPhoneNumber Method for ListPhoneNumber
* @param TrunkSid The SID of the Trunk from which to read the PhoneNumber resources.
* @param optional nil or *ListPhoneNumberParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListPhoneNumberResponse
 */
func (c *DefaultApiService) ListPhoneNumber(TrunkSid string, params *ListPhoneNumberParams) (*ListPhoneNumberResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListTrunkParams Optional parameters for the method 'ListTrunk'
type ListTrunkParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListTrunk Method for ListTrunk
* @param optional nil or *ListTrunkParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListTrunkResponse
 */
func (c *DefaultApiService) ListTrunk(params *ListTrunkParams) (*ListTrunkResponse, error) {
	path := "/v1/Trunks"

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrunkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateOriginationUrlParams Optional parameters for the method 'UpdateOriginationUrl'
type UpdateOriginationUrlParams struct {
	Enabled      *bool   `json:"Enabled,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	Priority     *int32  `json:"Priority,omitempty"`
	SipUrl       *string `json:"SipUrl,omitempty"`
	Weight       *int32  `json:"Weight,omitempty"`
}

/*
* UpdateOriginationUrl Method for UpdateOriginationUrl
* @param TrunkSid The SID of the Trunk from which to update the OriginationUrl.
* @param Sid The unique string that we created to identify the OriginationUrl resource to update.
* @param optional nil or *UpdateOriginationUrlParams - Optional Parameters:
* @param "Enabled" (bool) - Whether the URL is enabled. The default is `true`.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
* @param "Priority" (int32) - The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
* @param "SipUrl" (string) - The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema. `sips` is NOT supported.
* @param "Weight" (int32) - The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
* @return TrunkingV1TrunkOriginationUrl
 */
func (c *DefaultApiService) UpdateOriginationUrl(TrunkSid string, Sid string, params *UpdateOriginationUrlParams) (*TrunkingV1TrunkOriginationUrl, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.SipUrl != nil {
		data.Set("SipUrl", *params.SipUrl)
	}
	if params != nil && params.Weight != nil {
		data.Set("Weight", fmt.Sprint(*params.Weight))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkOriginationUrl{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateRecordingParams Optional parameters for the method 'UpdateRecording'
type UpdateRecordingParams struct {
	Mode *string `json:"Mode,omitempty"`
	Trim *string `json:"Trim,omitempty"`
}

/*
* UpdateRecording Method for UpdateRecording
* @param TrunkSid The SID of the Trunk that will have its recording settings updated.
* @param optional nil or *UpdateRecordingParams - Optional Parameters:
* @param "Mode" (string) - The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual.
* @param "Trim" (string) - The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence.
* @return TrunkingV1TrunkRecording
 */
func (c *DefaultApiService) UpdateRecording(TrunkSid string, params *UpdateRecordingParams) (*TrunkingV1TrunkRecording, error) {
	path := "/v1/Trunks/{TrunkSid}/Recording"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Mode != nil {
		data.Set("Mode", *params.Mode)
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", *params.Trim)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateTrunkParams Optional parameters for the method 'UpdateTrunk'
type UpdateTrunkParams struct {
	CnamLookupEnabled      *bool   `json:"CnamLookupEnabled,omitempty"`
	DisasterRecoveryMethod *string `json:"DisasterRecoveryMethod,omitempty"`
	DisasterRecoveryUrl    *string `json:"DisasterRecoveryUrl,omitempty"`
	DomainName             *string `json:"DomainName,omitempty"`
	FriendlyName           *string `json:"FriendlyName,omitempty"`
	Secure                 *bool   `json:"Secure,omitempty"`
	TransferMode           *string `json:"TransferMode,omitempty"`
}

/*
* UpdateTrunk Method for UpdateTrunk
* @param Sid The unique string that we created to identify the OriginationUrl resource to update.
* @param optional nil or *UpdateTrunkParams - Optional Parameters:
* @param "CnamLookupEnabled" (bool) - Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
* @param "DisasterRecoveryMethod" (string) - The HTTP method we should use to call the `disaster_recovery_url`. Can be: `GET` or `POST`.
* @param "DisasterRecoveryUrl" (string) - The URL we should call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
* @param "DomainName" (string) - The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
* @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
* @param "Secure" (bool) - Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
* @param "TransferMode" (string) - The call transfer settings for the trunk. Can be: `enable-all`, `sip-only` and `disable-all`. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.
* @return TrunkingV1Trunk
 */
func (c *DefaultApiService) UpdateTrunk(Sid string, params *UpdateTrunkParams) (*TrunkingV1Trunk, error) {
	path := "/v1/Trunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.CnamLookupEnabled != nil {
		data.Set("CnamLookupEnabled", fmt.Sprint(*params.CnamLookupEnabled))
	}
	if params != nil && params.DisasterRecoveryMethod != nil {
		data.Set("DisasterRecoveryMethod", *params.DisasterRecoveryMethod)
	}
	if params != nil && params.DisasterRecoveryUrl != nil {
		data.Set("DisasterRecoveryUrl", *params.DisasterRecoveryUrl)
	}
	if params != nil && params.DomainName != nil {
		data.Set("DomainName", *params.DomainName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Secure != nil {
		data.Set("Secure", fmt.Sprint(*params.Secure))
	}
	if params != nil && params.TransferMode != nil {
		data.Set("TransferMode", *params.TransferMode)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Trunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
