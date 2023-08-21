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

// UserSourceCreateOkta Okta Configuration
type UserSourceCreateOkta struct {
	// Okta URL
	Url *string `json:"url,omitempty"`
	// Administrator API Token
	AdministratorAPIToken *string `json:"administratorAPIToken,omitempty"`
	// Required Group
	RequiredGroup *string `json:"requiredGroup,omitempty"`
}

// NewUserSourceCreateOkta instantiates a new UserSourceCreateOkta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceCreateOkta() *UserSourceCreateOkta {
	this := UserSourceCreateOkta{}
	return &this
}

// NewUserSourceCreateOktaWithDefaults instantiates a new UserSourceCreateOkta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceCreateOktaWithDefaults() *UserSourceCreateOkta {
	this := UserSourceCreateOkta{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserSourceCreateOkta) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateOkta) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserSourceCreateOkta) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserSourceCreateOkta) SetUrl(v string) {
	o.Url = &v
}

// GetAdministratorAPIToken returns the AdministratorAPIToken field value if set, zero value otherwise.
func (o *UserSourceCreateOkta) GetAdministratorAPIToken() string {
	if o == nil || o.AdministratorAPIToken == nil {
		var ret string
		return ret
	}
	return *o.AdministratorAPIToken
}

// GetAdministratorAPITokenOk returns a tuple with the AdministratorAPIToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateOkta) GetAdministratorAPITokenOk() (*string, bool) {
	if o == nil || o.AdministratorAPIToken == nil {
		return nil, false
	}
	return o.AdministratorAPIToken, true
}

// HasAdministratorAPIToken returns a boolean if a field has been set.
func (o *UserSourceCreateOkta) HasAdministratorAPIToken() bool {
	if o != nil && o.AdministratorAPIToken != nil {
		return true
	}

	return false
}

// SetAdministratorAPIToken gets a reference to the given string and assigns it to the AdministratorAPIToken field.
func (o *UserSourceCreateOkta) SetAdministratorAPIToken(v string) {
	o.AdministratorAPIToken = &v
}

// GetRequiredGroup returns the RequiredGroup field value if set, zero value otherwise.
func (o *UserSourceCreateOkta) GetRequiredGroup() string {
	if o == nil || o.RequiredGroup == nil {
		var ret string
		return ret
	}
	return *o.RequiredGroup
}

// GetRequiredGroupOk returns a tuple with the RequiredGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateOkta) GetRequiredGroupOk() (*string, bool) {
	if o == nil || o.RequiredGroup == nil {
		return nil, false
	}
	return o.RequiredGroup, true
}

// HasRequiredGroup returns a boolean if a field has been set.
func (o *UserSourceCreateOkta) HasRequiredGroup() bool {
	if o != nil && o.RequiredGroup != nil {
		return true
	}

	return false
}

// SetRequiredGroup gets a reference to the given string and assigns it to the RequiredGroup field.
func (o *UserSourceCreateOkta) SetRequiredGroup(v string) {
	o.RequiredGroup = &v
}

func (o UserSourceCreateOkta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.AdministratorAPIToken != nil {
		toSerialize["administratorAPIToken"] = o.AdministratorAPIToken
	}
	if o.RequiredGroup != nil {
		toSerialize["requiredGroup"] = o.RequiredGroup
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceCreateOkta struct {
	value *UserSourceCreateOkta
	isSet bool
}

func (v NullableUserSourceCreateOkta) Get() *UserSourceCreateOkta {
	return v.value
}

func (v *NullableUserSourceCreateOkta) Set(val *UserSourceCreateOkta) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceCreateOkta) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceCreateOkta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceCreateOkta(val *UserSourceCreateOkta) *NullableUserSourceCreateOkta {
	return &NullableUserSourceCreateOkta{value: val, isSet: true}
}

func (v NullableUserSourceCreateOkta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceCreateOkta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


