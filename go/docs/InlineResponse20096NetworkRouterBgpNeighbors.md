# InlineResponse20096NetworkRouterBgpNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**ForwardingAddress** | Pointer to **NullableString** |  | [optional] 
**ProtocolAddress** | Pointer to **NullableString** |  | [optional] 
**RemoteAs** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 
**KeepAlive** | Pointer to **int64** |  | [optional] 
**HoldDown** | Pointer to **int64** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**RouteFilteringType** | Pointer to **string** |  | [optional] 
**RouteFilteringIn** | Pointer to **string** |  | [optional] 
**RouteFilteringOut** | Pointer to **string** |  | [optional] 
**BfdEnabled** | Pointer to **bool** |  | [optional] 
**BfdInterval** | Pointer to **int64** |  | [optional] 
**BfdMultiple** | Pointer to **int64** |  | [optional] 
**AllowAsIn** | Pointer to **bool** |  | [optional] 
**HopLimit** | Pointer to **int64** |  | [optional] 
**RestartMode** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**SyncSource** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to [**InlineResponse20096Config**](inline_response_200_96_config.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse20096NetworkRouterBgpNeighbors

`func NewInlineResponse20096NetworkRouterBgpNeighbors() *InlineResponse20096NetworkRouterBgpNeighbors`

NewInlineResponse20096NetworkRouterBgpNeighbors instantiates a new InlineResponse20096NetworkRouterBgpNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20096NetworkRouterBgpNeighborsWithDefaults

`func NewInlineResponse20096NetworkRouterBgpNeighborsWithDefaults() *InlineResponse20096NetworkRouterBgpNeighbors`

NewInlineResponse20096NetworkRouterBgpNeighborsWithDefaults instantiates a new InlineResponse20096NetworkRouterBgpNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### SetForwardingAddressNil

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetForwardingAddressNil(b bool)`

 SetForwardingAddressNil sets the value for ForwardingAddress to be an explicit nil

### UnsetForwardingAddress
`func (o *InlineResponse20096NetworkRouterBgpNeighbors) UnsetForwardingAddress()`

UnsetForwardingAddress ensures that no value is present for ForwardingAddress, not even an explicit nil
### GetProtocolAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### SetProtocolAddressNil

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetProtocolAddressNil(b bool)`

 SetProtocolAddressNil sets the value for ProtocolAddress to be an explicit nil

### UnsetProtocolAddress
`func (o *InlineResponse20096NetworkRouterBgpNeighbors) UnsetProtocolAddress()`

UnsetProtocolAddress ensures that no value is present for ProtocolAddress, not even an explicit nil
### GetRemoteAs

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *InlineResponse20096NetworkRouterBgpNeighbors) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRouteFilteringType

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetBfdEnabled() bool`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetBfdEnabledOk() (*bool, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetBfdEnabled(v bool)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetAllowAsIn() bool`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetAllowAsInOk() (*bool, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetAllowAsIn(v bool)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetProviderId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetSyncSource

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetInternalId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InlineResponse20096NetworkRouterBgpNeighbors) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRefType

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *InlineResponse20096NetworkRouterBgpNeighbors) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *InlineResponse20096NetworkRouterBgpNeighbors) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetConfig

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetConfig() InlineResponse20096Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetConfigOk() (*InlineResponse20096Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetConfig(v InlineResponse20096Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20096NetworkRouterBgpNeighbors) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


