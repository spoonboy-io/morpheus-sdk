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

// UserSourceCreateSaml SAML Configuration
type UserSourceCreateSaml struct {
	// Login Redirect URL
	Url *string `json:"url,omitempty"`
	// Exclude SAML Request Parameter
	DoNotIncludeSAMLRequest *bool `json:"doNotIncludeSAMLRequest,omitempty"`
	// Lougout Post URL
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// SAML Request signing mode
	SAMLSignatureMode *string `json:"SAMLSignatureMode,omitempty"`
	// Only applies when `SAMLSignatureMode=CustomSignature`
	Request509Certificate *string `json:"request509Certificate,omitempty"`
	// RSA Private Key. Only applies when `SAMLSignatureMode=CustomSignature`
	RequestPrivateKey *string `json:"requestPrivateKey,omitempty"`
	// SAML Response Signing Flag
	DoNotValidateSignature *bool `json:"doNotValidateSignature,omitempty"`
	// Signing Public Key. Only applies when `doNotValidateSignature=true`
	PublicKey *string `json:"publicKey,omitempty"`
	// Encryption RSA Private Key
	PrivateKey *string `json:"privateKey,omitempty"`
	// Given Name Attribute Name
	GivenNameAttribute *string `json:"givenNameAttribute,omitempty"`
	// Surname Attribute Name
	SurnameAttribute *string `json:"surnameAttribute,omitempty"`
	// Role Attribute Name
	RoleAttributeName *string `json:"roleAttributeName,omitempty"`
	// Role Attibute Required Value
	RequiredAttributeValue *string `json:"requiredAttributeValue,omitempty"`
}

// NewUserSourceCreateSaml instantiates a new UserSourceCreateSaml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceCreateSaml() *UserSourceCreateSaml {
	this := UserSourceCreateSaml{}
	var doNotIncludeSAMLRequest bool = false
	this.DoNotIncludeSAMLRequest = &doNotIncludeSAMLRequest
	var sAMLSignatureMode string = "NoSignature"
	this.SAMLSignatureMode = &sAMLSignatureMode
	var doNotValidateSignature bool = true
	this.DoNotValidateSignature = &doNotValidateSignature
	return &this
}

// NewUserSourceCreateSamlWithDefaults instantiates a new UserSourceCreateSaml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceCreateSamlWithDefaults() *UserSourceCreateSaml {
	this := UserSourceCreateSaml{}
	var doNotIncludeSAMLRequest bool = false
	this.DoNotIncludeSAMLRequest = &doNotIncludeSAMLRequest
	var sAMLSignatureMode string = "NoSignature"
	this.SAMLSignatureMode = &sAMLSignatureMode
	var doNotValidateSignature bool = true
	this.DoNotValidateSignature = &doNotValidateSignature
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserSourceCreateSaml) SetUrl(v string) {
	o.Url = &v
}

// GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetDoNotIncludeSAMLRequest() bool {
	if o == nil || o.DoNotIncludeSAMLRequest == nil {
		var ret bool
		return ret
	}
	return *o.DoNotIncludeSAMLRequest
}

// GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetDoNotIncludeSAMLRequestOk() (*bool, bool) {
	if o == nil || o.DoNotIncludeSAMLRequest == nil {
		return nil, false
	}
	return o.DoNotIncludeSAMLRequest, true
}

// HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasDoNotIncludeSAMLRequest() bool {
	if o != nil && o.DoNotIncludeSAMLRequest != nil {
		return true
	}

	return false
}

// SetDoNotIncludeSAMLRequest gets a reference to the given bool and assigns it to the DoNotIncludeSAMLRequest field.
func (o *UserSourceCreateSaml) SetDoNotIncludeSAMLRequest(v bool) {
	o.DoNotIncludeSAMLRequest = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *UserSourceCreateSaml) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetSAMLSignatureMode returns the SAMLSignatureMode field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetSAMLSignatureMode() string {
	if o == nil || o.SAMLSignatureMode == nil {
		var ret string
		return ret
	}
	return *o.SAMLSignatureMode
}

// GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetSAMLSignatureModeOk() (*string, bool) {
	if o == nil || o.SAMLSignatureMode == nil {
		return nil, false
	}
	return o.SAMLSignatureMode, true
}

