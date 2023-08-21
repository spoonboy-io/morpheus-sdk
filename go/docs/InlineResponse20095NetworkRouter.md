# InlineResponse20095NetworkRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RouterType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**InlineResponse20095NetworkRouterType**](inline_response_200_95_networkRouter_type.md) |  | [optional] 
**NetworkServer** | Pointer to [**InlineResponse20095NetworkRouterNetworkServer**](inline_response_200_95_networkRouter_networkServer.md) |  | [optional] 
**Zone** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Instance** | Pointer to **NullableString** |  | [optional] 
**ExternalNetwork** | Pointer to **NullableString** |  | [optional] 
**Site** | Pointer to **NullableString** |  | [optional] 
**Interfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Firewall** | Pointer to [**InlineResponse20095NetworkRouterFirewall**](inline_response_200_95_networkRouter_firewall.md) |  | [optional] 
**Routes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Nats** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Permissions** | Pointer to [**InlineResponse20095NetworkRouterPermissions**](inline_response_200_95_networkRouter_permissions.md) |  | [optional] 

## Methods

### NewInlineResponse20095NetworkRouter

`func NewInlineResponse20095NetworkRouter() *InlineResponse20095NetworkRouter`

NewInlineResponse20095NetworkRouter instantiates a new InlineResponse20095NetworkRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095NetworkRouterWithDefaults

`func NewInlineResponse20095NetworkRouterWithDefaults() *InlineResponse20095NetworkRouter`

NewInlineResponse20095NetworkRouterWithDefaults instantiates a new InlineResponse20095NetworkRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20095NetworkRouter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20095NetworkRouter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20095NetworkRouter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20095NetworkRouter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20095NetworkRouter) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20095NetworkRouter) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20095NetworkRouter) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20095NetworkRouter) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20095NetworkRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20095NetworkRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20095NetworkRouter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20095NetworkRouter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20095NetworkRouter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20095NetworkRouter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20095NetworkRouter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20095NetworkRouter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse20095NetworkRouter) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse20095NetworkRouter) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *InlineResponse20095NetworkRouter) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20095NetworkRouter) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20095NetworkRouter) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20095NetworkRouter) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20095NetworkRouter) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20095NetworkRouter) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20095NetworkRouter) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20095NetworkRouter) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20095NetworkRouter) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20095NetworkRouter) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20095NetworkRouter) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20095NetworkRouter) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRouterType

`func (o *InlineResponse20095NetworkRouter) GetRouterType() string`

GetRouterType returns the RouterType field if non-nil, zero value otherwise.

### GetRouterTypeOk

`func (o *InlineResponse20095NetworkRouter) GetRouterTypeOk() (*string, bool)`

GetRouterTypeOk returns a tuple with the RouterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterType

`func (o *InlineResponse20095NetworkRouter) SetRouterType(v string)`

SetRouterType sets RouterType field to given value.

### HasRouterType

`func (o *InlineResponse20095NetworkRouter) HasRouterType() bool`

HasRouterType returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20095NetworkRouter) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20095NetworkRouter) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20095NetworkRouter) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20095NetworkRouter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20095NetworkRouter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20095NetworkRouter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20095NetworkRouter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20095NetworkRouter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalIp

`func (o *InlineResponse20095NetworkRouter) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *InlineResponse20095NetworkRouter) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *InlineResponse20095NetworkRouter) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *InlineResponse20095NetworkRouter) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *InlineResponse20095NetworkRouter) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *InlineResponse20095NetworkRouter) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetExternalId

`func (o *InlineResponse20095NetworkRouter) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20095NetworkRouter) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20095NetworkRouter) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20095NetworkRouter) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *InlineResponse20095NetworkRouter) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *InlineResponse20095NetworkRouter) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *InlineResponse20095NetworkRouter) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *InlineResponse20095NetworkRouter) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20095NetworkRouter) GetType() InlineResponse20095NetworkRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20095NetworkRouter) GetTypeOk() (*InlineResponse20095NetworkRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20095NetworkRouter) SetType(v InlineResponse20095NetworkRouterType)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20095NetworkRouter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkServer

