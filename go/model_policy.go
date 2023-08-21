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

// Policy struct for Policy
type Policy struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	PolicyType *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"policyType,omitempty"`
	Zone NullableInlineResponse20082LoadBalancerInstanceSslCert `json:"zone,omitempty"`
	Site NullableInlineResponse20082LoadBalancerInstanceSslCert `json:"site,omitempty"`
	User NullableInlineResponse20083LoadBalancerNodeCreatedBy `json:"user,omitempty"`
	Role NullablePolicyRole `json:"role,omitempty"`
	RefType NullableString `json:"refType,omitempty"`
	RefId NullableString `json:"refId,omitempty"`
	EachUser NullableBool `json:"eachUser,omitempty"`
	Config *OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject `json:"config,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Owner *InlineResponse20040AppDeployInstance `json:"owner,omitempty"`
	Accounts []InlineResponse20040AppDeployInstance `json:"accounts,omitempty"`
}

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy() *Policy {
	this := Policy{}
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Policy) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Policy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Policy) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Policy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Policy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Policy) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Policy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Policy) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Policy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Policy) UnsetDescription() {
	o.Description.Unset()
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *Policy) GetPolicyType() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.PolicyType == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPolicyTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.PolicyType == nil {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *Policy) HasPolicyType() bool {
	if o != nil && o.PolicyType != nil {
		return true
	}

	return false
}

// SetPolicyType gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the PolicyType field.
func (o *Policy) SetPolicyType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.PolicyType = &v
}

// GetZone returns the Zone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetZone() InlineResponse20082LoadBalancerInstanceSslCert {
	if o == nil || o.Zone.Get() == nil {
		var ret InlineResponse20082LoadBalancerInstanceSslCert
		return ret
	}
	return *o.Zone.Get()
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetZoneOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Zone.Get(), o.Zone.IsSet()
}

// HasZone returns a boolean if a field has been set.
func (o *Policy) HasZone() bool {
	if o != nil && o.Zone.IsSet() {
		return true
	}

	return false
}

// SetZone gets a reference to the given NullableInlineResponse20082LoadBalancerInstanceSslCert and assigns it to the Zone field.
func (o *Policy) SetZone(v InlineResponse20082LoadBalancerInstanceSslCert) {
	o.Zone.Set(&v)
}
// SetZoneNil sets the value for Zone to be an explicit nil
func (o *Policy) SetZoneNil() {
	o.Zone.Set(nil)
}

