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

// OptionTypeList struct for OptionTypeList
type OptionTypeList struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	Type *string `json:"type,omitempty"`
	SourceUrl *string `json:"sourceUrl,omitempty"`
	SourceMethod *string `json:"sourceMethod,omitempty"`
	ApiType NullableString `json:"apiType,omitempty"`
	IgnoreSSLErrors *bool `json:"ignoreSSLErrors,omitempty"`
	RealTime *bool `json:"realTime,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Config *OptionTypeListConfig `json:"config,omitempty"`
	Credential *OptionTypeListCredential `json:"credential,omitempty"`
	ServiceUsername NullableString `json:"serviceUsername,omitempty"`
	ServicePassword NullableString `json:"servicePassword,omitempty"`
	InitialDataset NullableString `json:"initialDataset,omitempty"`
	TranslationScript *string `json:"translationScript,omitempty"`
	RequestScript NullableString `json:"requestScript,omitempty"`
	Account *InlineResponse20040AppDeployInstance `json:"account,omitempty"`
}

// NewOptionTypeList instantiates a new OptionTypeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionTypeList() *OptionTypeList {
	this := OptionTypeList{}
	return &this
}

// NewOptionTypeListWithDefaults instantiates a new OptionTypeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionTypeListWithDefaults() *OptionTypeList {
	this := OptionTypeList{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionTypeList) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionTypeList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OptionTypeList) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OptionTypeList) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OptionTypeList) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OptionTypeList) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OptionTypeList) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionTypeList) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OptionTypeList) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OptionTypeList) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OptionTypeList) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OptionTypeList) UnsetDescription() {
	o.Description.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *OptionTypeList) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *OptionTypeList) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *OptionTypeList) SetLabels(v []string) {
	o.Labels = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OptionTypeList) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OptionTypeList) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OptionTypeList) SetType(v string) {
	o.Type = &v
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise.
func (o *OptionTypeList) GetSourceUrl() string {
	if o == nil || o.SourceUrl == nil {
		var ret string
		return ret
	}
	return *o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetSourceUrlOk() (*string, bool) {
	if o == nil || o.SourceUrl == nil {
		return nil, false
	}
	return o.SourceUrl, true
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *OptionTypeList) HasSourceUrl() bool {
	if o != nil && o.SourceUrl != nil {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given string and assigns it to the SourceUrl field.
func (o *OptionTypeList) SetSourceUrl(v string) {
	o.SourceUrl = &v
}

// GetSourceMethod returns the SourceMethod field value if set, zero value otherwise.
func (o *OptionTypeList) GetSourceMethod() string {
	if o == nil || o.SourceMethod == nil {
		var ret string
		return ret
	}
	return *o.SourceMethod
}

// GetSourceMethodOk returns a tuple with the SourceMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetSourceMethodOk() (*string, bool) {
	if o == nil || o.SourceMethod == nil {
		return nil, false
	}
	return o.SourceMethod, true
}

// HasSourceMethod returns a boolean if a field has been set.
func (o *OptionTypeList) HasSourceMethod() bool {
	if o != nil && o.SourceMethod != nil {
		return true
	}

	return false
}

// SetSourceMethod gets a reference to the given string and assigns it to the SourceMethod field.
func (o *OptionTypeList) SetSourceMethod(v string) {
	o.SourceMethod = &v
}

// GetApiType returns the ApiType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OptionTypeList) GetApiType() string {
	if o == nil || o.ApiType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApiType.Get()
}

// GetApiTypeOk returns a tuple with the ApiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionTypeList) GetApiTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApiType.Get(), o.ApiType.IsSet()
}

// HasApiType returns a boolean if a field has been set.
func (o *OptionTypeList) HasApiType() bool {
	if o != nil && o.ApiType.IsSet() {
		return true
	}

	return false
}

// SetApiType gets a reference to the given NullableString and assigns it to the ApiType field.
func (o *OptionTypeList) SetApiType(v string) {
	o.ApiType.Set(&v)
}
// SetApiTypeNil sets the value for ApiType to be an explicit nil
func (o *OptionTypeList) SetApiTypeNil() {
	o.ApiType.Set(nil)
}

// UnsetApiType ensures that no value is present for ApiType, not even an explicit nil
func (o *OptionTypeList) UnsetApiType() {
	o.ApiType.Unset()
}

// GetIgnoreSSLErrors returns the IgnoreSSLErrors field value if set, zero value otherwise.
func (o *OptionTypeList) GetIgnoreSSLErrors() bool {
	if o == nil || o.IgnoreSSLErrors == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreSSLErrors
}

// GetIgnoreSSLErrorsOk returns a tuple with the IgnoreSSLErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetIgnoreSSLErrorsOk() (*bool, bool) {
	if o == nil || o.IgnoreSSLErrors == nil {
		return nil, false
	}
	return o.IgnoreSSLErrors, true
}

// HasIgnoreSSLErrors returns a boolean if a field has been set.
func (o *OptionTypeList) HasIgnoreSSLErrors() bool {
	if o != nil && o.IgnoreSSLErrors != nil {
		return true
	}

	return false
}

// SetIgnoreSSLErrors gets a reference to the given bool and assigns it to the IgnoreSSLErrors field.
func (o *OptionTypeList) SetIgnoreSSLErrors(v bool) {
	o.IgnoreSSLErrors = &v
}

// GetRealTime returns the RealTime field value if set, zero value otherwise.
func (o *OptionTypeList) GetRealTime() bool {
	if o == nil || o.RealTime == nil {
		var ret bool
		return ret
	}
	return *o.RealTime
}

// GetRealTimeOk returns a tuple with the RealTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetRealTimeOk() (*bool, bool) {
	if o == nil || o.RealTime == nil {
		return nil, false
	}
	return o.RealTime, true
}

// HasRealTime returns a boolean if a field has been set.
func (o *OptionTypeList) HasRealTime() bool {
	if o != nil && o.RealTime != nil {
		return true
	}

	return false
}

// SetRealTime gets a reference to the given bool and assigns it to the RealTime field.
func (o *OptionTypeList) SetRealTime(v bool) {
	o.RealTime = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OptionTypeList) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OptionTypeList) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OptionTypeList) SetVisibility(v string) {
	o.Visibility = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *OptionTypeList) GetConfig() OptionTypeListConfig {
	if o == nil || o.Config == nil {
		var ret OptionTypeListConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetConfigOk() (*OptionTypeListConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *OptionTypeList) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OptionTypeListConfig and assigns it to the Config field.
func (o *OptionTypeList) SetConfig(v OptionTypeListConfig) {
	o.Config = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *OptionTypeList) GetCredential() OptionTypeListCredential {
	if o == nil || o.Credential == nil {
		var ret OptionTypeListCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetCredentialOk() (*OptionTypeListCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *OptionTypeList) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given OptionTypeListCredential and assigns it to the Credential field.
func (o *OptionTypeList) SetCredential(v OptionTypeListCredential) {
	o.Credential = &v
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OptionTypeList) GetServiceUsername() string {
	if o == nil || o.ServiceUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceUsername.Get()
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionTypeList) GetServiceUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUsername.Get(), o.ServiceUsername.IsSet()
}

// HasServiceUsername returns a boolean if a field has been set.
func (o *OptionTypeList) HasServiceUsername() bool {
	if o != nil && o.ServiceUsername.IsSet() {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given NullableString and assigns it to the ServiceUsername field.
func (o *OptionTypeList) SetServiceUsername(v string) {
	o.ServiceUsername.Set(&v)
}
// SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil
func (o *OptionTypeList) SetServiceUsernameNil() {
	o.ServiceUsername.Set(nil)
}

// UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
func (o *OptionTypeList) UnsetServiceUsername() {
	o.ServiceUsername.Unset()
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OptionTypeList) GetServicePassword() string {
	if o == nil || o.ServicePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServicePassword.Get()
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionTypeList) GetServicePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePassword.Get(), o.ServicePassword.IsSet()
}

// HasServicePassword returns a boolean if a field has been set.
func (o *OptionTypeList) HasServicePassword() bool {
	if o != nil && o.ServicePassword.IsSet() {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given NullableString and assigns it to the ServicePassword field.
func (o *OptionTypeList) SetServicePassword(v string) {
	o.ServicePassword.Set(&v)
}
// SetServicePasswordNil sets the value for ServicePassword to be an explicit nil
func (o *OptionTypeList) SetServicePasswordNil() {
	o.ServicePassword.Set(nil)
}

// UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
func (o *OptionTypeList) UnsetServicePassword() {
	o.ServicePassword.Unset()
}

// GetInitialDataset returns the InitialDataset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OptionTypeList) GetInitialDataset() string {
	if o == nil || o.InitialDataset.Get() == nil {
		var ret string
		return ret
	}
	return *o.InitialDataset.Get()
}

// GetInitialDatasetOk returns a tuple with the InitialDataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionTypeList) GetInitialDatasetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InitialDataset.Get(), o.InitialDataset.IsSet()
}

// HasInitialDataset returns a boolean if a field has been set.
func (o *OptionTypeList) HasInitialDataset() bool {
	if o != nil && o.InitialDataset.IsSet() {
		return true
	}

	return false
}

// SetInitialDataset gets a reference to the given NullableString and assigns it to the InitialDataset field.
func (o *OptionTypeList) SetInitialDataset(v string) {
	o.InitialDataset.Set(&v)
}
// SetInitialDatasetNil sets the value for InitialDataset to be an explicit nil
func (o *OptionTypeList) SetInitialDatasetNil() {
	o.InitialDataset.Set(nil)
}

// UnsetInitialDataset ensures that no value is present for InitialDataset, not even an explicit nil
func (o *OptionTypeList) UnsetInitialDataset() {
	o.InitialDataset.Unset()
}

// GetTranslationScript returns the TranslationScript field value if set, zero value otherwise.
func (o *OptionTypeList) GetTranslationScript() string {
	if o == nil || o.TranslationScript == nil {
		var ret string
		return ret
	}
	return *o.TranslationScript
}

// GetTranslationScriptOk returns a tuple with the TranslationScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetTranslationScriptOk() (*string, bool) {
	if o == nil || o.TranslationScript == nil {
		return nil, false
	}
	return o.TranslationScript, true
}

// HasTranslationScript returns a boolean if a field has been set.
func (o *OptionTypeList) HasTranslationScript() bool {
	if o != nil && o.TranslationScript != nil {
		return true
	}

	return false
}

// SetTranslationScript gets a reference to the given string and assigns it to the TranslationScript field.
func (o *OptionTypeList) SetTranslationScript(v string) {
	o.TranslationScript = &v
}

// GetRequestScript returns the RequestScript field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OptionTypeList) GetRequestScript() string {
	if o == nil || o.RequestScript.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestScript.Get()
}

// GetRequestScriptOk returns a tuple with the RequestScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OptionTypeList) GetRequestScriptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestScript.Get(), o.RequestScript.IsSet()
}

// HasRequestScript returns a boolean if a field has been set.
func (o *OptionTypeList) HasRequestScript() bool {
	if o != nil && o.RequestScript.IsSet() {
		return true
	}

	return false
}

// SetRequestScript gets a reference to the given NullableString and assigns it to the RequestScript field.
func (o *OptionTypeList) SetRequestScript(v string) {
	o.RequestScript.Set(&v)
}
// SetRequestScriptNil sets the value for RequestScript to be an explicit nil
func (o *OptionTypeList) SetRequestScriptNil() {
	o.RequestScript.Set(nil)
}

// UnsetRequestScript ensures that no value is present for RequestScript, not even an explicit nil
func (o *OptionTypeList) UnsetRequestScript() {
	o.RequestScript.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *OptionTypeList) GetAccount() InlineResponse20040AppDeployInstance {
	if o == nil || o.Account == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeList) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *OptionTypeList) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Account field.
func (o *OptionTypeList) SetAccount(v InlineResponse20040AppDeployInstance) {
	o.Account = &v
}

func (o OptionTypeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SourceUrl != nil {
		toSerialize["sourceUrl"] = o.SourceUrl
	}
	if o.SourceMethod != nil {
		toSerialize["sourceMethod"] = o.SourceMethod
	}
	if o.ApiType.IsSet() {
		toSerialize["apiType"] = o.ApiType.Get()
	}
	if o.IgnoreSSLErrors != nil {
		toSerialize["ignoreSSLErrors"] = o.IgnoreSSLErrors
	}
	if o.RealTime != nil {
		toSerialize["realTime"] = o.RealTime
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	if o.ServiceUsername.IsSet() {
		toSerialize["serviceUsername"] = o.ServiceUsername.Get()
	}
	if o.ServicePassword.IsSet() {
		toSerialize["servicePassword"] = o.ServicePassword.Get()
	}
	if o.InitialDataset.IsSet() {
		toSerialize["initialDataset"] = o.InitialDataset.Get()
	}
	if o.TranslationScript != nil {
		toSerialize["translationScript"] = o.TranslationScript
	}
	if o.RequestScript.IsSet() {
		toSerialize["requestScript"] = o.RequestScript.Get()
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableOptionTypeList struct {
	value *OptionTypeList
	isSet bool
}

func (v NullableOptionTypeList) Get() *OptionTypeList {
	return v.value
}

func (v *NullableOptionTypeList) Set(val *OptionTypeList) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionTypeList) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionTypeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionTypeList(val *OptionTypeList) *NullableOptionTypeList {
	return &NullableOptionTypeList{value: val, isSet: true}
}

func (v NullableOptionTypeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionTypeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


