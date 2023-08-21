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

// UserSourceCreateActiveDirectory Active Directory Configuration
type UserSourceCreateActiveDirectory struct {
	// AD Server IP/FQDN
	Url *string `json:"url,omitempty"`
	// Domain
	Domain *string `json:"domain,omitempty"`
	// Use SSL
	UseSSL *bool `json:"useSSL,omitempty"`
	// Binding Username
	BindingUsername *string `json:"bindingUsername,omitempty"`
	// Binding Password
	BindingPassword *string `json:"bindingPassword,omitempty"`
	// Required Group
	RequiredGroup *string `json:"requiredGroup,omitempty"`
	// Include Member Groups
	SearchMemberGroups *bool `json:"searchMemberGroups,omitempty"`
}

// NewUserSourceCreateActiveDirectory instantiates a new UserSourceCreateActiveDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceCreateActiveDirectory() *UserSourceCreateActiveDirectory {
	this := UserSourceCreateActiveDirectory{}
	var useSSL bool = false
	this.UseSSL = &useSSL
	var searchMemberGroups bool = false
	this.SearchMemberGroups = &searchMemberGroups
	return &this
}

// NewUserSourceCreateActiveDirectoryWithDefaults instantiates a new UserSourceCreateActiveDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceCreateActiveDirectoryWithDefaults() *UserSourceCreateActiveDirectory {
	this := UserSourceCreateActiveDirectory{}
	var useSSL bool = false
	this.UseSSL = &useSSL
	var searchMemberGroups bool = false
	this.SearchMemberGroups = &searchMemberGroups
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserSourceCreateActiveDirectory) SetUrl(v string) {
	o.Url = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *UserSourceCreateActiveDirectory) SetDomain(v string) {
	o.Domain = &v
}

// GetUseSSL returns the UseSSL field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetUseSSL() bool {
	if o == nil || o.UseSSL == nil {
		var ret bool
		return ret
	}
	return *o.UseSSL
}

// GetUseSSLOk returns a tuple with the UseSSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetUseSSLOk() (*bool, bool) {
	if o == nil || o.UseSSL == nil {
		return nil, false
	}
	return o.UseSSL, true
}

// HasUseSSL returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasUseSSL() bool {
	if o != nil && o.UseSSL != nil {
		return true
	}

	return false
}

// SetUseSSL gets a reference to the given bool and assigns it to the UseSSL field.
func (o *UserSourceCreateActiveDirectory) SetUseSSL(v bool) {
	o.UseSSL = &v
}

// GetBindingUsername returns the BindingUsername field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetBindingUsername() string {
	if o == nil || o.BindingUsername == nil {
		var ret string
		return ret
	}
	return *o.BindingUsername
}

// GetBindingUsernameOk returns a tuple with the BindingUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetBindingUsernameOk() (*string, bool) {
	if o == nil || o.BindingUsername == nil {
		return nil, false
	}
	return o.BindingUsername, true
}

// HasBindingUsername returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasBindingUsername() bool {
	if o != nil && o.BindingUsername != nil {
		return true
	}

	return false
}

// SetBindingUsername gets a reference to the given string and assigns it to the BindingUsername field.
func (o *UserSourceCreateActiveDirectory) SetBindingUsername(v string) {
	o.BindingUsername = &v
}

// GetBindingPassword returns the BindingPassword field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetBindingPassword() string {
	if o == nil || o.BindingPassword == nil {
		var ret string
		return ret
	}
	return *o.BindingPassword
}

// GetBindingPasswordOk returns a tuple with the BindingPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetBindingPasswordOk() (*string, bool) {
	if o == nil || o.BindingPassword == nil {
		return nil, false
	}
	return o.BindingPassword, true
}

// HasBindingPassword returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasBindingPassword() bool {
	if o != nil && o.BindingPassword != nil {
		return true
	}

	return false
}

// SetBindingPassword gets a reference to the given string and assigns it to the BindingPassword field.
func (o *UserSourceCreateActiveDirectory) SetBindingPassword(v string) {
	o.BindingPassword = &v
}

// GetRequiredGroup returns the RequiredGroup field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetRequiredGroup() string {
	if o == nil || o.RequiredGroup == nil {
		var ret string
		return ret
	}
	return *o.RequiredGroup
}

// GetRequiredGroupOk returns a tuple with the RequiredGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetRequiredGroupOk() (*string, bool) {
	if o == nil || o.RequiredGroup == nil {
		return nil, false
	}
	return o.RequiredGroup, true
}

// HasRequiredGroup returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasRequiredGroup() bool {
	if o != nil && o.RequiredGroup != nil {
		return true
	}

	return false
}

// SetRequiredGroup gets a reference to the given string and assigns it to the RequiredGroup field.
func (o *UserSourceCreateActiveDirectory) SetRequiredGroup(v string) {
	o.RequiredGroup = &v
}

// GetSearchMemberGroups returns the SearchMemberGroups field value if set, zero value otherwise.
func (o *UserSourceCreateActiveDirectory) GetSearchMemberGroups() bool {
	if o == nil || o.SearchMemberGroups == nil {
		var ret bool
		return ret
	}
	return *o.SearchMemberGroups
}

// GetSearchMemberGroupsOk returns a tuple with the SearchMemberGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateActiveDirectory) GetSearchMemberGroupsOk() (*bool, bool) {
	if o == nil || o.SearchMemberGroups == nil {
		return nil, false
	}
	return o.SearchMemberGroups, true
}

// HasSearchMemberGroups returns a boolean if a field has been set.
func (o *UserSourceCreateActiveDirectory) HasSearchMemberGroups() bool {
	if o != nil && o.SearchMemberGroups != nil {
		return true
	}

	return false
}

// SetSearchMemberGroups gets a reference to the given bool and assigns it to the SearchMemberGroups field.
func (o *UserSourceCreateActiveDirectory) SetSearchMemberGroups(v bool) {
	o.SearchMemberGroups = &v
}

func (o UserSourceCreateActiveDirectory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.UseSSL != nil {
		toSerialize["useSSL"] = o.UseSSL
	}
	if o.BindingUsername != nil {
		toSerialize["bindingUsername"] = o.BindingUsername
	}
	if o.BindingPassword != nil {
		toSerialize["bindingPassword"] = o.BindingPassword
	}
	if o.RequiredGroup != nil {
		toSerialize["requiredGroup"] = o.RequiredGroup
	}
	if o.SearchMemberGroups != nil {
		toSerialize["searchMemberGroups"] = o.SearchMemberGroups
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceCreateActiveDirectory struct {
	value *UserSourceCreateActiveDirectory
	isSet bool
}

func (v NullableUserSourceCreateActiveDirectory) Get() *UserSourceCreateActiveDirectory {
	return v.value
}

func (v *NullableUserSourceCreateActiveDirectory) Set(val *UserSourceCreateActiveDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceCreateActiveDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceCreateActiveDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceCreateActiveDirectory(val *UserSourceCreateActiveDirectory) *NullableUserSourceCreateActiveDirectory {
	return &NullableUserSourceCreateActiveDirectory{value: val, isSet: true}
}

func (v NullableUserSourceCreateActiveDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceCreateActiveDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


