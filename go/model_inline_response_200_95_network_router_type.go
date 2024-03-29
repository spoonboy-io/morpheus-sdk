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

// InlineResponse20095NetworkRouterType struct for InlineResponse20095NetworkRouterType
type InlineResponse20095NetworkRouterType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Creatable *bool `json:"creatable,omitempty"`
	Selectable *bool `json:"selectable,omitempty"`
	HasFirewall *bool `json:"hasFirewall,omitempty"`
	HasDhcp *bool `json:"hasDhcp,omitempty"`
	HasRouting *bool `json:"hasRouting,omitempty"`
	HasNat *bool `json:"hasNat,omitempty"`
	HasNetworkServer *bool `json:"hasNetworkServer,omitempty"`
	HasFirewallGroups *bool `json:"hasFirewallGroups,omitempty"`
	HasSecurityGroupPriority *bool `json:"hasSecurityGroupPriority,omitempty"`
	OptionTypes *[]OptionType `json:"optionTypes,omitempty"`
	RuleOptionTypes *[]OptionType `json:"ruleOptionTypes,omitempty"`
	FirewallGroupOptionTypes *[]OptionType `json:"firewallGroupOptionTypes,omitempty"`
	NatOptionTypes *[]OptionType `json:"natOptionTypes,omitempty"`
}

// NewInlineResponse20095NetworkRouterType instantiates a new InlineResponse20095NetworkRouterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095NetworkRouterType() *InlineResponse20095NetworkRouterType {
	this := InlineResponse20095NetworkRouterType{}
	return &this
}

// NewInlineResponse20095NetworkRouterTypeWithDefaults instantiates a new InlineResponse20095NetworkRouterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095NetworkRouterTypeWithDefaults() *InlineResponse20095NetworkRouterType {
	this := InlineResponse20095NetworkRouterType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse20095NetworkRouterType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse20095NetworkRouterType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20095NetworkRouterType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20095NetworkRouterType) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20095NetworkRouterType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetCreatable() bool {
	if o == nil || o.Creatable == nil {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetCreatableOk() (*bool, bool) {
	if o == nil || o.Creatable == nil {
		return nil, false
	}
	return o.Creatable, true
}

// HasCreatable returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasCreatable() bool {
	if o != nil && o.Creatable != nil {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *InlineResponse20095NetworkRouterType) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetSelectable returns the Selectable field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetSelectable() bool {
	if o == nil || o.Selectable == nil {
		var ret bool
		return ret
	}
	return *o.Selectable
}

// GetSelectableOk returns a tuple with the Selectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetSelectableOk() (*bool, bool) {
	if o == nil || o.Selectable == nil {
		return nil, false
	}
	return o.Selectable, true
}

// HasSelectable returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasSelectable() bool {
	if o != nil && o.Selectable != nil {
		return true
	}

	return false
}

// SetSelectable gets a reference to the given bool and assigns it to the Selectable field.
func (o *InlineResponse20095NetworkRouterType) SetSelectable(v bool) {
	o.Selectable = &v
}

// GetHasFirewall returns the HasFirewall field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasFirewall() bool {
	if o == nil || o.HasFirewall == nil {
		var ret bool
		return ret
	}
	return *o.HasFirewall
}

// GetHasFirewallOk returns a tuple with the HasFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasFirewallOk() (*bool, bool) {
	if o == nil || o.HasFirewall == nil {
		return nil, false
	}
	return o.HasFirewall, true
}

// HasHasFirewall returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasFirewall() bool {
	if o != nil && o.HasFirewall != nil {
		return true
	}

	return false
}

// SetHasFirewall gets a reference to the given bool and assigns it to the HasFirewall field.
func (o *InlineResponse20095NetworkRouterType) SetHasFirewall(v bool) {
	o.HasFirewall = &v
}

// GetHasDhcp returns the HasDhcp field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasDhcp() bool {
	if o == nil || o.HasDhcp == nil {
		var ret bool
		return ret
	}
	return *o.HasDhcp
}

// GetHasDhcpOk returns a tuple with the HasDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasDhcpOk() (*bool, bool) {
	if o == nil || o.HasDhcp == nil {
		return nil, false
	}
	return o.HasDhcp, true
}

// HasHasDhcp returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasDhcp() bool {
	if o != nil && o.HasDhcp != nil {
		return true
	}

	return false
}

// SetHasDhcp gets a reference to the given bool and assigns it to the HasDhcp field.
func (o *InlineResponse20095NetworkRouterType) SetHasDhcp(v bool) {
	o.HasDhcp = &v
}

