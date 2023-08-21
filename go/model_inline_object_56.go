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

// InlineObject56 struct for InlineObject56
type InlineObject56 struct {
	Namespace *ClusterNamespaceCreate `json:"namespace,omitempty"`
}

// NewInlineObject56 instantiates a new InlineObject56 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject56() *InlineObject56 {
	this := InlineObject56{}
	return &this
}

// NewInlineObject56WithDefaults instantiates a new InlineObject56 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject56WithDefaults() *InlineObject56 {
	this := InlineObject56{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *InlineObject56) GetNamespace() ClusterNamespaceCreate {
	if o == nil || o.Namespace == nil {
		var ret ClusterNamespaceCreate
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject56) GetNamespaceOk() (*ClusterNamespaceCreate, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *InlineObject56) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given ClusterNamespaceCreate and assigns it to the Namespace field.
func (o *InlineObject56) SetNamespace(v ClusterNamespaceCreate) {
	o.Namespace = &v
}

func (o InlineObject56) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject56 struct {
	value *InlineObject56
	isSet bool
}

func (v NullableInlineObject56) Get() *InlineObject56 {
	return v.value
}

func (v *NullableInlineObject56) Set(val *InlineObject56) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject56) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject56) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject56(val *InlineObject56) *NullableInlineObject56 {
	return &NullableInlineObject56{value: val, isSet: true}
}

func (v NullableInlineObject56) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject56) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


