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

// ClusterServerCreate struct for ClusterServerCreate
type ClusterServerCreate struct {
	// Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available. 
	Config map[string]interface{} `json:"config"`
	ServerType *ClusterServerCreateServerType `json:"serverType,omitempty"`
	// Name to be used for host(s) created in the cluster
	Name string `json:"name"`
	Plan ClusterServerCreatePlan `json:"plan"`
	// The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects
	Volumes *[]ClusterServerCreateVolumes `json:"volumes,omitempty"`
	// The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes 
	NetworkInterfaces *[]ClusterServerCreateNetworkInterfaces `json:"networkInterfaces,omitempty"`
	// Key for security group configuration.
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
	// Visibility for server host
	Visibility *string `json:"visibility,omitempty"`
	UserGroup *ClusterServerCreateUserGroup `json:"userGroup,omitempty"`
	// Network domain
	NetworkDomain NullableString `json:"networkDomain,omitempty"`
	// Hostname for server host
	Hostname NullableString `json:"hostname,omitempty"`
	// Number of workers or hosts
	NodeCount *int64 `json:"nodeCount,omitempty"`
	// Metadata tags, Array of objects having a name and value.
	Tags *[]ApiServersIdMakeManagedServerTags `json:"tags,omitempty"`
	// Array of strings (keywords). This will set labels on the server and also on the cluster as well by default.
	Labels *[]string `json:"labels,omitempty"`
}

