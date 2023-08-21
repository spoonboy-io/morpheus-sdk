# InlineResponse200108NetworkFloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Cloud** | Pointer to [**InlineResponse200108NetworkFloatingIpCloud**](inline_response_200_108_networkFloatingIp_cloud.md) |  | [optional] 
**Server** | Pointer to [**NullableInlineResponse200108NetworkFloatingIpServer**](inline_response_200_108_networkFloatingIp_server.md) |  | [optional] 
**IpStatus** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**IpRange** | Pointer to **NullableString** |  | [optional] 
**PtrId** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**NullableInlineResponse200108NetworkFloatingIpNetworkDomain**](inline_response_200_108_networkFloatingIp_networkDomain.md) |  | [optional] 
**CreatedBy** | Pointer to [**NullableInlineResponse200108NetworkFloatingIpCreatedBy**](inline_response_200_108_networkFloatingIp_createdBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse200108NetworkFloatingIp

`func NewInlineResponse200108NetworkFloatingIp() *InlineResponse200108NetworkFloatingIp`

NewInlineResponse200108NetworkFloatingIp instantiates a new InlineResponse200108NetworkFloatingIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108NetworkFloatingIpWithDefaults

`func NewInlineResponse200108NetworkFloatingIpWithDefaults() *InlineResponse200108NetworkFloatingIp`

NewInlineResponse200108NetworkFloatingIpWithDefaults instantiates a new InlineResponse200108NetworkFloatingIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200108NetworkFloatingIp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200108NetworkFloatingIp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200108NetworkFloatingIp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200108NetworkFloatingIp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *InlineResponse200108NetworkFloatingIp) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse200108NetworkFloatingIp) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse200108NetworkFloatingIp) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse200108NetworkFloatingIp) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InlineResponse200108NetworkFloatingIp) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InlineResponse200108NetworkFloatingIp) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCloud

`func (o *InlineResponse200108NetworkFloatingIp) GetCloud() InlineResponse200108NetworkFloatingIpCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InlineResponse200108NetworkFloatingIp) GetCloudOk() (*InlineResponse200108NetworkFloatingIpCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InlineResponse200108NetworkFloatingIp) SetCloud(v InlineResponse200108NetworkFloatingIpCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InlineResponse200108NetworkFloatingIp) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetServer

`func (o *InlineResponse200108NetworkFloatingIp) GetServer() InlineResponse200108NetworkFloatingIpServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InlineResponse200108NetworkFloatingIp) GetServerOk() (*InlineResponse200108NetworkFloatingIpServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InlineResponse200108NetworkFloatingIp) SetServer(v InlineResponse200108NetworkFloatingIpServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *InlineResponse200108NetworkFloatingIp) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *InlineResponse200108NetworkFloatingIp) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *InlineResponse200108NetworkFloatingIp) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetIpStatus

`func (o *InlineResponse200108NetworkFloatingIp) GetIpStatus() string`

GetIpStatus returns the IpStatus field if non-nil, zero value otherwise.

### GetIpStatusOk

`func (o *InlineResponse200108NetworkFloatingIp) GetIpStatusOk() (*string, bool)`

GetIpStatusOk returns a tuple with the IpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStatus

`func (o *InlineResponse200108NetworkFloatingIp) SetIpStatus(v string)`

SetIpStatus sets IpStatus field to given value.

### HasIpStatus

`func (o *InlineResponse200108NetworkFloatingIp) HasIpStatus() bool`

HasIpStatus returns a boolean if a field has been set.

### GetIpAddress

`func (o *InlineResponse200108NetworkFloatingIp) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineResponse200108NetworkFloatingIp) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineResponse200108NetworkFloatingIp) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineResponse200108NetworkFloatingIp) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpRange

`func (o *InlineResponse200108NetworkFloatingIp) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *InlineResponse200108NetworkFloatingIp) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *InlineResponse200108NetworkFloatingIp) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *InlineResponse200108NetworkFloatingIp) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### SetIpRangeNil

`func (o *InlineResponse200108NetworkFloatingIp) SetIpRangeNil(b bool)`

 SetIpRangeNil sets the value for IpRange to be an explicit nil

### UnsetIpRange
`func (o *InlineResponse200108NetworkFloatingIp) UnsetIpRange()`

UnsetIpRange ensures that no value is present for IpRange, not even an explicit nil
### GetPtrId

`func (o *InlineResponse200108NetworkFloatingIp) GetPtrId() string`

GetPtrId returns the PtrId field if non-nil, zero value otherwise.

### GetPtrIdOk

`func (o *InlineResponse200108NetworkFloatingIp) GetPtrIdOk() (*string, bool)`

GetPtrIdOk returns a tuple with the PtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrId

`func (o *InlineResponse200108NetworkFloatingIp) SetPtrId(v string)`

SetPtrId sets PtrId field to given value.

### HasPtrId

`func (o *InlineResponse200108NetworkFloatingIp) HasPtrId() bool`

HasPtrId returns a boolean if a field has been set.

### SetPtrIdNil

`func (o *InlineResponse200108NetworkFloatingIp) SetPtrIdNil(b bool)`

 SetPtrIdNil sets the value for PtrId to be an explicit nil

### UnsetPtrId
`func (o *InlineResponse200108NetworkFloatingIp) UnsetPtrId()`

UnsetPtrId ensures that no value is present for PtrId, not even an explicit nil
### GetNetworkDomain

`func (o *InlineResponse200108NetworkFloatingIp) GetNetworkDomain() InlineResponse200108NetworkFloatingIpNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *InlineResponse200108NetworkFloatingIp) GetNetworkDomainOk() (*InlineResponse200108NetworkFloatingIpNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *InlineResponse200108NetworkFloatingIp) SetNetworkDomain(v InlineResponse200108NetworkFloatingIpNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *InlineResponse200108NetworkFloatingIp) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *InlineResponse200108NetworkFloatingIp) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *InlineResponse200108NetworkFloatingIp) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetCreatedBy

`func (o *InlineResponse200108NetworkFloatingIp) GetCreatedBy() InlineResponse200108NetworkFloatingIpCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse200108NetworkFloatingIp) GetCreatedByOk() (*InlineResponse200108NetworkFloatingIpCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse200108NetworkFloatingIp) SetCreatedBy(v InlineResponse200108NetworkFloatingIpCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse200108NetworkFloatingIp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *InlineResponse200108NetworkFloatingIp) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InlineResponse200108NetworkFloatingIp) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *InlineResponse200108NetworkFloatingIp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse200108NetworkFloatingIp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse200108NetworkFloatingIp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse200108NetworkFloatingIp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse200108NetworkFloatingIp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse200108NetworkFloatingIp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse200108NetworkFloatingIp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse200108NetworkFloatingIp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


