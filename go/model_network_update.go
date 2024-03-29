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

// NetworkUpdate struct for NetworkUpdate
type NetworkUpdate struct {
	// Display Name
	DisplayName *string `json:"displayName,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Description
	Description NullableString `json:"description,omitempty"`
	// CIDR Network
	Cidr *string `json:"cidr,omitempty"`
	// Network Gateway
	Gateway *string `json:"gateway,omitempty"`
	// Primary DNS Server
	DnsPrimary *string `json:"dnsPrimary,omitempty"`
	// Secondary DNS Server
	DnsSecondary *string `json:"dnsSecondary,omitempty"`
	VlanId *int64 `json:"vlanId,omitempty"`
	// Network Pool ID
	Pool NullableInt64 `json:"pool,omitempty"`
	// Allow IP Override
	AllowStaticOverride *bool `json:"allowStaticOverride,omitempty"`
	// Assign Public IP
	AssignPublicIp *bool `json:"assignPublicIp,omitempty"`
	// Activate (true) or disable (false) the network
	Active *bool `json:"active,omitempty"`
	// DHCP Server enabled network
	DhcpServer *bool `json:"dhcpServer,omitempty"`
	NetworkDomain *NetworkCreateNetworkDomain `json:"networkDomain,omitempty"`
	// Search Domains
	SearchDomains *string `json:"searchDomains,omitempty"`
	NetworkProxy *NetworkCreateNetworkProxy `json:"networkProxy,omitempty"`
	// Bypass Proxy for Appliance URL
	ApplianceUrlProxyBypass *bool `json:"applianceUrlProxyBypass,omitempty"`
	// Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded.
	NoProxy NullableString `json:"noProxy,omitempty"`
	// Visibility, private or public.
	Visibility *string `json:"visibility,omitempty"`
	// Configuration object. Settings vary by type.
	Config *map[string]interface{} `json:"config,omitempty"`
	// Array of tenant account ids that are allowed access
	Tenants *[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"tenants,omitempty"`
	ResourcePermissions *NetworkCreateResourcePermissions `json:"resourcePermissions,omitempty"`
}

// NewNetworkUpdate instantiates a new NetworkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkUpdate() *NetworkUpdate {
	this := NetworkUpdate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewNetworkUpdateWithDefaults instantiates a new NetworkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkUpdateWithDefaults() *NetworkUpdate {
	this := NetworkUpdate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkUpdate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkUpdate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NetworkUpdate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *NetworkUpdate) SetLabels(v []string) {
	o.Labels = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkUpdate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkUpdate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetworkUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetworkUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetworkUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *NetworkUpdate) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *NetworkUpdate) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *NetworkUpdate) SetCidr(v string) {
	o.Cidr = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *NetworkUpdate) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *NetworkUpdate) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *NetworkUpdate) SetGateway(v string) {
	o.Gateway = &v
}

// GetDnsPrimary returns the DnsPrimary field value if set, zero value otherwise.
func (o *NetworkUpdate) GetDnsPrimary() string {
	if o == nil || o.DnsPrimary == nil {
		var ret string
		return ret
	}
	return *o.DnsPrimary
}

// GetDnsPrimaryOk returns a tuple with the DnsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetDnsPrimaryOk() (*string, bool) {
	if o == nil || o.DnsPrimary == nil {
		return nil, false
	}
	return o.DnsPrimary, true
}

// HasDnsPrimary returns a boolean if a field has been set.
func (o *NetworkUpdate) HasDnsPrimary() bool {
	if o != nil && o.DnsPrimary != nil {
		return true
	}

	return false
}

// SetDnsPrimary gets a reference to the given string and assigns it to the DnsPrimary field.
func (o *NetworkUpdate) SetDnsPrimary(v string) {
	o.DnsPrimary = &v
}

// GetDnsSecondary returns the DnsSecondary field value if set, zero value otherwise.
func (o *NetworkUpdate) GetDnsSecondary() string {
	if o == nil || o.DnsSecondary == nil {
		var ret string
		return ret
	}
	return *o.DnsSecondary
}

// GetDnsSecondaryOk returns a tuple with the DnsSecondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetDnsSecondaryOk() (*string, bool) {
	if o == nil || o.DnsSecondary == nil {
		return nil, false
	}
	return o.DnsSecondary, true
}

// HasDnsSecondary returns a boolean if a field has been set.
func (o *NetworkUpdate) HasDnsSecondary() bool {
	if o != nil && o.DnsSecondary != nil {
		return true
	}

	return false
}

// SetDnsSecondary gets a reference to the given string and assigns it to the DnsSecondary field.
func (o *NetworkUpdate) SetDnsSecondary(v string) {
	o.DnsSecondary = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *NetworkUpdate) GetVlanId() int64 {
	if o == nil || o.VlanId == nil {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetVlanIdOk() (*int64, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *NetworkUpdate) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *NetworkUpdate) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkUpdate) GetPool() int64 {
	if o == nil || o.Pool.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkUpdate) GetPoolOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *NetworkUpdate) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableInt64 and assigns it to the Pool field.