`func (o *InlineResponse20095NetworkRouter) GetNetworkServer() InlineResponse20095NetworkRouterNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *InlineResponse20095NetworkRouter) GetNetworkServerOk() (*InlineResponse20095NetworkRouterNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *InlineResponse20095NetworkRouter) SetNetworkServer(v InlineResponse20095NetworkRouterNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *InlineResponse20095NetworkRouter) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetZone

`func (o *InlineResponse20095NetworkRouter) GetZone() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InlineResponse20095NetworkRouter) GetZoneOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InlineResponse20095NetworkRouter) SetZone(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InlineResponse20095NetworkRouter) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetInstance

`func (o *InlineResponse20095NetworkRouter) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InlineResponse20095NetworkRouter) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InlineResponse20095NetworkRouter) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InlineResponse20095NetworkRouter) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *InlineResponse20095NetworkRouter) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *InlineResponse20095NetworkRouter) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetExternalNetwork

`func (o *InlineResponse20095NetworkRouter) GetExternalNetwork() string`

GetExternalNetwork returns the ExternalNetwork field if non-nil, zero value otherwise.

### GetExternalNetworkOk

`func (o *InlineResponse20095NetworkRouter) GetExternalNetworkOk() (*string, bool)`

GetExternalNetworkOk returns a tuple with the ExternalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetwork

`func (o *InlineResponse20095NetworkRouter) SetExternalNetwork(v string)`

SetExternalNetwork sets ExternalNetwork field to given value.

### HasExternalNetwork

`func (o *InlineResponse20095NetworkRouter) HasExternalNetwork() bool`

HasExternalNetwork returns a boolean if a field has been set.

### SetExternalNetworkNil

`func (o *InlineResponse20095NetworkRouter) SetExternalNetworkNil(b bool)`

 SetExternalNetworkNil sets the value for ExternalNetwork to be an explicit nil

### UnsetExternalNetwork
`func (o *InlineResponse20095NetworkRouter) UnsetExternalNetwork()`

UnsetExternalNetwork ensures that no value is present for ExternalNetwork, not even an explicit nil
### GetSite

`func (o *InlineResponse20095NetworkRouter) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InlineResponse20095NetworkRouter) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InlineResponse20095NetworkRouter) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *InlineResponse20095NetworkRouter) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *InlineResponse20095NetworkRouter) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *InlineResponse20095NetworkRouter) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetInterfaces

`func (o *InlineResponse20095NetworkRouter) GetInterfaces() []map[string]interface{}`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse20095NetworkRouter) GetInterfacesOk() (*[]map[string]interface{}, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse20095NetworkRouter) SetInterfaces(v []map[string]interface{})`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse20095NetworkRouter) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetFirewall

`func (o *InlineResponse20095NetworkRouter) GetFirewall() InlineResponse20095NetworkRouterFirewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *InlineResponse20095NetworkRouter) GetFirewallOk() (*InlineResponse20095NetworkRouterFirewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *InlineResponse20095NetworkRouter) SetFirewall(v InlineResponse20095NetworkRouterFirewall)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *InlineResponse20095NetworkRouter) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetRoutes

`func (o *InlineResponse20095NetworkRouter) GetRoutes() []map[string]interface{}`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *InlineResponse20095NetworkRouter) GetRoutesOk() (*[]map[string]interface{}, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *InlineResponse20095NetworkRouter) SetRoutes(v []map[string]interface{})`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *InlineResponse20095NetworkRouter) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetNats

`func (o *InlineResponse20095NetworkRouter) GetNats() []map[string]interface{}`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *InlineResponse20095NetworkRouter) GetNatsOk() (*[]map[string]interface{}, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *InlineResponse20095NetworkRouter) SetNats(v []map[string]interface{})`

SetNats sets Nats field to given value.

### HasNats

`func (o *InlineResponse20095NetworkRouter) HasNats() bool`

HasNats returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse20095NetworkRouter) GetPermissions() InlineResponse20095NetworkRouterPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse20095NetworkRouter) GetPermissionsOk() (*InlineResponse20095NetworkRouterPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse20095NetworkRouter) SetPermissions(v InlineResponse20095NetworkRouterPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse20095NetworkRouter) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


