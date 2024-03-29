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

// WhitelabelSettingsSupportMenuLinks struct for WhitelabelSettingsSupportMenuLinks
type WhitelabelSettingsSupportMenuLinks struct {
	Url *string `json:"url,omitempty"`
	Label *string `json:"label,omitempty"`
	LabelCode *string `json:"labelCode,omitempty"`
}

// NewWhitelabelSettingsSupportMenuLinks instantiates a new WhitelabelSettingsSupportMenuLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhitelabelSettingsSupportMenuLinks() *WhitelabelSettingsSupportMenuLinks {
	this := WhitelabelSettingsSupportMenuLinks{}
	return &this
}

// NewWhitelabelSettingsSupportMenuLinksWithDefaults instantiates a new WhitelabelSettingsSupportMenuLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhitelabelSettingsSupportMenuLinksWithDefaults() *WhitelabelSettingsSupportMenuLinks {
	this := WhitelabelSettingsSupportMenuLinks{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WhitelabelSettingsSupportMenuLinks) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsSupportMenuLinks) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WhitelabelSettingsSupportMenuLinks) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WhitelabelSettingsSupportMenuLinks) SetUrl(v string) {
	o.Url = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WhitelabelSettingsSupportMenuLinks) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsSupportMenuLinks) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WhitelabelSettingsSupportMenuLinks) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WhitelabelSettingsSupportMenuLinks) SetLabel(v string) {
	o.Label = &v
}

// GetLabelCode returns the LabelCode field value if set, zero value otherwise.
func (o *WhitelabelSettingsSupportMenuLinks) GetLabelCode() string {
	if o == nil || o.LabelCode == nil {
		var ret string
		return ret
	}
	return *o.LabelCode
}

// GetLabelCodeOk returns a tuple with the LabelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsSupportMenuLinks) GetLabelCodeOk() (*string, bool) {
	if o == nil || o.LabelCode == nil {
		return nil, false
	}
	return o.LabelCode, true
}

// HasLabelCode returns a boolean if a field has been set.
func (o *WhitelabelSettingsSupportMenuLinks) HasLabelCode() bool {
	if o != nil && o.LabelCode != nil {
		return true
	}

	return false
}

// SetLabelCode gets a reference to the given string and assigns it to the LabelCode field.
func (o *WhitelabelSettingsSupportMenuLinks) SetLabelCode(v string) {
	o.LabelCode = &v
}

func (o WhitelabelSettingsSupportMenuLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LabelCode != nil {
		toSerialize["labelCode"] = o.LabelCode
	}
	return json.Marshal(toSerialize)
}

type NullableWhitelabelSettingsSupportMenuLinks struct {
	value *WhitelabelSettingsSupportMenuLinks
	isSet bool
}

func (v NullableWhitelabelSettingsSupportMenuLinks) Get() *WhitelabelSettingsSupportMenuLinks {
	return v.value
}

func (v *NullableWhitelabelSettingsSupportMenuLinks) Set(val *WhitelabelSettingsSupportMenuLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWhitelabelSettingsSupportMenuLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWhitelabelSettingsSupportMenuLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhitelabelSettingsSupportMenuLinks(val *WhitelabelSettingsSupportMenuLinks) *NullableWhitelabelSettingsSupportMenuLinks {
	return &NullableWhitelabelSettingsSupportMenuLinks{value: val, isSet: true}
}

func (v NullableWhitelabelSettingsSupportMenuLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhitelabelSettingsSupportMenuLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


