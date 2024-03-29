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

// SpecTemplateCreateConfigCloudformation struct for SpecTemplateCreateConfigCloudformation
type SpecTemplateCreateConfigCloudformation struct {
	IAM OneOfstringboolean `json:"IAM"`
	CAPABILITY_NAMED_IAM OneOfstringboolean `json:"CAPABILITY_NAMED_IAM"`
	CAPABILITY_AUTO_EXPAND OneOfstringboolean `json:"CAPABILITY_AUTO_EXPAND"`
}

// NewSpecTemplateCreateConfigCloudformation instantiates a new SpecTemplateCreateConfigCloudformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecTemplateCreateConfigCloudformation(iAM OneOfstringboolean, cAPABILITYNAMEDIAM OneOfstringboolean, cAPABILITYAUTOEXPAND OneOfstringboolean, ) *SpecTemplateCreateConfigCloudformation {
	this := SpecTemplateCreateConfigCloudformation{}
	this.IAM = iAM
	this.CAPABILITY_NAMED_IAM = cAPABILITYNAMEDIAM
	this.CAPABILITY_AUTO_EXPAND = cAPABILITYAUTOEXPAND
	return &this
}

// NewSpecTemplateCreateConfigCloudformationWithDefaults instantiates a new SpecTemplateCreateConfigCloudformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecTemplateCreateConfigCloudformationWithDefaults() *SpecTemplateCreateConfigCloudformation {
	this := SpecTemplateCreateConfigCloudformation{}
	return &this
}

// GetIAM returns the IAM field value
func (o *SpecTemplateCreateConfigCloudformation) GetIAM() OneOfstringboolean {
	if o == nil  {
		var ret OneOfstringboolean
		return ret
	}

	return o.IAM
}

// GetIAMOk returns a tuple with the IAM field value
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateConfigCloudformation) GetIAMOk() (*OneOfstringboolean, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IAM, true
}

// SetIAM sets field value
func (o *SpecTemplateCreateConfigCloudformation) SetIAM(v OneOfstringboolean) {
	o.IAM = v
}

// GetCAPABILITY_NAMED_IAM returns the CAPABILITY_NAMED_IAM field value
func (o *SpecTemplateCreateConfigCloudformation) GetCAPABILITY_NAMED_IAM() OneOfstringboolean {
	if o == nil  {
		var ret OneOfstringboolean
		return ret
	}

	return o.CAPABILITY_NAMED_IAM
}

// GetCAPABILITY_NAMED_IAMOk returns a tuple with the CAPABILITY_NAMED_IAM field value
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateConfigCloudformation) GetCAPABILITY_NAMED_IAMOk() (*OneOfstringboolean, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CAPABILITY_NAMED_IAM, true
}

// SetCAPABILITY_NAMED_IAM sets field value
func (o *SpecTemplateCreateConfigCloudformation) SetCAPABILITY_NAMED_IAM(v OneOfstringboolean) {
	o.CAPABILITY_NAMED_IAM = v
}

// GetCAPABILITY_AUTO_EXPAND returns the CAPABILITY_AUTO_EXPAND field value
func (o *SpecTemplateCreateConfigCloudformation) GetCAPABILITY_AUTO_EXPAND() OneOfstringboolean {
	if o == nil  {
		var ret OneOfstringboolean
		return ret
	}

	return o.CAPABILITY_AUTO_EXPAND
}

// GetCAPABILITY_AUTO_EXPANDOk returns a tuple with the CAPABILITY_AUTO_EXPAND field value
// and a boolean to check if the value has been set.
func (o *SpecTemplateCreateConfigCloudformation) GetCAPABILITY_AUTO_EXPANDOk() (*OneOfstringboolean, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CAPABILITY_AUTO_EXPAND, true
}

// SetCAPABILITY_AUTO_EXPAND sets field value
func (o *SpecTemplateCreateConfigCloudformation) SetCAPABILITY_AUTO_EXPAND(v OneOfstringboolean) {
	o.CAPABILITY_AUTO_EXPAND = v
}

func (o SpecTemplateCreateConfigCloudformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["IAM"] = o.IAM
	}
	if true {
		toSerialize["CAPABILITY_NAMED_IAM"] = o.CAPABILITY_NAMED_IAM
	}
	if true {
		toSerialize["CAPABILITY_AUTO_EXPAND"] = o.CAPABILITY_AUTO_EXPAND
	}
	return json.Marshal(toSerialize)
}

type NullableSpecTemplateCreateConfigCloudformation struct {
	value *SpecTemplateCreateConfigCloudformation
	isSet bool
}

func (v NullableSpecTemplateCreateConfigCloudformation) Get() *SpecTemplateCreateConfigCloudformation {
	return v.value
}

func (v *NullableSpecTemplateCreateConfigCloudformation) Set(val *SpecTemplateCreateConfigCloudformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecTemplateCreateConfigCloudformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecTemplateCreateConfigCloudformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecTemplateCreateConfigCloudformation(val *SpecTemplateCreateConfigCloudformation) *NullableSpecTemplateCreateConfigCloudformation {
	return &NullableSpecTemplateCreateConfigCloudformation{value: val, isSet: true}
}

func (v NullableSpecTemplateCreateConfigCloudformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecTemplateCreateConfigCloudformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


