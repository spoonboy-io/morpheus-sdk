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
	"fmt"
)

// CheckConfig struct for CheckConfig
type CheckConfig struct {
	ERRORUNKNOWN *ERRORUNKNOWN
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckConfig) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ERRORUNKNOWN
	err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
	if err == nil {
		jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
		if string(jsonERRORUNKNOWN) == "{}" { // empty struct
			dst.ERRORUNKNOWN = nil
		} else {
			return nil // data stored in dst.ERRORUNKNOWN, return on the first match
		}
	} else {
		dst.ERRORUNKNOWN = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckConfig)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckConfig) MarshalJSON() ([]byte, error) {
	if src.ERRORUNKNOWN != nil {
		return json.Marshal(&src.ERRORUNKNOWN)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckConfig struct {
	value *CheckConfig
	isSet bool
}

func (v NullableCheckConfig) Get() *CheckConfig {
	return v.value
}

func (v *NullableCheckConfig) Set(val *CheckConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckConfig(val *CheckConfig) *NullableCheckConfig {
	return &NullableCheckConfig{value: val, isSet: true}
}

func (v NullableCheckConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

