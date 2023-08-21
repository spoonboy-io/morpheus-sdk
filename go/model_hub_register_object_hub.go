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

// HubRegisterObjectHub struct for HubRegisterObjectHub
type HubRegisterObjectHub struct {
	// Company Name for new Morpheus Hub organization
	CompanyName string `json:"companyName"`
	// First Name for new Morpheus Hub user
	FirstName string `json:"firstName"`
	// Last Name for new Morpheus Hub user
	LastName string `json:"lastName"`
	// Email for new Morpheus Hub user
	Email string `json:"email"`
	// Password for new Morpheus Hub user
	Password string `json:"password"`
	// Job title of new Morpheus Hub user
	JobTitle string `json:"jobTitle"`
}

// NewHubRegisterObjectHub instantiates a new HubRegisterObjectHub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubRegisterObjectHub(companyName string, firstName string, lastName string, email string, password string, jobTitle string, ) *HubRegisterObjectHub {
	this := HubRegisterObjectHub{}
	this.CompanyName = companyName
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Password = password
	this.JobTitle = jobTitle
	return &this
}

// NewHubRegisterObjectHubWithDefaults instantiates a new HubRegisterObjectHub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubRegisterObjectHubWithDefaults() *HubRegisterObjectHub {
	this := HubRegisterObjectHub{}
	return &this
}

// GetCompanyName returns the CompanyName field value
func (o *HubRegisterObjectHub) GetCompanyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *HubRegisterObjectHub) GetCompanyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *HubRegisterObjectHub) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetFirstName returns the FirstName field value
func (o *HubRegisterObjectHub) GetFirstName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *HubRegisterObjectHub) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *HubRegisterObjectHub) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *HubRegisterObjectHub) GetLastName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *HubRegisterObjectHub) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *HubRegisterObjectHub) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *HubRegisterObjectHub) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *HubRegisterObjectHub) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *HubRegisterObjectHub) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *HubRegisterObjectHub) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *HubRegisterObjectHub) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *HubRegisterObjectHub) SetPassword(v string) {
	o.Password = v
}

// GetJobTitle returns the JobTitle field value
func (o *HubRegisterObjectHub) GetJobTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value
// and a boolean to check if the value has been set.
func (o *HubRegisterObjectHub) GetJobTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JobTitle, true
}

// SetJobTitle sets field value
func (o *HubRegisterObjectHub) SetJobTitle(v string) {
	o.JobTitle = v
}

func (o HubRegisterObjectHub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["companyName"] = o.CompanyName
	}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["jobTitle"] = o.JobTitle
	}
	return json.Marshal(toSerialize)
}

type NullableHubRegisterObjectHub struct {
	value *HubRegisterObjectHub
	isSet bool
}

func (v NullableHubRegisterObjectHub) Get() *HubRegisterObjectHub {
	return v.value
}

func (v *NullableHubRegisterObjectHub) Set(val *HubRegisterObjectHub) {
	v.value = val
	v.isSet = true
}

func (v NullableHubRegisterObjectHub) IsSet() bool {
	return v.isSet
}

func (v *NullableHubRegisterObjectHub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubRegisterObjectHub(val *HubRegisterObjectHub) *NullableHubRegisterObjectHub {
	return &NullableHubRegisterObjectHub{value: val, isSet: true}
}

func (v NullableHubRegisterObjectHub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubRegisterObjectHub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

