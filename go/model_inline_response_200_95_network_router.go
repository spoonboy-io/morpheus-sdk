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

// InlineResponse20095NetworkRouter struct for InlineResponse20095NetworkRouter
type InlineResponse20095NetworkRouter struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Category *string `json:"category,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	RouterType *string `json:"routerType,omitempty"`
	Status *string `json:"status,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ExternalIp NullableString `json:"externalIp,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	ProviderId *string `json:"providerId,omitempty"`
	Type *InlineResponse20095NetworkRouterType `json:"type,omitempty"`
	NetworkServer *InlineResponse20095NetworkRouterNetworkServer `json:"networkServer,omitempty"`
	Zone *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"zone,omitempty"`
	Instance NullableString `json:"instance,omitempty"`
	ExternalNetwork NullableString `json:"externalNetwork,omitempty"`
	Site NullableString `json:"site,omitempty"`
	Interfaces *[]map[string]interface{} `json:"interfaces,omitempty"`
	Firewall *InlineResponse20095NetworkRouterFirewall `json:"firewall,omitempty"`
	Routes *[]map[string]interface{} `json:"routes,omitempty"`
	Nats *[]map[string]interface{} `json:"nats,omitempty"`
	Permissions *InlineResponse20095NetworkRouterPermissions `json:"permissions,omitempty"`
}

// NewInlineResponse20095NetworkRouter instantiates a new InlineResponse20095NetworkRouter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095NetworkRouter() *InlineResponse20095NetworkRouter {
	this := InlineResponse20095NetworkRouter{}
	return &this
}

// NewInlineResponse20095NetworkRouterWithDefaults instantiates a new InlineResponse20095NetworkRouter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095NetworkRouterWithDefaults() *InlineResponse20095NetworkRouter {
	this := InlineResponse20095NetworkRouter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse20095NetworkRouter) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse20095NetworkRouter) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20095NetworkRouter) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20095NetworkRouter) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20095NetworkRouter) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InlineResponse20095NetworkRouter) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InlineResponse20095NetworkRouter) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InlineResponse20095NetworkRouter) UnsetDescription() {
	o.Description.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse20095NetworkRouter) SetCategory(v string) {
	o.Category = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *InlineResponse20095NetworkRouter) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *InlineResponse20095NetworkRouter) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRouterType returns the RouterType field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetRouterType() string {
	if o == nil || o.RouterType == nil {
		var ret string
		return ret
	}
	return *o.RouterType
}

// GetRouterTypeOk returns a tuple with the RouterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetRouterTypeOk() (*string, bool) {
	if o == nil || o.RouterType == nil {
		return nil, false
	}
	return o.RouterType, true
}

// HasRouterType returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasRouterType() bool {
	if o != nil && o.RouterType != nil {
		return true
	}

	return false
}

// SetRouterType gets a reference to the given string and assigns it to the RouterType field.
func (o *InlineResponse20095NetworkRouter) SetRouterType(v string) {
	o.RouterType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20095NetworkRouter) SetStatus(v string) {
	o.Status = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20095NetworkRouter) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExternalIp returns the ExternalIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20095NetworkRouter) GetExternalIp() string {
	if o == nil || o.ExternalIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIp.Get()
}

// GetExternalIpOk returns a tuple with the ExternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20095NetworkRouter) GetExternalIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIp.Get(), o.ExternalIp.IsSet()
}

// HasExternalIp returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasExternalIp() bool {
	if o != nil && o.ExternalIp.IsSet() {
		return true
	}

	return false
}

// SetExternalIp gets a reference to the given NullableString and assigns it to the ExternalIp field.
func (o *InlineResponse20095NetworkRouter) SetExternalIp(v string) {
	o.ExternalIp.Set(&v)
}
// SetExternalIpNil sets the value for ExternalIp to be an explicit nil
func (o *InlineResponse20095NetworkRouter) SetExternalIpNil() {
	o.ExternalIp.Set(nil)
}

// UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
func (o *InlineResponse20095NetworkRouter) UnsetExternalIp() {
	o.ExternalIp.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *InlineResponse20095NetworkRouter) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *InlineResponse20095NetworkRouter) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetType() InlineResponse20095NetworkRouterType {
	if o == nil || o.Type == nil {
		var ret InlineResponse20095NetworkRouterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetTypeOk() (*InlineResponse20095NetworkRouterType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given InlineResponse20095NetworkRouterType and assigns it to the Type field.
func (o *InlineResponse20095NetworkRouter) SetType(v InlineResponse20095NetworkRouterType) {
	o.Type = &v
}

// GetNetworkServer returns the NetworkServer field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetNetworkServer() InlineResponse20095NetworkRouterNetworkServer {
	if o == nil || o.NetworkServer == nil {
		var ret InlineResponse20095NetworkRouterNetworkServer
		return ret
	}
	return *o.NetworkServer
}

// GetNetworkServerOk returns a tuple with the NetworkServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetNetworkServerOk() (*InlineResponse20095NetworkRouterNetworkServer, bool) {
	if o == nil || o.NetworkServer == nil {
		return nil, false
	}
	return o.NetworkServer, true
}

// HasNetworkServer returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasNetworkServer() bool {
	if o != nil && o.NetworkServer != nil {
		return true
	}

	return false
}

// SetNetworkServer gets a reference to the given InlineResponse20095NetworkRouterNetworkServer and assigns it to the NetworkServer field.
func (o *InlineResponse20095NetworkRouter) SetNetworkServer(v InlineResponse20095NetworkRouterNetworkServer) {
	o.NetworkServer = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetZone() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.Zone == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetZoneOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the Zone field.
func (o *InlineResponse20095NetworkRouter) SetZone(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.Zone = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20095NetworkRouter) GetInstance() string {
	if o == nil || o.Instance.Get() == nil {
		var ret string
		return ret
	}
	return *o.Instance.Get()
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20095NetworkRouter) GetInstanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Instance.Get(), o.Instance.IsSet()
}

// HasInstance returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasInstance() bool {
	if o != nil && o.Instance.IsSet() {
		return true
	}

	return false
}

// SetInstance gets a reference to the given NullableString and assigns it to the Instance field.
func (o *InlineResponse20095NetworkRouter) SetInstance(v string) {
	o.Instance.Set(&v)
}
// SetInstanceNil sets the value for Instance to be an explicit nil
func (o *InlineResponse20095NetworkRouter) SetInstanceNil() {
	o.Instance.Set(nil)
}

// UnsetInstance ensures that no value is present for Instance, not even an explicit nil
func (o *InlineResponse20095NetworkRouter) UnsetInstance() {
	o.Instance.Unset()
}

// GetExternalNetwork returns the ExternalNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20095NetworkRouter) GetExternalNetwork() string {
	if o == nil || o.ExternalNetwork.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalNetwork.Get()
}

// GetExternalNetworkOk returns a tuple with the ExternalNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20095NetworkRouter) GetExternalNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalNetwork.Get(), o.ExternalNetwork.IsSet()
}

// HasExternalNetwork returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasExternalNetwork() bool {
	if o != nil && o.ExternalNetwork.IsSet() {
		return true
	}

	return false
}

// SetExternalNetwork gets a reference to the given NullableString and assigns it to the ExternalNetwork field.
func (o *InlineResponse20095NetworkRouter) SetExternalNetwork(v string) {
	o.ExternalNetwork.Set(&v)
}
// SetExternalNetworkNil sets the value for ExternalNetwork to be an explicit nil
func (o *InlineResponse20095NetworkRouter) SetExternalNetworkNil() {
	o.ExternalNetwork.Set(nil)
}

// UnsetExternalNetwork ensures that no value is present for ExternalNetwork, not even an explicit nil
func (o *InlineResponse20095NetworkRouter) UnsetExternalNetwork() {
	o.ExternalNetwork.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20095NetworkRouter) GetSite() string {
	if o == nil || o.Site.Get() == nil {
		var ret string
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20095NetworkRouter) GetSiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableString and assigns it to the Site field.
func (o *InlineResponse20095NetworkRouter) SetSite(v string) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *InlineResponse20095NetworkRouter) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *InlineResponse20095NetworkRouter) UnsetSite() {
	o.Site.Unset()
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetInterfaces() []map[string]interface{} {
	if o == nil || o.Interfaces == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetInterfacesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []map[string]interface{} and assigns it to the Interfaces field.
func (o *InlineResponse20095NetworkRouter) SetInterfaces(v []map[string]interface{}) {
	o.Interfaces = &v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetFirewall() InlineResponse20095NetworkRouterFirewall {
	if o == nil || o.Firewall == nil {
		var ret InlineResponse20095NetworkRouterFirewall
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetFirewallOk() (*InlineResponse20095NetworkRouterFirewall, bool) {
	if o == nil || o.Firewall == nil {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasFirewall() bool {
	if o != nil && o.Firewall != nil {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given InlineResponse20095NetworkRouterFirewall and assigns it to the Firewall field.
func (o *InlineResponse20095NetworkRouter) SetFirewall(v InlineResponse20095NetworkRouterFirewall) {
	o.Firewall = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetRoutes() []map[string]interface{} {
	if o == nil || o.Routes == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetRoutesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []map[string]interface{} and assigns it to the Routes field.
func (o *InlineResponse20095NetworkRouter) SetRoutes(v []map[string]interface{}) {
	o.Routes = &v
}

// GetNats returns the Nats field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetNats() []map[string]interface{} {
	if o == nil || o.Nats == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Nats
}

// GetNatsOk returns a tuple with the Nats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetNatsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Nats == nil {
		return nil, false
	}
	return o.Nats, true
}

// HasNats returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasNats() bool {
	if o != nil && o.Nats != nil {
		return true
	}

	return false
}

// SetNats gets a reference to the given []map[string]interface{} and assigns it to the Nats field.
func (o *InlineResponse20095NetworkRouter) SetNats(v []map[string]interface{}) {
	o.Nats = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouter) GetPermissions() InlineResponse20095NetworkRouterPermissions {
	if o == nil || o.Permissions == nil {
		var ret InlineResponse20095NetworkRouterPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouter) GetPermissionsOk() (*InlineResponse20095NetworkRouterPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouter) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given InlineResponse20095NetworkRouterPermissions and assigns it to the Permissions field.
func (o *InlineResponse20095NetworkRouter) SetPermissions(v InlineResponse20095NetworkRouterPermissions) {
	o.Permissions = &v
}

func (o InlineResponse20095NetworkRouter) MarshalJSON() ([]byte, error) {
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.RouterType != nil {
		toSerialize["routerType"] = o.RouterType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExternalIp.IsSet() {
		toSerialize["externalIp"] = o.ExternalIp.Get()
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.ProviderId != nil {
		toSerialize["providerId"] = o.ProviderId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.NetworkServer != nil {
		toSerialize["networkServer"] = o.NetworkServer
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Instance.IsSet() {
		toSerialize["instance"] = o.Instance.Get()
	}
	if o.ExternalNetwork.IsSet() {
		toSerialize["externalNetwork"] = o.ExternalNetwork.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Interfaces != nil {
		toSerialize["interfaces"] = o.Interfaces
	}
	if o.Firewall != nil {
		toSerialize["firewall"] = o.Firewall
	}
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	if o.Nats != nil {
		toSerialize["nats"] = o.Nats
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095NetworkRouter struct {
	value *InlineResponse20095NetworkRouter
	isSet bool
}

func (v NullableInlineResponse20095NetworkRouter) Get() *InlineResponse20095NetworkRouter {
	return v.value
}

func (v *NullableInlineResponse20095NetworkRouter) Set(val *InlineResponse20095NetworkRouter) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095NetworkRouter) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095NetworkRouter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095NetworkRouter(val *InlineResponse20095NetworkRouter) *NullableInlineResponse20095NetworkRouter {
	return &NullableInlineResponse20095NetworkRouter{value: val, isSet: true}
}

func (v NullableInlineResponse20095NetworkRouter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095NetworkRouter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


