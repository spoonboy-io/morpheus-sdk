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

// InlineObject273 struct for InlineObject273
type InlineObject273 struct {
	Page *ApiAppsIdWikiPage `json:"page,omitempty"`
}

// NewInlineObject273 instantiates a new InlineObject273 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject273() *InlineObject273 {
	this := InlineObject273{}
	return &this
}

// NewInlineObject273WithDefaults instantiates a new InlineObject273 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject273WithDefaults() *InlineObject273 {
	this := InlineObject273{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *InlineObject273) GetPage() ApiAppsIdWikiPage {
	if o == nil || o.Page == nil {
		var ret ApiAppsIdWikiPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject273) GetPageOk() (*ApiAppsIdWikiPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *InlineObject273) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given ApiAppsIdWikiPage and assigns it to the Page field.
func (o *InlineObject273) SetPage(v ApiAppsIdWikiPage) {
	o.Page = &v
}

func (o InlineObject273) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject273 struct {
	value *InlineObject273
	isSet bool
}

func (v NullableInlineObject273) Get() *InlineObject273 {
	return v.value
}

func (v *NullableInlineObject273) Set(val *InlineObject273) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject273) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject273) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject273(val *InlineObject273) *NullableInlineObject273 {
	return &NullableInlineObject273{value: val, isSet: true}
}

func (v NullableInlineObject273) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject273) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