// UnsetZone ensures that no value is present for Zone, not even an explicit nil
func (o *Policy) UnsetZone() {
	o.Zone.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetSite() InlineResponse20082LoadBalancerInstanceSslCert {
	if o == nil || o.Site.Get() == nil {
		var ret InlineResponse20082LoadBalancerInstanceSslCert
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetSiteOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *Policy) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableInlineResponse20082LoadBalancerInstanceSslCert and assigns it to the Site field.
func (o *Policy) SetSite(v InlineResponse20082LoadBalancerInstanceSslCert) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *Policy) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *Policy) UnsetSite() {
	o.Site.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetUser() InlineResponse20083LoadBalancerNodeCreatedBy {
	if o == nil || o.User.Get() == nil {
		var ret InlineResponse20083LoadBalancerNodeCreatedBy
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetUserOk() (*InlineResponse20083LoadBalancerNodeCreatedBy, bool) {
	if o == nil  {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *Policy) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInlineResponse20083LoadBalancerNodeCreatedBy and assigns it to the User field.
func (o *Policy) SetUser(v InlineResponse20083LoadBalancerNodeCreatedBy) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *Policy) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *Policy) UnsetUser() {
	o.User.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetRole() PolicyRole {
	if o == nil || o.Role.Get() == nil {
		var ret PolicyRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetRoleOk() (*PolicyRole, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *Policy) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullablePolicyRole and assigns it to the Role field.
func (o *Policy) SetRole(v PolicyRole) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *Policy) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *Policy) UnsetRole() {
	o.Role.Unset()
}

// GetRefType returns the RefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetRefType() string {
	if o == nil || o.RefType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefType.Get()
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetRefTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefType.Get(), o.RefType.IsSet()
}

// HasRefType returns a boolean if a field has been set.
func (o *Policy) HasRefType() bool {
	if o != nil && o.RefType.IsSet() {
		return true
	}

	return false
}

// SetRefType gets a reference to the given NullableString and assigns it to the RefType field.
func (o *Policy) SetRefType(v string) {
	o.RefType.Set(&v)
}
// SetRefTypeNil sets the value for RefType to be an explicit nil
func (o *Policy) SetRefTypeNil() {
	o.RefType.Set(nil)
}

// UnsetRefType ensures that no value is present for RefType, not even an explicit nil
func (o *Policy) UnsetRefType() {
	o.RefType.Unset()
}

// GetRefId returns the RefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetRefId() string {
	if o == nil || o.RefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefId.Get()
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefId.Get(), o.RefId.IsSet()
}

// HasRefId returns a boolean if a field has been set.
func (o *Policy) HasRefId() bool {
	if o != nil && o.RefId.IsSet() {
		return true
	}

	return false
}

// SetRefId gets a reference to the given NullableString and assigns it to the RefId field.
func (o *Policy) SetRefId(v string) {
	o.RefId.Set(&v)
}
// SetRefIdNil sets the value for RefId to be an explicit nil
func (o *Policy) SetRefIdNil() {
	o.RefId.Set(nil)
}

// UnsetRefId ensures that no value is present for RefId, not even an explicit nil
func (o *Policy) UnsetRefId() {
	o.RefId.Unset()
}

// GetEachUser returns the EachUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetEachUser() bool {
	if o == nil || o.EachUser.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EachUser.Get()
}

// GetEachUserOk returns a tuple with the EachUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetEachUserOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EachUser.Get(), o.EachUser.IsSet()
}

// HasEachUser returns a boolean if a field has been set.
func (o *Policy) HasEachUser() bool {
	if o != nil && o.EachUser.IsSet() {
		return true
	}

	return false
}

// SetEachUser gets a reference to the given NullableBool and assigns it to the EachUser field.
func (o *Policy) SetEachUser(v bool) {
	o.EachUser.Set(&v)
}
// SetEachUserNil sets the value for EachUser to be an explicit nil
func (o *Policy) SetEachUserNil() {
	o.EachUser.Set(nil)
}

// UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
func (o *Policy) UnsetEachUser() {
	o.EachUser.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Policy) GetConfig() OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject {
	if o == nil || o.Config == nil {
		var ret OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetConfigOk() (*OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Policy) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject and assigns it to the Config field.
func (o *Policy) SetConfig(v OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Policy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Policy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Policy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Policy) GetOwner() InlineResponse20040AppDeployInstance {
	if o == nil || o.Owner == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetOwnerOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Policy) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Owner field.
func (o *Policy) SetOwner(v InlineResponse20040AppDeployInstance) {
	o.Owner = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetAccounts() []InlineResponse20040AppDeployInstance {
	if o == nil  {
		var ret []InlineResponse20040AppDeployInstance
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetAccountsOk() (*[]InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Policy) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []InlineResponse20040AppDeployInstance and assigns it to the Accounts field.
func (o *Policy) SetAccounts(v []InlineResponse20040AppDeployInstance) {
	o.Accounts = v
}

func (o Policy) MarshalJSON() ([]byte, error) {
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
	if o.PolicyType != nil {
		toSerialize["policyType"] = o.PolicyType
	}
	if o.Zone.IsSet() {
		toSerialize["zone"] = o.Zone.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.RefType.IsSet() {
		toSerialize["refType"] = o.RefType.Get()
	}
	if o.RefId.IsSet() {
		toSerialize["refId"] = o.RefId.Get()
	}
	if o.EachUser.IsSet() {
		toSerialize["eachUser"] = o.EachUser.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullablePolicy struct {
	value *Policy
	isSet bool
}

func (v NullablePolicy) Get() *Policy {
	return v.value
}

func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