func (o *NetworkUpdate) SetPool(v int64) {
	o.Pool.Set(&v)
}
// SetPoolNil sets the value for Pool to be an explicit nil
func (o *NetworkUpdate) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *NetworkUpdate) UnsetPool() {
	o.Pool.Unset()
}

// GetAllowStaticOverride returns the AllowStaticOverride field value if set, zero value otherwise.
func (o *NetworkUpdate) GetAllowStaticOverride() bool {
	if o == nil || o.AllowStaticOverride == nil {
		var ret bool
		return ret
	}
	return *o.AllowStaticOverride
}

// GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetAllowStaticOverrideOk() (*bool, bool) {
	if o == nil || o.AllowStaticOverride == nil {
		return nil, false
	}
	return o.AllowStaticOverride, true
}

// HasAllowStaticOverride returns a boolean if a field has been set.
func (o *NetworkUpdate) HasAllowStaticOverride() bool {
	if o != nil && o.AllowStaticOverride != nil {
		return true
	}

	return false
}

// SetAllowStaticOverride gets a reference to the given bool and assigns it to the AllowStaticOverride field.
func (o *NetworkUpdate) SetAllowStaticOverride(v bool) {
	o.AllowStaticOverride = &v
}

// GetAssignPublicIp returns the AssignPublicIp field value if set, zero value otherwise.
func (o *NetworkUpdate) GetAssignPublicIp() bool {
	if o == nil || o.AssignPublicIp == nil {
		var ret bool
		return ret
	}
	return *o.AssignPublicIp
}

// GetAssignPublicIpOk returns a tuple with the AssignPublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetAssignPublicIpOk() (*bool, bool) {
	if o == nil || o.AssignPublicIp == nil {
		return nil, false
	}
	return o.AssignPublicIp, true
}

// HasAssignPublicIp returns a boolean if a field has been set.
func (o *NetworkUpdate) HasAssignPublicIp() bool {
	if o != nil && o.AssignPublicIp != nil {
		return true
	}

	return false
}

