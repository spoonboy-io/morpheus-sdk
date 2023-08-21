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

// ApiAppsIdWikiPage struct for ApiAppsIdWikiPage
type ApiAppsIdWikiPage struct {
	Name *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
	Content *string `json:"content,omitempty"`
}

// NewApiAppsIdWikiPage instantiates a new ApiAppsIdWikiPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAppsIdWikiPage() *ApiAppsIdWikiPage {
	this := ApiAppsIdWikiPage{}
	return &this
}

// NewApiAppsIdWikiPageWithDefaults instantiates a new ApiAppsIdWikiPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAppsIdWikiPageWithDefaults() *ApiAppsIdWikiPage {
	this := ApiAppsIdWikiPage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAppsIdWikiPage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppsIdWikiPage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAppsIdWikiPage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAppsIdWikiPage) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ApiAppsIdWikiPage) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppsIdWikiPage) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ApiAppsIdWikiPage) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ApiAppsIdWikiPage) SetCategory(v string) {
	o.Category = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ApiAppsIdWikiPage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppsIdWikiPage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ApiAppsIdWikiPage) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ApiAppsIdWikiPage) SetContent(v string) {
	o.Content = &v
}

func (o ApiAppsIdWikiPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableApiAppsIdWikiPage struct {
	value *ApiAppsIdWikiPage
	isSet bool
}

func (v NullableApiAppsIdWikiPage) Get() *ApiAppsIdWikiPage {
	return v.value
}

func (v *NullableApiAppsIdWikiPage) Set(val *ApiAppsIdWikiPage) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAppsIdWikiPage) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAppsIdWikiPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAppsIdWikiPage(val *ApiAppsIdWikiPage) *NullableApiAppsIdWikiPage {
	return &NullableApiAppsIdWikiPage{value: val, isSet: true}
}

func (v NullableApiAppsIdWikiPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAppsIdWikiPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


