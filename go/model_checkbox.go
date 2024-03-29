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
	"fmt"
)

// Checkbox the model 'Checkbox'
type Checkbox string

// List of checkbox
const (
	ON Checkbox = "on"
	OFF Checkbox = "off"
)

func (v *Checkbox) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Checkbox(value)
	for _, existing := range []Checkbox{ "on", "off",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Checkbox", value)
}

// Ptr returns reference to checkbox value
func (v Checkbox) Ptr() *Checkbox {
	return &v
}

type NullableCheckbox struct {
	value *Checkbox
	isSet bool
}

func (v NullableCheckbox) Get() *Checkbox {
	return v.value
}

func (v *NullableCheckbox) Set(val *Checkbox) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckbox) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckbox(val *Checkbox) *NullableCheckbox {
	return &NullableCheckbox{value: val, isSet: true}
}

func (v NullableCheckbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

