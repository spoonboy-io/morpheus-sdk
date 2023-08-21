# NetworkPoolCreateIpRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAddress** | Pointer to **string** | Starting IP Address | [optional] 
**EndAddress** | Pointer to **string** | Ending IP Address | [optional] 
**CidrIPv6** | Pointer to **string** | IPv6 Network CIDR | [optional] 

## Methods

### NewNetworkPoolCreateIpRanges

`func NewNetworkPoolCreateIpRanges() *NetworkPoolCreateIpRanges`

NewNetworkPoolCreateIpRanges instantiates a new NetworkPoolCreateIpRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolCreateIpRangesWithDefaults

`func NewNetworkPoolCreateIpRangesWithDefaults() *NetworkPoolCreateIpRanges`

NewNetworkPoolCreateIpRangesWithDefaults instantiates a new NetworkPoolCreateIpRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAddress

`func (o *NetworkPoolCreateIpRanges) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *NetworkPoolCreateIpRanges) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *NetworkPoolCreateIpRanges) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *NetworkPoolCreateIpRanges) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *NetworkPoolCreateIpRanges) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *NetworkPoolCreateIpRanges) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *NetworkPoolCreateIpRanges) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *NetworkPoolCreateIpRanges) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetCidrIPv6

`func (o *NetworkPoolCreateIpRanges) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *NetworkPoolCreateIpRanges) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *NetworkPoolCreateIpRanges) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *NetworkPoolCreateIpRanges) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


