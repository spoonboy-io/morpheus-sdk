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

// ApiAccountsAccount Payload for creating a new tenant
type ApiAccountsAccount struct {
	// Name
	Name string `json:"name"`
	// Description
	Description NullableString `json:"description,omitempty"`
	Role *ApiAccountsAccountRole `json:"role,omitempty"`
	// The subdomain. This will be part of the login URL and username for sub tenant users.
	Subdomain NullableString `json:"subdomain,omitempty"`
	Currency *CurrencyCode `json:"currency,omitempty"`
}

// NewApiAccountsAccount instantiates a new ApiAccountsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccountsAccount(name string, ) *ApiAccountsAccount {
	this := ApiAccountsAccount{}
	this.Name = name
	var currency CurrencyCode = "USD"
	this.Currency = &currency
	return &this
}

// NewApiAccountsAccountWithDefaults instantiates a new ApiAccountsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccountsAccountWithDefaults() *ApiAccountsAccount {
	this := ApiAccountsAccount{}
	var currency CurrencyCode = "USD"
	this.Currency = &currency
	return &this
}

// GetName returns the Name field value
func (o *ApiAccountsAccount) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccount) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAccountsAccount) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAccountsAccount) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAccountsAccount) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAccountsAccount) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApiAccountsAccount) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApiAccountsAccount) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApiAccountsAccount) UnsetDescription() {
	o.Description.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ApiAccountsAccount) GetRole() ApiAccountsAccountRole {
	if o == nil || o.Role == nil {
		var ret ApiAccountsAccountRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccount) GetRoleOk() (*ApiAccountsAccountRole, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ApiAccountsAccount) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given ApiAccountsAccountRole and assigns it to the Role field.
func (o *ApiAccountsAccount) SetRole(v ApiAccountsAccountRole) {
	o.Role = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAccountsAccount) GetSubdomain() string {
	if o == nil || o.Subdomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subdomain.Get()
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAccountsAccount) GetSubdomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subdomain.Get(), o.Subdomain.IsSet()
}

// HasSubdomain returns a boolean if a field has been set.
func (o *ApiAccountsAccount) HasSubdomain() bool {
	if o != nil && o.Subdomain.IsSet() {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given NullableString and assigns it to the Subdomain field.
func (o *ApiAccountsAccount) SetSubdomain(v string) {
	o.Subdomain.Set(&v)
}
// SetSubdomainNil sets the value for Subdomain to be an explicit nil
func (o *ApiAccountsAccount) SetSubdomainNil() {
	o.Subdomain.Set(nil)
}

// UnsetSubdomain ensures that no value is present for Subdomain, not even an explicit nil
func (o *ApiAccountsAccount) UnsetSubdomain() {
	o.Subdomain.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ApiAccountsAccount) GetCurrency() CurrencyCode {
	if o == nil || o.Currency == nil {
		var ret CurrencyCode
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccount) GetCurrencyOk() (*CurrencyCode, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ApiAccountsAccount) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyCode and assigns it to the Currency field.
func (o *ApiAccountsAccount) SetCurrency(v CurrencyCode) {
	o.Currency = &v
}

func (o ApiAccountsAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Subdomain.IsSet() {
		toSerialize["subdomain"] = o.Subdomain.Get()
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableApiAccountsAccount struct {
	value *ApiAccountsAccount
	isSet bool
}

func (v NullableApiAccountsAccount) Get() *ApiAccountsAccount {
	return v.value
}

func (v *NullableApiAccountsAccount) Set(val *ApiAccountsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccountsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccountsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccountsAccount(val *ApiAccountsAccount) *NullableApiAccountsAccount {
	return &NullableApiAccountsAccount{value: val, isSet: true}
}

func (v NullableApiAccountsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccountsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


