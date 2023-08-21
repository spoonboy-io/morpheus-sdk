/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiMonitoringContactsContact Payload for creating a new monitoring contact
type ApiMonitoringContactsContact struct {
	// Unique name scoped to your account for the contact
	Name string `json:"name"`
	// Email notification address
	EmailAddress *string `json:"emailAddress,omitempty"`
	// SMS notification address
	SmsAddress *string `json:"smsAddress,omitempty"`
	// Slack Hook
	SlackHook *string `json:"slackHook,omitempty"`
}

// NewApiMonitoringContactsContact instantiates a new ApiMonitoringContactsContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitoringContactsContact(name string, ) *ApiMonitoringContactsContact {
	this := ApiMonitoringContactsContact{}
	this.Name = name
	return &this
}

// NewApiMonitoringContactsContactWithDefaults instantiates a new ApiMonitoringContactsContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitoringContactsContactWithDefaults() *ApiMonitoringContactsContact {
	this := ApiMonitoringContactsContact{}
	return &this
}

// GetName returns the Name field value
func (o *ApiMonitoringContactsContact) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiMonitoringContactsContact) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiMonitoringContactsContact) SetName(v string) {
	o.Name = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ApiMonitoringContactsContact) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringContactsContact) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ApiMonitoringContactsContact) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ApiMonitoringContactsContact) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetSmsAddress returns the SmsAddress field value if set, zero value otherwise.
func (o *ApiMonitoringContactsContact) GetSmsAddress() string {
	if o == nil || o.SmsAddress == nil {
		var ret string
		return ret
	}
	return *o.SmsAddress
}

// GetSmsAddressOk returns a tuple with the SmsAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringContactsContact) GetSmsAddressOk() (*string, bool) {
	if o == nil || o.SmsAddress == nil {
		return nil, false
	}
	return o.SmsAddress, true
}

// HasSmsAddress returns a boolean if a field has been set.
func (o *ApiMonitoringContactsContact) HasSmsAddress() bool {
	if o != nil && o.SmsAddress != nil {
		return true
	}

	return false
}

// SetSmsAddress gets a reference to the given string and assigns it to the SmsAddress field.
func (o *ApiMonitoringContactsContact) SetSmsAddress(v string) {
	o.SmsAddress = &v
}

// GetSlackHook returns the SlackHook field value if set, zero value otherwise.
func (o *ApiMonitoringContactsContact) GetSlackHook() string {
	if o == nil || o.SlackHook == nil {
		var ret string
		return ret
	}
	return *o.SlackHook
}

// GetSlackHookOk returns a tuple with the SlackHook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringContactsContact) GetSlackHookOk() (*string, bool) {
	if o == nil || o.SlackHook == nil {
		return nil, false
	}
	return o.SlackHook, true
}

// HasSlackHook returns a boolean if a field has been set.
func (o *ApiMonitoringContactsContact) HasSlackHook() bool {
	if o != nil && o.SlackHook != nil {
		return true
	}

	return false
}

// SetSlackHook gets a reference to the given string and assigns it to the SlackHook field.
func (o *ApiMonitoringContactsContact) SetSlackHook(v string) {
	o.SlackHook = &v
}

func (o ApiMonitoringContactsContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.SmsAddress != nil {
		toSerialize["smsAddress"] = o.SmsAddress
	}
	if o.SlackHook != nil {
		toSerialize["slackHook"] = o.SlackHook
	}
	return json.Marshal(toSerialize)
}

type NullableApiMonitoringContactsContact struct {
	value *ApiMonitoringContactsContact
	isSet bool
}

func (v NullableApiMonitoringContactsContact) Get() *ApiMonitoringContactsContact {
	return v.value
}

func (v *NullableApiMonitoringContactsContact) Set(val *ApiMonitoringContactsContact) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitoringContactsContact) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitoringContactsContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitoringContactsContact(val *ApiMonitoringContactsContact) *NullableApiMonitoringContactsContact {
	return &NullableApiMonitoringContactsContact{value: val, isSet: true}
}

func (v NullableApiMonitoringContactsContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitoringContactsContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

