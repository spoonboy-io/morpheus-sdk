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
	"time"
)

// Credential struct for Credential
type Credential struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"type,omitempty"`
	Integration NullableString `json:"integration,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	AuthKey NullableInlineResponse20082LoadBalancerInstanceSslCert `json:"authKey,omitempty"`
	AuthPath NullableString `json:"authPath,omitempty"`
	ExternalId NullableString `json:"externalId,omitempty"`
	RefType NullableString `json:"refType,omitempty"`
	RefId NullableString `json:"refId,omitempty"`
	Category NullableString `json:"category,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	StatusDate NullableTime `json:"statusDate,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Account *InlineResponse20040AppDeployInstance `json:"account,omitempty"`
	User *CredentialUser `json:"user,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Config *CredentialConfig `json:"config,omitempty"`
}

// NewCredential instantiates a new Credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredential() *Credential {
	this := Credential{}
	return &this
}

// NewCredentialWithDefaults instantiates a new Credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialWithDefaults() *Credential {
	this := Credential{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Credential) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Credential) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Credential) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Credential) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Credential) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Credential) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Credential) GetType() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.Type == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Credential) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the Type field.
func (o *Credential) SetType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.Type = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetIntegration() string {
	if o == nil || o.Integration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Integration.Get()
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetIntegrationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Integration.Get(), o.Integration.IsSet()
}

// HasIntegration returns a boolean if a field has been set.
func (o *Credential) HasIntegration() bool {
	if o != nil && o.Integration.IsSet() {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given NullableString and assigns it to the Integration field.
func (o *Credential) SetIntegration(v string) {
	o.Integration.Set(&v)
}
// SetIntegrationNil sets the value for Integration to be an explicit nil
func (o *Credential) SetIntegrationNil() {
	o.Integration.Set(nil)
}

// UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
func (o *Credential) UnsetIntegration() {
	o.Integration.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Credential) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Credential) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Credential) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Credential) UnsetDescription() {
	o.Description.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *Credential) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *Credential) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *Credential) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *Credential) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *Credential) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *Credential) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *Credential) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *Credential) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetPasswordHash() string {
	if o == nil || o.PasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetPasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *Credential) HasPasswordHash() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *Credential) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *Credential) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *Credential) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetAuthKey() InlineResponse20082LoadBalancerInstanceSslCert {
	if o == nil || o.AuthKey.Get() == nil {
		var ret InlineResponse20082LoadBalancerInstanceSslCert
		return ret
	}
	return *o.AuthKey.Get()
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetAuthKeyOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthKey.Get(), o.AuthKey.IsSet()
}

// HasAuthKey returns a boolean if a field has been set.
func (o *Credential) HasAuthKey() bool {
	if o != nil && o.AuthKey.IsSet() {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given NullableInlineResponse20082LoadBalancerInstanceSslCert and assigns it to the AuthKey field.
func (o *Credential) SetAuthKey(v InlineResponse20082LoadBalancerInstanceSslCert) {
	o.AuthKey.Set(&v)
}
// SetAuthKeyNil sets the value for AuthKey to be an explicit nil
func (o *Credential) SetAuthKeyNil() {
	o.AuthKey.Set(nil)
}

// UnsetAuthKey ensures that no value is present for AuthKey, not even an explicit nil
func (o *Credential) UnsetAuthKey() {
	o.AuthKey.Unset()
}

// GetAuthPath returns the AuthPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetAuthPath() string {
	if o == nil || o.AuthPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthPath.Get()
}

// GetAuthPathOk returns a tuple with the AuthPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetAuthPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthPath.Get(), o.AuthPath.IsSet()
}

// HasAuthPath returns a boolean if a field has been set.
func (o *Credential) HasAuthPath() bool {
	if o != nil && o.AuthPath.IsSet() {
		return true
	}

	return false
}

// SetAuthPath gets a reference to the given NullableString and assigns it to the AuthPath field.
func (o *Credential) SetAuthPath(v string) {
	o.AuthPath.Set(&v)
}
// SetAuthPathNil sets the value for AuthPath to be an explicit nil
func (o *Credential) SetAuthPathNil() {
	o.AuthPath.Set(nil)
}

// UnsetAuthPath ensures that no value is present for AuthPath, not even an explicit nil
func (o *Credential) UnsetAuthPath() {
	o.AuthPath.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *Credential) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *Credential) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *Credential) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *Credential) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetRefType returns the RefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetRefType() string {
	if o == nil || o.RefType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefType.Get()
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetRefTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefType.Get(), o.RefType.IsSet()
}

// HasRefType returns a boolean if a field has been set.
func (o *Credential) HasRefType() bool {
	if o != nil && o.RefType.IsSet() {
		return true
	}

	return false
}

// SetRefType gets a reference to the given NullableString and assigns it to the RefType field.
func (o *Credential) SetRefType(v string) {
	o.RefType.Set(&v)
}
// SetRefTypeNil sets the value for RefType to be an explicit nil
func (o *Credential) SetRefTypeNil() {
	o.RefType.Set(nil)
}

// UnsetRefType ensures that no value is present for RefType, not even an explicit nil
func (o *Credential) UnsetRefType() {
	o.RefType.Unset()
}

// GetRefId returns the RefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetRefId() string {
	if o == nil || o.RefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefId.Get()
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefId.Get(), o.RefId.IsSet()
}

// HasRefId returns a boolean if a field has been set.
func (o *Credential) HasRefId() bool {
	if o != nil && o.RefId.IsSet() {
		return true
	}

	return false
}

// SetRefId gets a reference to the given NullableString and assigns it to the RefId field.
func (o *Credential) SetRefId(v string) {
	o.RefId.Set(&v)
}
// SetRefIdNil sets the value for RefId to be an explicit nil
func (o *Credential) SetRefIdNil() {
	o.RefId.Set(nil)
}

// UnsetRefId ensures that no value is present for RefId, not even an explicit nil
func (o *Credential) UnsetRefId() {
	o.RefId.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *Credential) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *Credential) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *Credential) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *Credential) UnsetCategory() {
	o.Category.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Credential) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Credential) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Credential) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Credential) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Credential) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Credential) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetStatusMessage() string {
	if o == nil || o.StatusMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetStatusMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Credential) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *Credential) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *Credential) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *Credential) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetStatusDate() time.Time {
	if o == nil || o.StatusDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusDate.Get()
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetStatusDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDate.Get(), o.StatusDate.IsSet()
}

// HasStatusDate returns a boolean if a field has been set.
func (o *Credential) HasStatusDate() bool {
	if o != nil && o.StatusDate.IsSet() {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given NullableTime and assigns it to the StatusDate field.
func (o *Credential) SetStatusDate(v time.Time) {
	o.StatusDate.Set(&v)
}
// SetStatusDateNil sets the value for StatusDate to be an explicit nil
func (o *Credential) SetStatusDateNil() {
	o.StatusDate.Set(nil)
}

// UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
func (o *Credential) UnsetStatusDate() {
	o.StatusDate.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Credential) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Credential) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Credential) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Credential) GetAccount() InlineResponse20040AppDeployInstance {
	if o == nil || o.Account == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Credential) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Account field.
func (o *Credential) SetAccount(v InlineResponse20040AppDeployInstance) {
	o.Account = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Credential) GetUser() CredentialUser {
	if o == nil || o.User == nil {
		var ret CredentialUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetUserOk() (*CredentialUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Credential) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given CredentialUser and assigns it to the User field.
func (o *Credential) SetUser(v CredentialUser) {
	o.User = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Credential) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Credential) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Credential) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Credential) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Credential) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Credential) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Credential) GetConfig() CredentialConfig {
	if o == nil || o.Config == nil {
		var ret CredentialConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetConfigOk() (*CredentialConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Credential) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given CredentialConfig and assigns it to the Config field.
func (o *Credential) SetConfig(v CredentialConfig) {
	o.Config = &v
}

func (o Credential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Integration.IsSet() {
		toSerialize["integration"] = o.Integration.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	if o.AuthKey.IsSet() {
		toSerialize["authKey"] = o.AuthKey.Get()
	}
	if o.AuthPath.IsSet() {
		toSerialize["authPath"] = o.AuthPath.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.RefType.IsSet() {
		toSerialize["refType"] = o.RefType.Get()
	}
	if o.RefId.IsSet() {
		toSerialize["refId"] = o.RefId.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage.IsSet() {
		toSerialize["statusMessage"] = o.StatusMessage.Get()
	}
	if o.StatusDate.IsSet() {
		toSerialize["statusDate"] = o.StatusDate.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableCredential struct {
	value *Credential
	isSet bool
}

func (v NullableCredential) Get() *Credential {
	return v.value
}

func (v *NullableCredential) Set(val *Credential) {
	v.value = val
	v.isSet = true
}

func (v NullableCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredential(val *Credential) *NullableCredential {
	return &NullableCredential{value: val, isSet: true}
}

func (v NullableCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


