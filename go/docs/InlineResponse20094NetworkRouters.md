# InlineResponse20094NetworkRouters

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
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Type**](inline_response_200_94_type.md) |  | [optional] 
**NetworkServer** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Instance** | Pointer to **NullableString** |  | [optional] 
**ExternalNetwork** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Site** | Pointer to **NullableString** |  | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse20094Interfaces**](InlineResponse20094Interfaces.md) |  | [optional] 

## Methods

### NewInlineResponse20094NetworkRouters

`func NewInlineResponse20094NetworkRouters() *InlineResponse20094NetworkRouters`

NewInlineResponse20094NetworkRouters instantiates a new InlineResponse20094NetworkRouters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20094NetworkRoutersWithDefaults

`func NewInlineResponse20094NetworkRoutersWithDefaults() *InlineResponse20094NetworkRouters`

NewInlineResponse20094NetworkRoutersWithDefaults instantiates a new InlineResponse20094NetworkRouters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20094NetworkRouters) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20094NetworkRouters) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20094NetworkRouters) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20094NetworkRouters) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20094NetworkRouters) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20094NetworkRouters) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20094NetworkRouters) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20094NetworkRouters) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20094NetworkRouters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20094NetworkRouters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20094NetworkRouters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20094NetworkRouters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20094NetworkRouters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20094NetworkRouters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20094NetworkRouters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20094NetworkRouters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse20094NetworkRouters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse20094NetworkRouters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *InlineResponse20094NetworkRouters) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20094NetworkRouters) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20094NetworkRouters) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20094NetworkRouters) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20094NetworkRouters) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20094NetworkRouters) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20094NetworkRouters) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20094NetworkRouters) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20094NetworkRouters) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20094NetworkRouters) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20094NetworkRouters) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20094NetworkRouters) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRouterType

`func (o *InlineResponse20094NetworkRouters) GetRouterType() string`

GetRouterType returns the RouterType field if non-nil, zero value otherwise.

### GetRouterTypeOk

`func (o *InlineResponse20094NetworkRouters) GetRouterTypeOk() (*string, bool)`

GetRouterTypeOk returns a tuple with the RouterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterType

`func (o *InlineResponse20094NetworkRouters) SetRouterType(v string)`

SetRouterType sets RouterType field to given value.

### HasRouterType

`func (o *InlineResponse20094NetworkRouters) HasRouterType() bool`

HasRouterType returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20094NetworkRouters) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20094NetworkRouters) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20094NetworkRouters) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20094NetworkRouters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20094NetworkRouters) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20094NetworkRouters) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20094NetworkRouters) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20094NetworkRouters) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalIp

`func (o *InlineResponse20094NetworkRouters) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *InlineResponse20094NetworkRouters) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *InlineResponse20094NetworkRouters) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *InlineResponse20094NetworkRouters) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *InlineResponse20094NetworkRouters) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *InlineResponse20094NetworkRouters) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetExternalId

`func (o *InlineResponse20094NetworkRouters) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20094NetworkRouters) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20094NetworkRouters) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20094NetworkRouters) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *InlineResponse20094NetworkRouters) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *InlineResponse20094NetworkRouters) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *InlineResponse20094NetworkRouters) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *InlineResponse20094NetworkRouters) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *InlineResponse20094NetworkRouters) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *InlineResponse20094NetworkRouters) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetType

`func (o *InlineResponse20094NetworkRouters) GetType() InlineResponse20094Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20094NetworkRouters) GetTypeOk() (*InlineResponse20094Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20094NetworkRouters) SetType(v InlineResponse20094Type)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20094NetworkRouters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkServer

`func (o *InlineResponse20094NetworkRouters) GetNetworkServer() string`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *InlineResponse20094NetworkRouters) GetNetworkServerOk() (*string, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *InlineResponse20094NetworkRouters) SetNetworkServer(v string)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *InlineResponse20094NetworkRouters) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### SetNetworkServerNil

`func (o *InlineResponse20094NetworkRouters) SetNetworkServerNil(b bool)`

 SetNetworkServerNil sets the value for NetworkServer to be an explicit nil

### UnsetNetworkServer
`func (o *InlineResponse20094NetworkRouters) UnsetNetworkServer()`

UnsetNetworkServer ensures that no value is present for NetworkServer, not even an explicit nil
### GetZone

`func (o *InlineResponse20094NetworkRouters) GetZone() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InlineResponse20094NetworkRouters) GetZoneOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InlineResponse20094NetworkRouters) SetZone(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InlineResponse20094NetworkRouters) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetInstance

`func (o *InlineResponse20094NetworkRouters) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InlineResponse20094NetworkRouters) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InlineResponse20094NetworkRouters) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InlineResponse20094NetworkRouters) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *InlineResponse20094NetworkRouters) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *InlineResponse20094NetworkRouters) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetExternalNetwork

`func (o *InlineResponse20094NetworkRouters) GetExternalNetwork() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetExternalNetwork returns the ExternalNetwork field if non-nil, zero value otherwise.

### GetExternalNetworkOk

`func (o *InlineResponse20094NetworkRouters) GetExternalNetworkOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetExternalNetworkOk returns a tuple with the ExternalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetwork

`func (o *InlineResponse20094NetworkRouters) SetExternalNetwork(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetExternalNetwork sets ExternalNetwork field to given value.

### HasExternalNetwork

`func (o *InlineResponse20094NetworkRouters) HasExternalNetwork() bool`

HasExternalNetwork returns a boolean if a field has been set.

### GetSite

`func (o *InlineResponse20094NetworkRouters) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InlineResponse20094NetworkRouters) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InlineResponse20094NetworkRouters) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *InlineResponse20094NetworkRouters) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *InlineResponse20094NetworkRouters) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *InlineResponse20094NetworkRouters) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetInterfaces

`func (o *InlineResponse20094NetworkRouters) GetInterfaces() []InlineResponse20094Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse20094NetworkRouters) GetInterfacesOk() (*[]InlineResponse20094Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse20094NetworkRouters) SetInterfaces(v []InlineResponse20094Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse20094NetworkRouters) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