// GetHasRouting returns the HasRouting field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasRouting() bool {
	if o == nil || o.HasRouting == nil {
		var ret bool
		return ret
	}
	return *o.HasRouting
}

// GetHasRoutingOk returns a tuple with the HasRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasRoutingOk() (*bool, bool) {
	if o == nil || o.HasRouting == nil {
		return nil, false
	}
	return o.HasRouting, true
}

// HasHasRouting returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasRouting() bool {
	if o != nil && o.HasRouting != nil {
		return true
	}

	return false
}

// SetHasRouting gets a reference to the given bool and assigns it to the HasRouting field.
func (o *InlineResponse20095NetworkRouterType) SetHasRouting(v bool) {
	o.HasRouting = &v
}

// GetHasNat returns the HasNat field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasNat() bool {
	if o == nil || o.HasNat == nil {
		var ret bool
		return ret
	}
	return *o.HasNat
}

// GetHasNatOk returns a tuple with the HasNat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasNatOk() (*bool, bool) {
	if o == nil || o.HasNat == nil {
		return nil, false
	}
	return o.HasNat, true
}

// HasHasNat returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasNat() bool {
	if o != nil && o.HasNat != nil {
		return true
	}

	return false
}

// SetHasNat gets a reference to the given bool and assigns it to the HasNat field.
func (o *InlineResponse20095NetworkRouterType) SetHasNat(v bool) {
	o.HasNat = &v
}

// GetHasNetworkServer returns the HasNetworkServer field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasNetworkServer() bool {
	if o == nil || o.HasNetworkServer == nil {
		var ret bool
		return ret
	}
	return *o.HasNetworkServer
}

// GetHasNetworkServerOk returns a tuple with the HasNetworkServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasNetworkServerOk() (*bool, bool) {
	if o == nil || o.HasNetworkServer == nil {
		return nil, false
	}
	return o.HasNetworkServer, true
}

// HasHasNetworkServer returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasNetworkServer() bool {
	if o != nil && o.HasNetworkServer != nil {
		return true
	}

	return false
}

// SetHasNetworkServer gets a reference to the given bool and assigns it to the HasNetworkServer field.
func (o *InlineResponse20095NetworkRouterType) SetHasNetworkServer(v bool) {
	o.HasNetworkServer = &v
}

// GetHasFirewallGroups returns the HasFirewallGroups field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasFirewallGroups() bool {
	if o == nil || o.HasFirewallGroups == nil {
		var ret bool
		return ret
	}
	return *o.HasFirewallGroups
}

// GetHasFirewallGroupsOk returns a tuple with the HasFirewallGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasFirewallGroupsOk() (*bool, bool) {
	if o == nil || o.HasFirewallGroups == nil {
		return nil, false
	}
	return o.HasFirewallGroups, true
}

// HasHasFirewallGroups returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasFirewallGroups() bool {
	if o != nil && o.HasFirewallGroups != nil {
		return true
	}

	return false
}

// SetHasFirewallGroups gets a reference to the given bool and assigns it to the HasFirewallGroups field.
func (o *InlineResponse20095NetworkRouterType) SetHasFirewallGroups(v bool) {
	o.HasFirewallGroups = &v
}

// GetHasSecurityGroupPriority returns the HasSecurityGroupPriority field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetHasSecurityGroupPriority() bool {
	if o == nil || o.HasSecurityGroupPriority == nil {
		var ret bool
		return ret
	}
	return *o.HasSecurityGroupPriority
}

// GetHasSecurityGroupPriorityOk returns a tuple with the HasSecurityGroupPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetHasSecurityGroupPriorityOk() (*bool, bool) {
	if o == nil || o.HasSecurityGroupPriority == nil {
		return nil, false
	}
	return o.HasSecurityGroupPriority, true
}

// HasHasSecurityGroupPriority returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasHasSecurityGroupPriority() bool {
	if o != nil && o.HasSecurityGroupPriority != nil {
		return true
	}

	return false
}

// SetHasSecurityGroupPriority gets a reference to the given bool and assigns it to the HasSecurityGroupPriority field.
func (o *InlineResponse20095NetworkRouterType) SetHasSecurityGroupPriority(v bool) {
	o.HasSecurityGroupPriority = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetOptionTypes() []OptionType {
	if o == nil || o.OptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []OptionType and assigns it to the OptionTypes field.
func (o *InlineResponse20095NetworkRouterType) SetOptionTypes(v []OptionType) {
	o.OptionTypes = &v
}

// GetRuleOptionTypes returns the RuleOptionTypes field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetRuleOptionTypes() []OptionType {
	if o == nil || o.RuleOptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.RuleOptionTypes
}

// GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetRuleOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.RuleOptionTypes == nil {
		return nil, false
	}
	return o.RuleOptionTypes, true
}

