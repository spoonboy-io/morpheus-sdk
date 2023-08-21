# InlineResponse200106NetworkPoolIpRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**StartAddress** | Pointer to **NullableString** |  | [optional] 
**EndAddress** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AddressCount** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Cidr** | Pointer to **NullableString** |  | [optional] 
**CidrIPv6** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineResponse200106NetworkPoolIpRanges

`func NewInlineResponse200106NetworkPoolIpRanges() *InlineResponse200106NetworkPoolIpRanges`

NewInlineResponse200106NetworkPoolIpRanges instantiates a new InlineResponse200106NetworkPoolIpRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200106NetworkPoolIpRangesWithDefaults

`func NewInlineResponse200106NetworkPoolIpRangesWithDefaults() *InlineResponse200106NetworkPoolIpRanges`

NewInlineResponse200106NetworkPoolIpRangesWithDefaults instantiates a new InlineResponse200106NetworkPoolIpRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200106NetworkPoolIpRanges) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200106NetworkPoolIpRanges) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200106NetworkPoolIpRanges) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartAddress

`func (o *InlineResponse200106NetworkPoolIpRanges) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *InlineResponse200106NetworkPoolIpRanges) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *InlineResponse200106NetworkPoolIpRanges) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### SetStartAddressNil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetStartAddressNil(b bool)`

 SetStartAddressNil sets the value for StartAddress to be an explicit nil

### UnsetStartAddress
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetStartAddress()`

UnsetStartAddress ensures that no value is present for StartAddress, not even an explicit nil
### GetEndAddress

`func (o *InlineResponse200106NetworkPoolIpRanges) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *InlineResponse200106NetworkPoolIpRanges) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *InlineResponse200106NetworkPoolIpRanges) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### SetEndAddressNil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetEndAddressNil(b bool)`

 SetEndAddressNil sets the value for EndAddress to be an explicit nil

### UnsetEndAddress
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetEndAddress()`

UnsetEndAddress ensures that no value is present for EndAddress, not even an explicit nil
### GetInternalId

`func (o *InlineResponse200106NetworkPoolIpRanges) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse200106NetworkPoolIpRanges) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse200106NetworkPoolIpRanges) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *InlineResponse200106NetworkPoolIpRanges) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse200106NetworkPoolIpRanges) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse200106NetworkPoolIpRanges) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDescription

`func (o *InlineResponse200106NetworkPoolIpRanges) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200106NetworkPoolIpRanges) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200106NetworkPoolIpRanges) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAddressCount

`func (o *InlineResponse200106NetworkPoolIpRanges) GetAddressCount() int64`

GetAddressCount returns the AddressCount field if non-nil, zero value otherwise.

### GetAddressCountOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetAddressCountOk() (*int64, bool)`

GetAddressCountOk returns a tuple with the AddressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCount

`func (o *InlineResponse200106NetworkPoolIpRanges) SetAddressCount(v int64)`

SetAddressCount sets AddressCount field to given value.

### HasAddressCount

`func (o *InlineResponse200106NetworkPoolIpRanges) HasAddressCount() bool`

HasAddressCount returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200106NetworkPoolIpRanges) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200106NetworkPoolIpRanges) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200106NetworkPoolIpRanges) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse200106NetworkPoolIpRanges) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse200106NetworkPoolIpRanges) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse200106NetworkPoolIpRanges) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse200106NetworkPoolIpRanges) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse200106NetworkPoolIpRanges) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse200106NetworkPoolIpRanges) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse200106NetworkPoolIpRanges) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse200106NetworkPoolIpRanges) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse200106NetworkPoolIpRanges) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse200106NetworkPoolIpRanges) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetCidrIPv6

`func (o *InlineResponse200106NetworkPoolIpRanges) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *InlineResponse200106NetworkPoolIpRanges) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *InlineResponse200106NetworkPoolIpRanges) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *InlineResponse200106NetworkPoolIpRanges) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### SetCidrIPv6Nil

`func (o *InlineResponse200106NetworkPoolIpRanges) SetCidrIPv6Nil(b bool)`

 SetCidrIPv6Nil sets the value for CidrIPv6 to be an explicit nil

### UnsetCidrIPv6
`func (o *InlineResponse200106NetworkPoolIpRanges) UnsetCidrIPv6()`

UnsetCidrIPv6 ensures that no value is present for CidrIPv6, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