// HasSAMLSignatureMode returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasSAMLSignatureMode() bool {
	if o != nil && o.SAMLSignatureMode != nil {
		return true
	}

	return false
}

// SetSAMLSignatureMode gets a reference to the given string and assigns it to the SAMLSignatureMode field.
func (o *UserSourceCreateSaml) SetSAMLSignatureMode(v string) {
	o.SAMLSignatureMode = &v
}

// GetRequest509Certificate returns the Request509Certificate field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetRequest509Certificate() string {
	if o == nil || o.Request509Certificate == nil {
		var ret string
		return ret
	}
	return *o.Request509Certificate
}

// GetRequest509CertificateOk returns a tuple with the Request509Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetRequest509CertificateOk() (*string, bool) {
	if o == nil || o.Request509Certificate == nil {
		return nil, false
	}
	return o.Request509Certificate, true
}

// HasRequest509Certificate returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasRequest509Certificate() bool {
	if o != nil && o.Request509Certificate != nil {
		return true
	}

	return false
}

// SetRequest509Certificate gets a reference to the given string and assigns it to the Request509Certificate field.
func (o *UserSourceCreateSaml) SetRequest509Certificate(v string) {
	o.Request509Certificate = &v
}

// GetRequestPrivateKey returns the RequestPrivateKey field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetRequestPrivateKey() string {
	if o == nil || o.RequestPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.RequestPrivateKey
}

// GetRequestPrivateKeyOk returns a tuple with the RequestPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetRequestPrivateKeyOk() (*string, bool) {
	if o == nil || o.RequestPrivateKey == nil {
		return nil, false
	}
	return o.RequestPrivateKey, true
}

// HasRequestPrivateKey returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasRequestPrivateKey() bool {
	if o != nil && o.RequestPrivateKey != nil {
		return true
	}

	return false
}

// SetRequestPrivateKey gets a reference to the given string and assigns it to the RequestPrivateKey field.
func (o *UserSourceCreateSaml) SetRequestPrivateKey(v string) {
	o.RequestPrivateKey = &v
}

// GetDoNotValidateSignature returns the DoNotValidateSignature field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetDoNotValidateSignature() bool {
	if o == nil || o.DoNotValidateSignature == nil {
		var ret bool
		return ret
	}
	return *o.DoNotValidateSignature
}

// GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetDoNotValidateSignatureOk() (*bool, bool) {
	if o == nil || o.DoNotValidateSignature == nil {
		return nil, false
	}
	return o.DoNotValidateSignature, true
}

// HasDoNotValidateSignature returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasDoNotValidateSignature() bool {
	if o != nil && o.DoNotValidateSignature != nil {
		return true
	}

	return false
}

// SetDoNotValidateSignature gets a reference to the given bool and assigns it to the DoNotValidateSignature field.
func (o *UserSourceCreateSaml) SetDoNotValidateSignature(v bool) {
	o.DoNotValidateSignature = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *UserSourceCreateSaml) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *UserSourceCreateSaml) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetGivenNameAttribute returns the GivenNameAttribute field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetGivenNameAttribute() string {
	if o == nil || o.GivenNameAttribute == nil {
		var ret string
		return ret
	}
	return *o.GivenNameAttribute
}

// GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetGivenNameAttributeOk() (*string, bool) {
	if o == nil || o.GivenNameAttribute == nil {
		return nil, false
	}
	return o.GivenNameAttribute, true
}

// HasGivenNameAttribute returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasGivenNameAttribute() bool {
	if o != nil && o.GivenNameAttribute != nil {
		return true
	}

	return false
}

// SetGivenNameAttribute gets a reference to the given string and assigns it to the GivenNameAttribute field.
func (o *UserSourceCreateSaml) SetGivenNameAttribute(v string) {
	o.GivenNameAttribute = &v
}

// GetSurnameAttribute returns the SurnameAttribute field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetSurnameAttribute() string {
	if o == nil || o.SurnameAttribute == nil {
		var ret string
		return ret
	}
	return *o.SurnameAttribute
}