// SetAssignPublicIp gets a reference to the given bool and assigns it to the AssignPublicIp field.
func (o *NetworkUpdate) SetAssignPublicIp(v bool) {
	o.AssignPublicIp = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NetworkUpdate) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *NetworkUpdate) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NetworkUpdate) SetActive(v bool) {
	o.Active = &v
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *NetworkUpdate) GetDhcpServer() bool {
	if o == nil || o.DhcpServer == nil {
		var ret bool
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetDhcpServerOk() (*bool, bool) {
	if o == nil || o.DhcpServer == nil {
		return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *NetworkUpdate) HasDhcpServer() bool {
	if o != nil && o.DhcpServer != nil {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given bool and assigns it to the DhcpServer field.
func (o *NetworkUpdate) SetDhcpServer(v bool) {
	o.DhcpServer = &v
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *NetworkUpdate) GetNetworkDomain() NetworkCreateNetworkDomain {
	if o == nil || o.NetworkDomain == nil {
		var ret NetworkCreateNetworkDomain
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetNetworkDomainOk() (*NetworkCreateNetworkDomain, bool) {
	if o == nil || o.NetworkDomain == nil {
		return nil, false
	}
	return o.NetworkDomain, true
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *NetworkUpdate) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain != nil {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given NetworkCreateNetworkDomain and assigns it to the NetworkDomain field.
func (o *NetworkUpdate) SetNetworkDomain(v NetworkCreateNetworkDomain) {
	o.NetworkDomain = &v
}

// GetSearchDomains returns the SearchDomains field value if set, zero value otherwise.
func (o *NetworkUpdate) GetSearchDomains() string {
	if o == nil || o.SearchDomains == nil {
		var ret string
		return ret
	}
	return *o.SearchDomains
}

// GetSearchDomainsOk returns a tuple with the SearchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetSearchDomainsOk() (*string, bool) {
	if o == nil || o.SearchDomains == nil {
		return nil, false
	}
	return o.SearchDomains, true
}

// HasSearchDomains returns a boolean if a field has been set.
func (o *NetworkUpdate) HasSearchDomains() bool {
	if o != nil && o.SearchDomains != nil {
		return true
	}

	return false
}

// SetSearchDomains gets a reference to the given string and assigns it to the SearchDomains field.
func (o *NetworkUpdate) SetSearchDomains(v string) {
	o.SearchDomains = &v
}

// GetNetworkProxy returns the NetworkProxy field value if set, zero value otherwise.
func (o *NetworkUpdate) GetNetworkProxy() NetworkCreateNetworkProxy {
	if o == nil || o.NetworkProxy == nil {
		var ret NetworkCreateNetworkProxy
		return ret
	}
	return *o.NetworkProxy
}

// GetNetworkProxyOk returns a tuple with the NetworkProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetNetworkProxyOk() (*NetworkCreateNetworkProxy, bool) {
	if o == nil || o.NetworkProxy == nil {
		return nil, false
	}
	return o.NetworkProxy, true
}

// HasNetworkProxy returns a boolean if a field has been set.
func (o *NetworkUpdate) HasNetworkProxy() bool {
	if o != nil && o.NetworkProxy != nil {
		return true
	}

	return false
}

// SetNetworkProxy gets a reference to the given NetworkCreateNetworkProxy and assigns it to the NetworkProxy field.
func (o *NetworkUpdate) SetNetworkProxy(v NetworkCreateNetworkProxy) {
	o.NetworkProxy = &v
}

// GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field value if set, zero value otherwise.
func (o *NetworkUpdate) GetApplianceUrlProxyBypass() bool {
	if o == nil || o.ApplianceUrlProxyBypass == nil {
		var ret bool
		return ret
	}
	return *o.ApplianceUrlProxyBypass
}

// GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetApplianceUrlProxyBypassOk() (*bool, bool) {
	if o == nil || o.ApplianceUrlProxyBypass == nil {
		return nil, false
	}
	return o.ApplianceUrlProxyBypass, true
}

// HasApplianceUrlProxyBypass returns a boolean if a field has been set.
func (o *NetworkUpdate) HasApplianceUrlProxyBypass() bool {
	if o != nil && o.ApplianceUrlProxyBypass != nil {
		return true
	}

	return false
}

// SetApplianceUrlProxyBypass gets a reference to the given bool and assigns it to the ApplianceUrlProxyBypass field.
func (o *NetworkUpdate) SetApplianceUrlProxyBypass(v bool) {
	o.ApplianceUrlProxyBypass = &v
}

// GetNoProxy returns the NoProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkUpdate) GetNoProxy() string {
	if o == nil || o.NoProxy.Get() == nil {
		var ret string
		return ret
	}
	return *o.NoProxy.Get()
}

// GetNoProxyOk returns a tuple with the NoProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkUpdate) GetNoProxyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NoProxy.Get(), o.NoProxy.IsSet()
}

// HasNoProxy returns a boolean if a field has been set.
func (o *NetworkUpdate) HasNoProxy() bool {
	if o != nil && o.NoProxy.IsSet() {
		return true
	}

	return false
}

// SetNoProxy gets a reference to the given NullableString and assigns it to the NoProxy field.
func (o *NetworkUpdate) SetNoProxy(v string) {
	o.NoProxy.Set(&v)
}
// SetNoProxyNil sets the value for NoProxy to be an explicit nil
func (o *NetworkUpdate) SetNoProxyNil() {
	o.NoProxy.Set(nil)
}

// UnsetNoProxy ensures that no value is present for NoProxy, not even an explicit nil
func (o *NetworkUpdate) UnsetNoProxy() {
	o.NoProxy.Unset()
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *NetworkUpdate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *NetworkUpdate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *NetworkUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NetworkUpdate) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NetworkUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *NetworkUpdate) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *NetworkUpdate) GetTenants() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.Tenants == nil {
		var ret []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetTenantsOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *NetworkUpdate) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the Tenants field.
func (o *NetworkUpdate) SetTenants(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Tenants = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *NetworkUpdate) GetResourcePermissions() NetworkCreateResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret NetworkCreateResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkUpdate) GetResourcePermissionsOk() (*NetworkCreateResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *NetworkUpdate) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given NetworkCreateResourcePermissions and assigns it to the ResourcePermissions field.
func (o *NetworkUpdate) SetResourcePermissions(v NetworkCreateResourcePermissions) {
	o.ResourcePermissions = &v
}

func (o NetworkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.DnsPrimary != nil {
		toSerialize["dnsPrimary"] = o.DnsPrimary
	}
	if o.DnsSecondary != nil {
		toSerialize["dnsSecondary"] = o.DnsSecondary
	}
	if o.VlanId != nil {
		toSerialize["vlanId"] = o.VlanId
	}
	if o.Pool.IsSet() {
		toSerialize["pool"] = o.Pool.Get()
	}
	if o.AllowStaticOverride != nil {
		toSerialize["allowStaticOverride"] = o.AllowStaticOverride
	}
	if o.AssignPublicIp != nil {
		toSerialize["assignPublicIp"] = o.AssignPublicIp
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.DhcpServer != nil {
		toSerialize["dhcpServer"] = o.DhcpServer
	}
	if o.NetworkDomain != nil {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	if o.SearchDomains != nil {
		toSerialize["searchDomains"] = o.SearchDomains
	}
	if o.NetworkProxy != nil {
		toSerialize["networkProxy"] = o.NetworkProxy
	}
	if o.ApplianceUrlProxyBypass != nil {
		toSerialize["applianceUrlProxyBypass"] = o.ApplianceUrlProxyBypass
	}
	if o.NoProxy.IsSet() {
		toSerialize["noProxy"] = o.NoProxy.Get()
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkUpdate struct {
	value *NetworkUpdate
	isSet bool
}

func (v NullableNetworkUpdate) Get() *NetworkUpdate {
	return v.value
}

func (v *NullableNetworkUpdate) Set(val *NetworkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkUpdate(val *NetworkUpdate) *NullableNetworkUpdate {
	return &NullableNetworkUpdate{value: val, isSet: true}
}

func (v NullableNetworkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