// HasRuleOptionTypes returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasRuleOptionTypes() bool {
	if o != nil && o.RuleOptionTypes != nil {
		return true
	}

	return false
}

// SetRuleOptionTypes gets a reference to the given []OptionType and assigns it to the RuleOptionTypes field.
func (o *InlineResponse20095NetworkRouterType) SetRuleOptionTypes(v []OptionType) {
	o.RuleOptionTypes = &v
}

// GetFirewallGroupOptionTypes returns the FirewallGroupOptionTypes field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetFirewallGroupOptionTypes() []OptionType {
	if o == nil || o.FirewallGroupOptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.FirewallGroupOptionTypes
}

// GetFirewallGroupOptionTypesOk returns a tuple with the FirewallGroupOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetFirewallGroupOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.FirewallGroupOptionTypes == nil {
		return nil, false
	}
	return o.FirewallGroupOptionTypes, true
}

// HasFirewallGroupOptionTypes returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasFirewallGroupOptionTypes() bool {
	if o != nil && o.FirewallGroupOptionTypes != nil {
		return true
	}

	return false
}

// SetFirewallGroupOptionTypes gets a reference to the given []OptionType and assigns it to the FirewallGroupOptionTypes field.
func (o *InlineResponse20095NetworkRouterType) SetFirewallGroupOptionTypes(v []OptionType) {
	o.FirewallGroupOptionTypes = &v
}

// GetNatOptionTypes returns the NatOptionTypes field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterType) GetNatOptionTypes() []OptionType {
	if o == nil || o.NatOptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.NatOptionTypes
}

// GetNatOptionTypesOk returns a tuple with the NatOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterType) GetNatOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.NatOptionTypes == nil {
		return nil, false
	}
	return o.NatOptionTypes, true
}

// HasNatOptionTypes returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterType) HasNatOptionTypes() bool {
	if o != nil && o.NatOptionTypes != nil {
		return true
	}

	return false
}

// SetNatOptionTypes gets a reference to the given []OptionType and assigns it to the NatOptionTypes field.
func (o *InlineResponse20095NetworkRouterType) SetNatOptionTypes(v []OptionType) {
	o.NatOptionTypes = &v
}

func (o InlineResponse20095NetworkRouterType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Creatable != nil {
		toSerialize["creatable"] = o.Creatable
	}
	if o.Selectable != nil {
		toSerialize["selectable"] = o.Selectable
	}
	if o.HasFirewall != nil {
		toSerialize["hasFirewall"] = o.HasFirewall
	}
	if o.HasDhcp != nil {
		toSerialize["hasDhcp"] = o.HasDhcp
	}
	if o.HasRouting != nil {
		toSerialize["hasRouting"] = o.HasRouting
	}
	if o.HasNat != nil {
		toSerialize["hasNat"] = o.HasNat
	}
	if o.HasNetworkServer != nil {
		toSerialize["hasNetworkServer"] = o.HasNetworkServer
	}
	if o.HasFirewallGroups != nil {
		toSerialize["hasFirewallGroups"] = o.HasFirewallGroups
	}
	if o.HasSecurityGroupPriority != nil {
		toSerialize["hasSecurityGroupPriority"] = o.HasSecurityGroupPriority
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if o.RuleOptionTypes != nil {
		toSerialize["ruleOptionTypes"] = o.RuleOptionTypes
	}
	if o.FirewallGroupOptionTypes != nil {
		toSerialize["firewallGroupOptionTypes"] = o.FirewallGroupOptionTypes
	}
	if o.NatOptionTypes != nil {
		toSerialize["natOptionTypes"] = o.NatOptionTypes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095NetworkRouterType struct {
	value *InlineResponse20095NetworkRouterType
	isSet bool
}

func (v NullableInlineResponse20095NetworkRouterType) Get() *InlineResponse20095NetworkRouterType {
	return v.value
}

func (v *NullableInlineResponse20095NetworkRouterType) Set(val *InlineResponse20095NetworkRouterType) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095NetworkRouterType) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095NetworkRouterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095NetworkRouterType(val *InlineResponse20095NetworkRouterType) *NullableInlineResponse20095NetworkRouterType {
	return &NullableInlineResponse20095NetworkRouterType{value: val, isSet: true}
}

func (v NullableInlineResponse20095NetworkRouterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095NetworkRouterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


