# SecurityGroupLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**ZonePool** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **NullableString** |  | [optional] 
**GroupLayer** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecurityGroupLocations

`func NewSecurityGroupLocations() *SecurityGroupLocations`

NewSecurityGroupLocations instantiates a new SecurityGroupLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupLocationsWithDefaults

`func NewSecurityGroupLocationsWithDefaults() *SecurityGroupLocations`

NewSecurityGroupLocationsWithDefaults instantiates a new SecurityGroupLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupLocations) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupLocations) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupLocations) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupLocations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupLocations) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupLocations) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupLocations) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupLocations) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityGroupLocations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupLocations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupLocations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupLocations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityGroupLocations) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityGroupLocations) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *SecurityGroupLocations) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SecurityGroupLocations) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SecurityGroupLocations) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SecurityGroupLocations) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *SecurityGroupLocations) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *SecurityGroupLocations) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *SecurityGroupLocations) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *SecurityGroupLocations) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### SetIacIdNil

`func (o *SecurityGroupLocations) SetIacIdNil(b bool)`

 SetIacIdNil sets the value for IacId to be an explicit nil

### UnsetIacId
`func (o *SecurityGroupLocations) UnsetIacId()`

UnsetIacId ensures that no value is present for IacId, not even an explicit nil
### GetZone

`func (o *SecurityGroupLocations) GetZone() InlineResponse20082LoadBalancerInstanceSslCert`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SecurityGroupLocations) GetZoneOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SecurityGroupLocations) SetZone(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SecurityGroupLocations) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *SecurityGroupLocations) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *SecurityGroupLocations) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZonePool

`func (o *SecurityGroupLocations) GetZonePool() InlineResponse20082LoadBalancerInstanceSslCert`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *SecurityGroupLocations) GetZonePoolOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *SecurityGroupLocations) SetZonePool(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *SecurityGroupLocations) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### SetZonePoolNil

`func (o *SecurityGroupLocations) SetZonePoolNil(b bool)`

 SetZonePoolNil sets the value for ZonePool to be an explicit nil

### UnsetZonePool
`func (o *SecurityGroupLocations) UnsetZonePool()`

UnsetZonePool ensures that no value is present for ZonePool, not even an explicit nil
### GetStatus

`func (o *SecurityGroupLocations) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityGroupLocations) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityGroupLocations) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityGroupLocations) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *SecurityGroupLocations) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SecurityGroupLocations) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SecurityGroupLocations) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SecurityGroupLocations) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *SecurityGroupLocations) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *SecurityGroupLocations) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetGroupLayer

`func (o *SecurityGroupLocations) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *SecurityGroupLocations) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *SecurityGroupLocations) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *SecurityGroupLocations) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.

### SetGroupLayerNil

`func (o *SecurityGroupLocations) SetGroupLayerNil(b bool)`

 SetGroupLayerNil sets the value for GroupLayer to be an explicit nil

### UnsetGroupLayer
`func (o *SecurityGroupLocations) UnsetGroupLayer()`

UnsetGroupLayer ensures that no value is present for GroupLayer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