// NewClusterServerCreate instantiates a new ClusterServerCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterServerCreate(config map[string]interface{}, name string, plan ClusterServerCreatePlan, ) *ClusterServerCreate {
	this := ClusterServerCreate{}
	this.Config = config
	this.Name = name
	this.Plan = plan
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewClusterServerCreateWithDefaults instantiates a new ClusterServerCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterServerCreateWithDefaults() *ClusterServerCreate {
	this := ClusterServerCreate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetConfig returns the Config field value
func (o *ClusterServerCreate) GetConfig() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ClusterServerCreate) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetServerType() ClusterServerCreateServerType {
	if o == nil || o.ServerType == nil {
		var ret ClusterServerCreateServerType
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetServerTypeOk() (*ClusterServerCreateServerType, bool) {
	if o == nil || o.ServerType == nil {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasServerType() bool {
	if o != nil && o.ServerType != nil {
		return true
	}

	return false
}

// SetServerType gets a reference to the given ClusterServerCreateServerType and assigns it to the ServerType field.
func (o *ClusterServerCreate) SetServerType(v ClusterServerCreateServerType) {
	o.ServerType = &v
}

// GetName returns the Name field value
func (o *ClusterServerCreate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterServerCreate) SetName(v string) {
	o.Name = v
}

// GetPlan returns the Plan field value
func (o *ClusterServerCreate) GetPlan() ClusterServerCreatePlan {
	if o == nil  {
		var ret ClusterServerCreatePlan
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetPlanOk() (*ClusterServerCreatePlan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *ClusterServerCreate) SetPlan(v ClusterServerCreatePlan) {
	o.Plan = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetVolumes() []ClusterServerCreateVolumes {
	if o == nil || o.Volumes == nil {
		var ret []ClusterServerCreateVolumes
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetVolumesOk() (*[]ClusterServerCreateVolumes, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []ClusterServerCreateVolumes and assigns it to the Volumes field.
func (o *ClusterServerCreate) SetVolumes(v []ClusterServerCreateVolumes) {
	o.Volumes = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetNetworkInterfaces() []ClusterServerCreateNetworkInterfaces {
	if o == nil || o.NetworkInterfaces == nil {
		var ret []ClusterServerCreateNetworkInterfaces
		return ret
	}
	return *o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetNetworkInterfacesOk() (*[]ClusterServerCreateNetworkInterfaces, bool) {
	if o == nil || o.NetworkInterfaces == nil {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasNetworkInterfaces() bool {
	if o != nil && o.NetworkInterfaces != nil {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []ClusterServerCreateNetworkInterfaces and assigns it to the NetworkInterfaces field.
func (o *ClusterServerCreate) SetNetworkInterfaces(v []ClusterServerCreateNetworkInterfaces) {
	o.NetworkInterfaces = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetSecurityGroups() []string {
	if o == nil || o.SecurityGroups == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *ClusterServerCreate) SetSecurityGroups(v []string) {
	o.SecurityGroups = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ClusterServerCreate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetUserGroup returns the UserGroup field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetUserGroup() ClusterServerCreateUserGroup {
	if o == nil || o.UserGroup == nil {
		var ret ClusterServerCreateUserGroup
		return ret
	}
	return *o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetUserGroupOk() (*ClusterServerCreateUserGroup, bool) {
	if o == nil || o.UserGroup == nil {
		return nil, false
	}
	return o.UserGroup, true
}

// HasUserGroup returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasUserGroup() bool {
	if o != nil && o.UserGroup != nil {
		return true
	}

	return false
}

// SetUserGroup gets a reference to the given ClusterServerCreateUserGroup and assigns it to the UserGroup field.
func (o *ClusterServerCreate) SetUserGroup(v ClusterServerCreateUserGroup) {
	o.UserGroup = &v
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterServerCreate) GetNetworkDomain() string {
	if o == nil || o.NetworkDomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkDomain.Get()
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterServerCreate) GetNetworkDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkDomain.Get(), o.NetworkDomain.IsSet()
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain.IsSet() {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given NullableString and assigns it to the NetworkDomain field.
func (o *ClusterServerCreate) SetNetworkDomain(v string) {
	o.NetworkDomain.Set(&v)
}
// SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil
func (o *ClusterServerCreate) SetNetworkDomainNil() {
	o.NetworkDomain.Set(nil)
}

// UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
func (o *ClusterServerCreate) UnsetNetworkDomain() {
	o.NetworkDomain.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterServerCreate) GetHostname() string {
	if o == nil || o.Hostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterServerCreate) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *ClusterServerCreate) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *ClusterServerCreate) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *ClusterServerCreate) UnsetHostname() {
	o.Hostname.Unset()
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetNodeCount() int64 {
	if o == nil || o.NodeCount == nil {
		var ret int64
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetNodeCountOk() (*int64, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int64 and assigns it to the NodeCount field.
func (o *ClusterServerCreate) SetNodeCount(v int64) {
	o.NodeCount = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetTags() []ApiServersIdMakeManagedServerTags {
	if o == nil || o.Tags == nil {
		var ret []ApiServersIdMakeManagedServerTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetTagsOk() (*[]ApiServersIdMakeManagedServerTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ApiServersIdMakeManagedServerTags and assigns it to the Tags field.
func (o *ClusterServerCreate) SetTags(v []ApiServersIdMakeManagedServerTags) {
	o.Tags = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ClusterServerCreate) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServerCreate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ClusterServerCreate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ClusterServerCreate) SetLabels(v []string) {
	o.Labels = &v
}

func (o ClusterServerCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.ServerType != nil {
		toSerialize["serverType"] = o.ServerType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.NetworkInterfaces != nil {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if o.SecurityGroups != nil {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.UserGroup != nil {
		toSerialize["userGroup"] = o.UserGroup
	}
	if o.NetworkDomain.IsSet() {
		toSerialize["networkDomain"] = o.NetworkDomain.Get()
	}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableClusterServerCreate struct {
	value *ClusterServerCreate
	isSet bool
}

func (v NullableClusterServerCreate) Get() *ClusterServerCreate {
	return v.value
}

func (v *NullableClusterServerCreate) Set(val *ClusterServerCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterServerCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterServerCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterServerCreate(val *ClusterServerCreate) *NullableClusterServerCreate {
	return &NullableClusterServerCreate{value: val, isSet: true}
}

func (v NullableClusterServerCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterServerCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


