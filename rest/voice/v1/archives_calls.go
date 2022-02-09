/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.26.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"fmt"
	"net/url"
	"strings"
)

// Delete an archived call record from Bulk Export. Note: this does not also delete the record from the Voice API.
func (c *ApiService) DeleteArchivedCall(Date string, Sid string) error {
	path := "/v1/Archives/{Date}/Calls/{Sid}"
	path = strings.Replace(path, "{"+"Date"+"}", fmt.Sprint(Date), -1)
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
