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

// IntegrationSNOWConfigServiceNowCmdbClassMapping struct for IntegrationSNOWConfigServiceNowCmdbClassMapping
type IntegrationSNOWConfigServiceNowCmdbClassMapping struct {
	Id *string `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	NowClass *string `json:"nowClass,omitempty"`
}

// NewIntegrationSNOWConfigServiceNowCmdbClassMapping instantiates a new IntegrationSNOWConfigServiceNowCmdbClassMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationSNOWConfigServiceNowCmdbClassMapping() *IntegrationSNOWConfigServiceNowCmdbClassMapping {
	this := IntegrationSNOWConfigServiceNowCmdbClassMapping{}
	return &this
}

// NewIntegrationSNOWConfigServiceNowCmdbClassMappingWithDefaults instantiates a new IntegrationSNOWConfigServiceNowCmdbClassMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationSNOWConfigServiceNowCmdbClassMappingWithDefaults() *IntegrationSNOWConfigServiceNowCmdbClassMapping {
	this := IntegrationSNOWConfigServiceNowCmdbClassMapping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) SetName(v string) {
	o.Name = &v
}

// GetNowClass returns the NowClass field value if set, zero value otherwise.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetNowClass() string {
	if o == nil || o.NowClass == nil {
		var ret string
		return ret
	}
	return *o.NowClass
}

// GetNowClassOk returns a tuple with the NowClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) GetNowClassOk() (*string, bool) {
	if o == nil || o.NowClass == nil {
		return nil, false
	}
	return o.NowClass, true
}

// HasNowClass returns a boolean if a field has been set.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) HasNowClass() bool {
	if o != nil && o.NowClass != nil {
		return true
	}

	return false
}

// SetNowClass gets a reference to the given string and assigns it to the NowClass field.
func (o *IntegrationSNOWConfigServiceNowCmdbClassMapping) SetNowClass(v string) {
	o.NowClass = &v
}

func (o IntegrationSNOWConfigServiceNowCmdbClassMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NowClass != nil {
		toSerialize["nowClass"] = o.NowClass
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationSNOWConfigServiceNowCmdbClassMapping struct {
	value *IntegrationSNOWConfigServiceNowCmdbClassMapping
	isSet bool
}

func (v NullableIntegrationSNOWConfigServiceNowCmdbClassMapping) Get() *IntegrationSNOWConfigServiceNowCmdbClassMapping {
	return v.value
}

func (v *NullableIntegrationSNOWConfigServiceNowCmdbClassMapping) Set(val *IntegrationSNOWConfigServiceNowCmdbClassMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationSNOWConfigServiceNowCmdbClassMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationSNOWConfigServiceNowCmdbClassMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationSNOWConfigServiceNowCmdbClassMapping(val *IntegrationSNOWConfigServiceNowCmdbClassMapping) *NullableIntegrationSNOWConfigServiceNowCmdbClassMapping {
	return &NullableIntegrationSNOWConfigServiceNowCmdbClassMapping{value: val, isSet: true}
}

func (v NullableIntegrationSNOWConfigServiceNowCmdbClassMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationSNOWConfigServiceNowCmdbClassMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


