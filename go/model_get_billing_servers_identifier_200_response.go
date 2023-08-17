/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetBillingServersIdentifier200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBillingServersIdentifier200Response{}

// GetBillingServersIdentifier200Response struct for GetBillingServersIdentifier200Response
type GetBillingServersIdentifier200Response struct {
	Success *bool `json:"success,omitempty"`
	BillingInfo *BillingServer `json:"billingInfo,omitempty"`
}

// NewGetBillingServersIdentifier200Response instantiates a new GetBillingServersIdentifier200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBillingServersIdentifier200Response() *GetBillingServersIdentifier200Response {
	this := GetBillingServersIdentifier200Response{}
	return &this
}

// NewGetBillingServersIdentifier200ResponseWithDefaults instantiates a new GetBillingServersIdentifier200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBillingServersIdentifier200ResponseWithDefaults() *GetBillingServersIdentifier200Response {
	this := GetBillingServersIdentifier200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetBillingServersIdentifier200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingServersIdentifier200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetBillingServersIdentifier200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetBillingServersIdentifier200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise.
func (o *GetBillingServersIdentifier200Response) GetBillingInfo() BillingServer {
	if o == nil || IsNil(o.BillingInfo) {
		var ret BillingServer
		return ret
	}
	return *o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingServersIdentifier200Response) GetBillingInfoOk() (*BillingServer, bool) {
	if o == nil || IsNil(o.BillingInfo) {
		return nil, false
	}
	return o.BillingInfo, true
}

// HasBillingInfo returns a boolean if a field has been set.
func (o *GetBillingServersIdentifier200Response) HasBillingInfo() bool {
	if o != nil && !IsNil(o.BillingInfo) {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given BillingServer and assigns it to the BillingInfo field.
func (o *GetBillingServersIdentifier200Response) SetBillingInfo(v BillingServer) {
	o.BillingInfo = &v
}

func (o GetBillingServersIdentifier200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBillingServersIdentifier200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.BillingInfo) {
		toSerialize["billingInfo"] = o.BillingInfo
	}
	return toSerialize, nil
}

type NullableGetBillingServersIdentifier200Response struct {
	value *GetBillingServersIdentifier200Response
	isSet bool
}

func (v NullableGetBillingServersIdentifier200Response) Get() *GetBillingServersIdentifier200Response {
	return v.value
}

func (v *NullableGetBillingServersIdentifier200Response) Set(val *GetBillingServersIdentifier200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBillingServersIdentifier200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBillingServersIdentifier200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBillingServersIdentifier200Response(val *GetBillingServersIdentifier200Response) *NullableGetBillingServersIdentifier200Response {
	return &NullableGetBillingServersIdentifier200Response{value: val, isSet: true}
}

func (v NullableGetBillingServersIdentifier200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBillingServersIdentifier200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


