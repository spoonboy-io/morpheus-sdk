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

// Client struct for Client
type Client struct {
	Id *int64 `json:"id,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	AccessTokenValiditySeconds *int64 `json:"accessTokenValiditySeconds,omitempty"`
	RefreshTokenValiditySeconds *int64 `json:"refreshTokenValiditySeconds,omitempty"`
	Authorities *[]string `json:"authorities,omitempty"`
	AuthorizedGrantTypes *[]string `json:"authorizedGrantTypes,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
}

// NewClient instantiates a new Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient() *Client {
	this := Client{}
	return &this
}

// NewClientWithDefaults instantiates a new Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWithDefaults() *Client {
	this := Client{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Client) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Client) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Client) SetId(v int64) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Client) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Client) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Client) SetClientId(v string) {
	o.ClientId = &v
}

// GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field value if set, zero value otherwise.
func (o *Client) GetAccessTokenValiditySeconds() int64 {
	if o == nil || o.AccessTokenValiditySeconds == nil {
		var ret int64
		return ret
	}
	return *o.AccessTokenValiditySeconds
}

// GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAccessTokenValiditySecondsOk() (*int64, bool) {
	if o == nil || o.AccessTokenValiditySeconds == nil {
		return nil, false
	}
	return o.AccessTokenValiditySeconds, true
}

// HasAccessTokenValiditySeconds returns a boolean if a field has been set.
func (o *Client) HasAccessTokenValiditySeconds() bool {
	if o != nil && o.AccessTokenValiditySeconds != nil {
		return true
	}

	return false
}

// SetAccessTokenValiditySeconds gets a reference to the given int64 and assigns it to the AccessTokenValiditySeconds field.
func (o *Client) SetAccessTokenValiditySeconds(v int64) {
	o.AccessTokenValiditySeconds = &v
}

// GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field value if set, zero value otherwise.
func (o *Client) GetRefreshTokenValiditySeconds() int64 {
	if o == nil || o.RefreshTokenValiditySeconds == nil {
		var ret int64
		return ret
	}
	return *o.RefreshTokenValiditySeconds
}

// GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRefreshTokenValiditySecondsOk() (*int64, bool) {
	if o == nil || o.RefreshTokenValiditySeconds == nil {
		return nil, false
	}
	return o.RefreshTokenValiditySeconds, true
}

// HasRefreshTokenValiditySeconds returns a boolean if a field has been set.
func (o *Client) HasRefreshTokenValiditySeconds() bool {
	if o != nil && o.RefreshTokenValiditySeconds != nil {
		return true
	}

	return false
}

// SetRefreshTokenValiditySeconds gets a reference to the given int64 and assigns it to the RefreshTokenValiditySeconds field.
func (o *Client) SetRefreshTokenValiditySeconds(v int64) {
	o.RefreshTokenValiditySeconds = &v
}

// GetAuthorities returns the Authorities field value if set, zero value otherwise.
func (o *Client) GetAuthorities() []string {
	if o == nil || o.Authorities == nil {
		var ret []string
		return ret
	}
	return *o.Authorities
}

// GetAuthoritiesOk returns a tuple with the Authorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthoritiesOk() (*[]string, bool) {
	if o == nil || o.Authorities == nil {
		return nil, false
	}
	return o.Authorities, true
}

// HasAuthorities returns a boolean if a field has been set.
func (o *Client) HasAuthorities() bool {
	if o != nil && o.Authorities != nil {
		return true
	}

	return false
}

// SetAuthorities gets a reference to the given []string and assigns it to the Authorities field.
func (o *Client) SetAuthorities(v []string) {
	o.Authorities = &v
}

// GetAuthorizedGrantTypes returns the AuthorizedGrantTypes field value if set, zero value otherwise.
func (o *Client) GetAuthorizedGrantTypes() []string {
	if o == nil || o.AuthorizedGrantTypes == nil {
		var ret []string
		return ret
	}
	return *o.AuthorizedGrantTypes
}

// GetAuthorizedGrantTypesOk returns a tuple with the AuthorizedGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthorizedGrantTypesOk() (*[]string, bool) {
	if o == nil || o.AuthorizedGrantTypes == nil {
		return nil, false
	}
	return o.AuthorizedGrantTypes, true
}

// HasAuthorizedGrantTypes returns a boolean if a field has been set.
func (o *Client) HasAuthorizedGrantTypes() bool {
	if o != nil && o.AuthorizedGrantTypes != nil {
		return true
	}

	return false
}

// SetAuthorizedGrantTypes gets a reference to the given []string and assigns it to the AuthorizedGrantTypes field.
func (o *Client) SetAuthorizedGrantTypes(v []string) {
	o.AuthorizedGrantTypes = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Client) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Client) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *Client) SetScopes(v []string) {
	o.Scopes = &v
}

func (o Client) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.AccessTokenValiditySeconds != nil {
		toSerialize["accessTokenValiditySeconds"] = o.AccessTokenValiditySeconds
	}
	if o.RefreshTokenValiditySeconds != nil {
		toSerialize["refreshTokenValiditySeconds"] = o.RefreshTokenValiditySeconds
	}
	if o.Authorities != nil {
		toSerialize["authorities"] = o.Authorities
	}
	if o.AuthorizedGrantTypes != nil {
		toSerialize["authorizedGrantTypes"] = o.AuthorizedGrantTypes
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}

type NullableClient struct {
	value *Client
	isSet bool
}

func (v NullableClient) Get() *Client {
	return v.value
}

func (v *NullableClient) Set(val *Client) {
	v.value = val
	v.isSet = true
}

func (v NullableClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient(val *Client) *NullableClient {
	return &NullableClient{value: val, isSet: true}
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