// GetSurnameAttributeOk returns a tuple with the SurnameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetSurnameAttributeOk() (*string, bool) {
	if o == nil || o.SurnameAttribute == nil {
		return nil, false
	}
	return o.SurnameAttribute, true
}

// HasSurnameAttribute returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasSurnameAttribute() bool {
	if o != nil && o.SurnameAttribute != nil {
		return true
	}

	return false
}

// SetSurnameAttribute gets a reference to the given string and assigns it to the SurnameAttribute field.
func (o *UserSourceCreateSaml) SetSurnameAttribute(v string) {
	o.SurnameAttribute = &v
}

// GetRoleAttributeName returns the RoleAttributeName field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetRoleAttributeName() string {
	if o == nil || o.RoleAttributeName == nil {
		var ret string
		return ret
	}
	return *o.RoleAttributeName
}

// GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetRoleAttributeNameOk() (*string, bool) {
	if o == nil || o.RoleAttributeName == nil {
		return nil, false
	}
	return o.RoleAttributeName, true
}

// HasRoleAttributeName returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasRoleAttributeName() bool {
	if o != nil && o.RoleAttributeName != nil {
		return true
	}

	return false
}

// SetRoleAttributeName gets a reference to the given string and assigns it to the RoleAttributeName field.
func (o *UserSourceCreateSaml) SetRoleAttributeName(v string) {
	o.RoleAttributeName = &v
}

// GetRequiredAttributeValue returns the RequiredAttributeValue field value if set, zero value otherwise.
func (o *UserSourceCreateSaml) GetRequiredAttributeValue() string {
	if o == nil || o.RequiredAttributeValue == nil {
		var ret string
		return ret
	}
	return *o.RequiredAttributeValue
}

// GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSourceCreateSaml) GetRequiredAttributeValueOk() (*string, bool) {
	if o == nil || o.RequiredAttributeValue == nil {
		return nil, false
	}
	return o.RequiredAttributeValue, true
}

// HasRequiredAttributeValue returns a boolean if a field has been set.
func (o *UserSourceCreateSaml) HasRequiredAttributeValue() bool {
	if o != nil && o.RequiredAttributeValue != nil {
		return true
	}

	return false
}

// SetRequiredAttributeValue gets a reference to the given string and assigns it to the RequiredAttributeValue field.
func (o *UserSourceCreateSaml) SetRequiredAttributeValue(v string) {
	o.RequiredAttributeValue = &v
}

func (o UserSourceCreateSaml) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.DoNotIncludeSAMLRequest != nil {
		toSerialize["doNotIncludeSAMLRequest"] = o.DoNotIncludeSAMLRequest
	}
	if o.LogoutUrl != nil {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if o.SAMLSignatureMode != nil {
		toSerialize["SAMLSignatureMode"] = o.SAMLSignatureMode
	}
	if o.Request509Certificate != nil {
		toSerialize["request509Certificate"] = o.Request509Certificate
	}
	if o.RequestPrivateKey != nil {
		toSerialize["requestPrivateKey"] = o.RequestPrivateKey
	}
	if o.DoNotValidateSignature != nil {
		toSerialize["doNotValidateSignature"] = o.DoNotValidateSignature
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.GivenNameAttribute != nil {
		toSerialize["givenNameAttribute"] = o.GivenNameAttribute
	}
	if o.SurnameAttribute != nil {
		toSerialize["surnameAttribute"] = o.SurnameAttribute
	}
	if o.RoleAttributeName != nil {
		toSerialize["roleAttributeName"] = o.RoleAttributeName
	}
	if o.RequiredAttributeValue != nil {
		toSerialize["requiredAttributeValue"] = o.RequiredAttributeValue
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceCreateSaml struct {
	value *UserSourceCreateSaml
	isSet bool
}

func (v NullableUserSourceCreateSaml) Get() *UserSourceCreateSaml {
	return v.value
}

func (v *NullableUserSourceCreateSaml) Set(val *UserSourceCreateSaml) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceCreateSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceCreateSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceCreateSaml(val *UserSourceCreateSaml) *NullableUserSourceCreateSaml {
	return &NullableUserSourceCreateSaml{value: val, isSet: true}
}

func (v NullableUserSourceCreateSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceCreateSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

